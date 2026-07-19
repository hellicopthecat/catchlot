UPDATE gak_soo_status
SET
  appear_count = appear_count + 1,
  bonus_count = bonus_count + CASE WHEN ? THEN 1 ELSE 0 END,
  updated_at = CURRENT_TIMESTAMP
WHERE 
  soo_id = ?;