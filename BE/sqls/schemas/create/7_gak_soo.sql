/*
* 로또의 각 번호
*/
CREATE TABLE gak_soo (
	id TEXT PRIMARY KEY, -- uuid
	gak_num INTEGER, -- (1~45)
	-- 1등 출현 확률 (매주 일요일 아침batch or server_calc)
  -- 2등 출현 확률 (매주 일요일 아침batch or server_calc )
  -- 3등 출현 확률 (매주 일요일 아침batch or server_calc)
  -- 4등 출현 확률 (매주 일요일 아침batch or server_calc)
  -- 5등 출현 확률 (매주 일요일 아침batch or server_calc)
);