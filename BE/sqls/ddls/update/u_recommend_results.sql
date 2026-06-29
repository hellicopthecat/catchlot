UPDATE recommend_results
SET
  updated_at = CURRENT_TIMESTAMP,
  matched_count = ?,
  bonus_match = ?,
  rank = ?,
WHERE
  recommend_set_id = ?;