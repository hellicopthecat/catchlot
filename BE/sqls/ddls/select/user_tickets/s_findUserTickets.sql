SELECT
  ut.id,
  ut.created_at,
  ut.updated_at,
  ut.user_id,
  ut.round_id,
  ut.rank,
  ut.bonus_match,
  ut.checked
FROM
  users_tickets ut
JOIN users u
  on ut.user_id = u.id
WHERE 
  u.email = ?
ORDER BY
  ut.created_at DESC
LIMIT ? OFFSET ?;
