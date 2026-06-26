/*
* 공식 당첨번호 테이블
*/
CREATE TABLE lotter_round_number (
	id INTEGER PRIMARY KEY,
	round_id TEXT, -- 로또회차별테이블 id
	goal_number INTEGER, -- 당첨번호

  FOREIGN KEY (round_id) REFERENCES (lotto_round.id)
);