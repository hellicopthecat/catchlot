SELECT
  utn.pick_number,
  utn.pick_number_id
FROM
  users_tickets_numbers utn
WHERE
  utn.ticket_id = ?;
