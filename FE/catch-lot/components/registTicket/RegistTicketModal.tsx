"use client";
import {ChangeEvent, MouseEvent, useState} from "react";
import Card from "../common/Card";
import {ITicketField} from "@/types/ticket.types";
import {api} from "@/constants/constants";

export default function RegistTicketModal() {
  const ticketFormat = {
    round_id: null,
    rank: null,
    numbers: [null, null, null, null, null, null],
  };
  const [ticketField, setTicketField] = useState<ITicketField[]>([
    {round_id: null, rank: null, numbers: [null, null, null, null, null, null]},
  ]);

  const increaseFieldArr = () => {
    setTicketField((prev) => [...prev, ticketFormat]);
  };
  const decreaseFieldArr = () => {
    setTicketField((prev) => (prev.length <= 1 ? prev : prev.slice(0, -1)));
  };
  const removeFieldArr = (e: MouseEvent, index: number) => {
    e.preventDefault();
    setTicketField((prev) =>
      prev.length <= 1 ? prev : prev.filter((_, i) => index !== i),
    );
  };
  const resetFieldArr = (e: MouseEvent) => {
    e.preventDefault();
    setTicketField([ticketFormat]);
  };
  const onChange = (
    e: ChangeEvent<HTMLInputElement>,
    i: number,
    key: string,
  ) => {
    const val = e.target.value;
    setTicketField((prev) =>
      prev.map((item, idx) => {
        return idx === i
          ? {
              ...item,
              [key]: Number(val),
            }
          : item;
      }),
    );
  };
  const onChangeSelect = (
    e: ChangeEvent<HTMLSelectElement>,
    i: number,
    key: string,
  ) => {
    const val = e.target.value;
    setTicketField((prev) =>
      prev.map((item, idx) => {
        return idx === i
          ? {
              ...item,
              [key]: Number(val),
            }
          : item;
      }),
    );
  };

  const onChangeNumber = (
    e: ChangeEvent<HTMLInputElement>,
    i: number,
    t_index: number,
  ) => {
    const val = e.target.value;
    if (val === "") {
      setTicketField((prev) =>
        prev.map((item, idx) => {
          if (idx !== i) return item;
          const numbers = [...item.numbers];
          numbers[t_index] = null;
          return {
            ...item,
            numbers,
          };
        }),
      );
    }
    if (Number.isNaN(Number(val)) || Number(val) < 1 || Number(val) > 45) {
      return;
    }
    setTicketField((prev) =>
      prev.map((item, idx) => {
        if (idx !== i) return item;
        const numbers = [...item.numbers];
        numbers[t_index] = val === "" ? null : Number(val);
        return {
          ...item,
          numbers,
        };
      }),
    );
  };

  const onSubmit = async (e: MouseEvent) => {
    e.preventDefault();
    const res = await fetch(`${api}/ticket/new`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(ticketField),
    });
    const data = await res.json();
    if (!res.ok) {
      // {status: false, msg: '올바른 요청이 아닙니다.', data: null}
      resetFieldArr(e);
    }
    // {status: true, msg: '등록성공했습니다.', data: null}
    resetFieldArr(e);
    console.log(data);
  };
  return (
    <Card>
      <button onClick={increaseFieldArr}>+</button>
      <button onClick={decreaseFieldArr}>-</button>
      <form>
        {ticketField.map((val, i) => (
          <fieldset key={i}>
            <legend>ticket field</legend>
            <div>
              <label htmlFor="round_id">회차</label>
              <input
                id="round_id"
                type="text"
                placeholder="회차를 입력해주세요."
                value={val.round_id ?? ""}
                onChange={(e) => onChange(e, i, "round_id")}
              />
            </div>
            <div>
              <label htmlFor="rank">등수</label>
              <select
                name=""
                id="rank"
                onChange={(e) => onChangeSelect(e, i, "rank")}
              >
                <option value="-">등수를 선택해주세요.</option>
                <option value="0">꽝</option>
                <option value="5">5등</option>
                <option value="4">4등</option>
                <option value="3">3등</option>
                <option value="2">2등</option>
                <option value="1">1등</option>
              </select>
            </div>
            <div>
              <label htmlFor="">1</label>
              <input
                type="text"
                placeholder="1번째 수를 작성하세요."
                value={val.numbers[0] ?? ""}
                onChange={(e) => onChangeNumber(e, i, 0)}
              />
            </div>
            <div>
              <label htmlFor="">2</label>
              <input
                type="text"
                placeholder="2번째 수를 작성하세요."
                value={val.numbers[1] ?? ""}
                onChange={(e) => onChangeNumber(e, i, 1)}
              />
            </div>
            <div>
              <label htmlFor="">3</label>
              <input
                type="text"
                placeholder="3번째 수를 작성하세요."
                value={val.numbers[2] ?? ""}
                onChange={(e) => onChangeNumber(e, i, 2)}
              />
            </div>
            <div>
              <label htmlFor="">4</label>
              <input
                type="text"
                placeholder="4번째 수를 작성하세요."
                value={val.numbers[3] ?? ""}
                onChange={(e) => onChangeNumber(e, i, 3)}
              />
            </div>
            <div>
              <label htmlFor="">5</label>
              <input
                type="text"
                placeholder="5번째 수를 작성하세요."
                value={val.numbers[4] ?? ""}
                onChange={(e) => onChangeNumber(e, i, 4)}
              />
            </div>
            <div>
              <label htmlFor="">6</label>
              <input
                type="text"
                placeholder="6번째 수를 작성하세요."
                value={val.numbers[5] ?? ""}
                onChange={(e) => onChangeNumber(e, i, 5)}
              />
            </div>
            <button type="button" onClick={(e) => resetFieldArr(e)}>
              초기화
            </button>
            <button type="button" onClick={(e) => removeFieldArr(e, i)}>
              삭제
            </button>
          </fieldset>
        ))}
        <button>취소</button>
        <button type="button" onClick={onSubmit}>
          확인
        </button>
      </form>
    </Card>
  );
}
