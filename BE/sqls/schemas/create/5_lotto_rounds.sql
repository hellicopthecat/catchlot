/*
* 공식 로또 회차 별 테이블
*/

CREATE TABLE IF NOT EXISTS lotto_round (
	id INTEGER PRIMARY KEY
	created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
  
	round_no INTEGER NOT NULL DEFAULT 0, -- 회차
	draw_date TEXT NOT NULL DEFAULT '', -- 추첨일
	bonus_number INTEGER NOT NULL DEFAULT 0, -- 보너스 번호

  FOREIGN KEY(round_no) REFERENCES (lotto_round_number.id)
);