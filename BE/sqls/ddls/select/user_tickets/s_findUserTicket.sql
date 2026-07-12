SELECT
  id,
  created_at,
  updated_at,
  user_id,
  round_id,
  rank,
  bonus_match,
  checked
FROM
  users_tickets
WHERE
  id = ?;
