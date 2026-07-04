# 로또 번호 추천 알고리즘

`sqls/schemas`, `sqls/ddls`에 정의된 테이블 구조를 기반으로 한 추천 번호 생성 알고리즘 설계.

## 1. 관련 테이블 매핑

| 테이블 | 역할 |
|---|---|
| `gak_soo` | 1~45 각 번호 마스터 |
| `gak_soo_status` | 번호별 등수 출현 횟수 / 확률 (배치로 갱신) |
| `lotto_rounds` | 공식 회차 (추첨일, 보너스번호) |
| `lotto_rounds_numbers` | 회차별 당첨번호 6개 |
| `recommend_sets` | 생성된 추천 세트 (회차, 알고리즘 버전, 신뢰도 점수) |
| `recommend_set_numbers` | 추천 세트에 속한 번호 6개 + 가중치 |
| `recommend_results` | 추천 세트가 실제 회차와 몇 개 맞았는지 사후 검증 |

즉 파이프라인은 **통계 갱신 → 세트 생성 → 사후 검증**의 3단계로 구성된다.

## 2. 1단계 — 번호별 통계 갱신 (`gak_soo_status`)

매주 일요일 아침 배치(신규 회차 등록 직후) 또는 서버 기동 시 계산.

1. 새로 등록된 `lotto_rounds_numbers` + `lotto_rounds.bonus_number`를 순회하며, 당첨번호 6개는 `rank`(1~5등, 회차의 `recommend_results`가 아니라 공식 등수 규칙: 6개 일치=1등, 5개+보너스=2등, 5개=3등, 4개=4등, 3개=5등)에 따라 해당 `gak_soo_status`의 `Nth_count`를 +1.
2. 전체 회차 수 대비 등장 횟수로 확률 계산:
   ```
   first_probability  = first_count  / total_rounds
   second_probability = second_count / total_rounds
   ```
   (스키마상 3~5등 확률 컬럼은 아직 없음 — 필요 시 `third_probability` 등 컬럼 추가 필요)
3. `u_gak_soo_status.sql`로 45개 번호 전부 UPSERT.

### 가중 스코어 (번호별 종합 점수)

등수별 실제 배당/희귀도가 다르므로 등수별 가중치를 곱해 하나의 점수로 합산한다. 가중치는 `recommend_sets.algorithm_version`으로 버저닝해서 튜닝 가능하게 둔다.

```
weighted_score(soo) =
    W1 * first_probability
  + W2 * second_probability
  + W3 * third_probability
  + W4 * fourth_probability
  + W5 * fifth_probability
```

기본값 제안: `W1=5.0, W2=3.0, W3=2.0, W4=1.0, W5=0.5` (등수가 높을수록 희귀하니 더 큰 가중치).

### 최근성 보정 (recency decay)

전체 기간 누적 확률만 쓰면 오래된 회차와 최근 회차 비중이 같아진다. 최근 N회차(예: 최근 52회 = 1년)에 지수 감쇠 가중치를 추가로 곱해서 "최근에 잘 나오는 번호"에 더 힘을 싣는다.

```
recency_weight(round) = decay^(current_round - round_no)   // decay = 0.98 정도
adjusted_score(soo) = weighted_score(soo) * (1 + Σ recency_weight(round) for rounds where soo appeared)
```

## 3. 2단계 — 추천 세트 생성 (`recommend_sets` / `recommend_set_numbers`)

단순히 `adjusted_score` 상위 6개만 뽑으면 매 회차 똑같은 번호만 추천하게 되므로, **가중 랜덤 샘플링 + 제약조건 필터**로 세트를 여러 개(예: 5세트) 생성한다.

### 3-1. 가중 랜덤 샘플링

- 45개 번호의 `adjusted_score`를 정규화해 확률분포로 변환.
- 확률에 비례한 가중 추출(weighted sampling without replacement)로 6개 번호 선택.
  - Go 구현 시 누적합(prefix sum) + `rand.Float64()` 이분탐색 방식이 일반적.
- 이 과정을 세트 개수만큼 반복(세트마다 독립적으로 추출 → 세트 간 다양성 확보).

### 3-2. 제약조건 필터 (통계적으로 흔한 당첨 패턴에 맞춰 이상치 제거)

샘플링된 6개 조합이 아래 조건을 만족하지 않으면 재추출한다.

| 조건 | 기준 |
|---|---|
| 합계 범위 | 100 ~ 175 (역대 당첨번호 합의 대부분이 이 구간) |
| 홀짝 비율 | 전부 홀수/전부 짝수 제외 (2:4, 3:3, 4:2 허용) |
| 구간 분포 | 1~15 / 16~30 / 31~45 각 구간에 최소 1개 이상 |
| 연속번호 | 3개 이상 연속 숫자 금지 (예: 12,13,14) |
| 과거 당첨조합 중복 | `lotto_rounds_numbers`에 존재하는 완전 동일 조합 제외 |

### 3-3. 세트 점수 산출

```
recommend_sets.score = Σ recommend_set_numbers.weights / 6   // 평균 가중치
```

`recommend_set_numbers.weights`에는 선택 당시 각 번호의 `adjusted_score`(정규화값)를 저장 → 세트 신뢰도와 개별 번호 근거를 동시에 남긴다.

## 4. 3단계 — 사후 검증 (`recommend_results`)

새 회차 결과가 `lotto_rounds_numbers`에 등록되면, 해당 회차에 걸린 모든 `recommend_sets`에 대해:

1. `recommend_set_numbers` 6개 vs `lotto_rounds_numbers` 6개 교집합 개수 = `matched_count`
2. 교집합에 `bonus_number` 포함 여부 = `bonus_match`
3. 공식 등수 규칙으로 `rank` 산출 (6개=1등, 5개+보너스=2등, 5개=3등, 4개=4등, 3개=5등, 그 외=낙첨)
4. `u_recommend_results.sql`로 UPDATE (없으면 INSERT).

이 결과는 다시 `algorithm_version`별로 집계해 어떤 가중치(W1~W5, decay)가 실제 적중률이 높았는지 비교하는 피드백 루프로 쓰인다 — 다음 버전 가중치 튜닝의 근거 데이터.

## 5. 처리 흐름 요약

```
[배치: 신규 회차 등록]
        │
        ▼
① gak_soo_status 갱신 (count, probability)
        │
        ▼
② weighted_score → recency 보정 → adjusted_score (45개 번호)
        │
        ▼
③ 가중 랜덤 샘플링 (세트 N개) → 제약조건 필터 통과할 때까지 재추출
        │
        ▼
④ recommend_sets / recommend_set_numbers 저장 (score, weights, rank)
        │
        ▼
   (한 주 대기, 실제 추첨 진행)
        │
        ▼
⑤ lotto_rounds_numbers 결과 등록
        │
        ▼
⑥ recommend_results 계산/저장 → 알고리즘 버전별 성능 비교
```

## 6. 구현 시 레이어 매핑 (기존 handler/service/repo 컨벤션)

- `recommend/repo`: `RUpsertGakSooStatus`, `RCreateRecommendSet`, `RCreateRecommendSetNumbers`, `RUpsertRecommendResult` — 각각 대응 SQL 파일(`sqls/ddls/*`)을 `os.ReadFile`로 읽어 실행.
- `recommend/service`: `SGenerateRecommendSets(ctx, roundID, setCount)` — 위 파이프라인 ①~④를 조합해 트랜잭션으로 실행.
- `recommend/service`: `SEvaluateRecommendResults(ctx, roundID)` — ⑤~⑥ 실행, 신규 회차 확정 후 배치/핸들러에서 호출.
- `recommend/handler`: 사용자가 "이번 회차 추천 세트 조회" API를 호출하면 `recommend_sets` + `recommend_set_numbers`를 JOIN해 JSON으로 반환.

## 7. 미확정 사항 (설계 시 결정 필요)

- `gak_soo_status`에 3~5등 확률 컬럼이 없음 → 컬럼 추가 여부 결정 필요.
- 세트 개수(N), decay 계수, 홀짝/구간 제약의 정확한 임계값은 초기값으로 시작 후 `recommend_results` 데이터가 쌓이면 실측 기반으로 조정.
