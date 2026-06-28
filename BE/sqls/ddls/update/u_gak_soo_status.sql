UPDATE gak_soo_status
SET
  first_count = ?,
  second_count = ?,
  third_count = ?,
  fourth_count = ?,
  fifth_count = ?,
  first_probability = ?,
  second_probability = ?,
  updated_at = CURRENT_TIMESTAMP
WHERE 
  soo_id = ?;