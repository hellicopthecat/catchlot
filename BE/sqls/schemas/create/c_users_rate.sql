/*
* 유저의 개인 확률
*/
CREATE TABLE IF NOT EXISTS users_rate_count (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	
	user_id TEXT NOT NULL, -- ref user.id

	count_first INTEGER NOT NULL DEFAULT 0,
	count_second INTEGER NOT NULL DEFAULT 0,
	count_third INTEGER NOT NULL DEFAULT 0,
	count_fourth INTEGER NOT NULL DEFAULT 0,
	count_fifth INTEGER NOT NULL DEFAULT 0,	
	rate_win_first REAL NOT NULL DEFAULT 0,
	rate_win_second REAL NOT NULL DEFAULT 0,
	rate_win_third REAL NOT NULL DEFAULT 0,
	rate_win_fourth REAL NOT NULL DEFAULT 0,
	rate_win_fifth REAL NOT NULL DEFAULT 0,
	
  FOREIGN KEY (user_id) REFERENCES users(id)
);