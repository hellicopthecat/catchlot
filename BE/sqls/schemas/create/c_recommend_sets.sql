/*
* 추천번호세트
*/

CREATE TABLE IF NOT EXISTS recommend_sets (
	id TEXT primary key,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	
	round_id INTEGER, -- lotto_round.id
	
	algorithm_version TEXT,
	
	score REAL, -- 추천 신뢰도
	generated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 생성일

  FOREIGN KEY(round_id) REFERENCES (lotto_rounds.id)
);