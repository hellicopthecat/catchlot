UPDATE users
SET 
  updated_at = CURRENT_TIMESTAMP,
  refresh_token = null
WHERE 
  email = ?;