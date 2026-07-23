/*
* 유저가 구매한 티켓
*/

CREATE TABLE IF NOT EXISTS users_tickets  (
  id INTEGER PRIMARY KEY NOT NULL,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

  user_id TEXT NOT NULL,
  round_id INTEGER NOT NULL,

  rank INTEGER NOT NULL,

  bonus_match INTEGER DEFAULT false, -- BOOLEAN -- 보너스 번호 맞았는지
  checked INTEGER DEFAULT false, -- BOOLEAN -- 당첨확인체크

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (round_id) REFERENCES lotto_rounds(round_no)
);