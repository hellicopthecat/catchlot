CREATE TABLE IF NOT EXISTS recommend_set_numbers (
	id INTEGER primary key,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

	recommend_set_id INTEGER not null,

	recommend_number INTEGER not null,
	weights REAL, -- 등장확률
	rank INTEGER,

  FOREIGN KEY (recommend_set_id) REFERENCES (recommend_sets.id)
)