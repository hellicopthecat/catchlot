UPDATE users_tickets 
SET
  rank = ?,
  bonus_match = ?,
  checked = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE 
  round_id = ?;