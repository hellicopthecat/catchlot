INSERT INTO users_tickets (
  user_id,
  round_id,
  rank
)
VALUES (?,?,?)
RETURNING id;