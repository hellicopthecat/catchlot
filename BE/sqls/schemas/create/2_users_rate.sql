CREATE TABLE IF NOT EXISTS users_rate_count (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id TEXT NOT NULL, -- ref user.id
	count_first INTEGER,
	count_second INTEGER,
	count_third INTEGER,
	count_fourth INTEGER,
	count_fifth INTEGER,	
	rate_win_first REAL,
	rate_win_second REAL,
	rate_win_third REAL,
	rate_win_fourth REAL,
	rate_win_fifth REAL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);