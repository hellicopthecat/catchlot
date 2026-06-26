/*
* 각 수의 상태
*/
CREATE TABLE IF NOT EXISTS gak_soo_status (
	id TEXT PRIMARY KEY, --uuid
	created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

	soo_id INTEGER, -- 숫자

	first_count INTEGER, -- 1등 등장
	second_count INTEGER, -- 2등 등장
	third_count INTEGER, -- 3등 등장
	fourth_count INTEGER, -- 4등 등장
	fifth_count INTEGER, -- 5등 등장
	first_probability REAL, -- 1등 등장 확률
	second_probability REAL, -- 2등 등장 확률

	FOREIGN KEY (soo_id) REFERENCES (gak_soo.id)
);