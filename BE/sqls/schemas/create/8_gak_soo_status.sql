/*
* 각 수의 상태
*/
CREATE TABLE lotto_number_status (
	 id text --uuid
	 number int -- 숫자
	 first_count -- 1등 등장
	 second_count -- 2등 등장
   third_count -- 3등 등장
	 fourth_count -- 4등 등장
   fifth_count -- 5등 등장
	 first_probability -- 1등 등장 확률
   second_probability -- 2등 등장 확률
	 created_at -- 생성일
	 updated_at -- 업데이트일
”);