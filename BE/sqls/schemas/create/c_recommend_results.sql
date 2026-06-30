CREATE TABLE IF NOT EXISTS recommend_results (
	id TEXT primary key,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

	recommend_set_id TEXT not null UNIQUE,

	matched_count INTEGER,
	bonus_match INTEGER,
	rank INTEGER,

  FOREIGN KEY(recommend_set_id) REFERENCES recommend_sets(id)
);