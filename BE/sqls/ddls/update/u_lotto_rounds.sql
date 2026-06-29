UPDATE lotto_rounds
SET 
  updated_at = CURRENT_TIMESTAMP,
  draw_date = ?,
  bonus_number = ?
WHERE 
  id = ?;