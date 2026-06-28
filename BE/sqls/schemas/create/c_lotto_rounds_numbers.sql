/*
* 공식 당첨번호 테이블
*/
CREATE TABLE IF NOT EXISTS lotto_rounds_numbers (
	id INTEGER PRIMARY KEY,
	created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	
	round_id TEXT, -- 로또회차별테이블 id
	goal_number INTEGER, -- 당첨번호

  FOREIGN KEY (round_id) REFERENCES (lotto_round.id)
);