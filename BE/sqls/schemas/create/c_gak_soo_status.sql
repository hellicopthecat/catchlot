/*
* 각 수의 상태
*/
CREATE TABLE IF NOT EXISTS gak_soo_status (
	id TEXT PRIMARY KEY, --uuid
	created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

	soo_id TEXT, -- 숫자

	appear_count INTEGER NOT NULL DEFAULT 0, -- 등장

	bonus_count INTEGER NOT NULL DEFAULT 0,

	first_probability REAL NOT NULL DEFAULT 0, -- 1등 등장 확률
	second_probability REAL NOT NULL DEFAULT 0, -- 2등 등장 확률
	third_probability REAL NOT NULL DEFAULT 0, -- 2등 등장 확률
	fourth_probability REAL NOT NULL DEFAULT 0, -- 2등 등장 확률
	fifth_probability REAL NOT NULL DEFAULT 0, -- 2등 등장 확률

	FOREIGN KEY (soo_id) REFERENCES gak_soo(id)
);