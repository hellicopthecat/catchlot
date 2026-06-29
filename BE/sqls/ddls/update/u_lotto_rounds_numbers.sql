UPDATE lotto_rounds_numbers
SET
  updated_at = CURRENT_TIMESTAMP,
  round_id = ?,
  goal_number = ?
WHERE
  id = ?;