UPDATE users
SET 
  updated_at = CURRENT_TIMESTAMP,
  nickname = CASE WHEN :nickname IS NOT NULL THEN :nickname ELSE nickname END, 
  refresh_token = CASE WHEN :refresh_token IS NOT NULL THEN :refresh_token ELSE refresh_token END
WHERE 
  email = ?;