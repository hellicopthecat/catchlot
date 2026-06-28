UPDATE users
SET 
  nickname = ?, 
  updated_at = CURRENT_TIMESTAMP
WHERE 
  id = ?;