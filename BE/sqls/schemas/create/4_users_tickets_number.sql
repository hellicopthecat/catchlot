CREATE TABLE IF NOT EXISTS users_ticket_number (
	id INTEGER PRIMARY KEY NOT NULL,
  created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,

	ticket_id INTEGER NOT NULL,
  pick_number_id TEXT NOT NULL,
	pick_number INTEGER NOT NULL,

  FOREIGN KEY (ticket_id) REFERENCES (users_tickets.id),
  FOREIGN KEY (pick_number_id) REFERENCES (gak_soo.id)
)