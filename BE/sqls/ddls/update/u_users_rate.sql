UPDATE users_rate
SET
  created_at = CURRENT_TIMESTAMP,
  count_first = ?,
  count_second = ?,
  count_third = ?,
  count_fourth = ?,
  count_fifth = ?,
  rate_win_first = ?,
  rate_win_second = ?,
  rate_win_third = ?,
  rate_win_fourth = ?,
  rate_win_fifth = ?
WHERE
  id = ?;