"use client";
import {ChangeEvent, MouseEvent, useState} from "react";
import Card from "../common/Card";
import {ITicketField} from "@/types/ticket.types";

export default function RegistTicketModal() {
  const ticketFormat = {
    round_id: "",
    rank: "",
    numbers: ["", "", "", "", "", ""],
  };
  const [ticketField, setTicketField] = useState<ITicketField[]>([
    {round_id: "", rank: "", numbers: ["", "", "", "", "", ""]},
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
      prev.map((item, idx) =>
        idx === i
          ? {
              ...item,
              [key]: val,
            }
          : item,
      ),
    );
  };
  const onChangeNumber = (
    e: ChangeEvent<HTMLInputElement>,
    i: number,
    t_index: number,
  ) => {
    const val = e.target.value;
    setTicketField((prev) =>
      prev.map((item, idx) => {
        if (idx !== i) return item;
        const numbers = [...item.numbers];
        numbers[t_index] = val;
        return {
          ...item,
          numbers,
        };
      }),
    );
  };
  const onSubmit = (e: MouseEvent) => {
    e.preventDefault();
    console.log(ticketField);
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
                value={val.round_id}
                onChange={(e) => onChange(e, i, "round_id")}
              />
            </div>
            <div>
              <label htmlFor="rank">등수</label>
              <input
                id="rank"
                type="text"
                placeholder="등수를 입력해주세요."
                value={val.rank}
                onChange={(e) => onChange(e, i, "rank")}
              />
            </div>
            <div>
              <label htmlFor="">1</label>
              <input
                type="text"
                placeholder="1번째 수를 작성하세요."
                value={val.numbers[0]}
                onChange={(e) => onChangeNumber(e, i, 0)}
              />
            </div>
            <div>
              <label htmlFor="">2</label>
              <input
                type="text"
                placeholder="2번째 수를 작성하세요."
                value={val.numbers[1]}
                onChange={(e) => onChangeNumber(e, i, 1)}
              />
            </div>
            <div>
              <label htmlFor="">3</label>
              <input
                type="text"
                placeholder="3번째 수를 작성하세요."
                value={val.numbers[2]}
                onChange={(e) => onChangeNumber(e, i, 2)}
              />
            </div>
            <div>
              <label htmlFor="">4</label>
              <input
                type="text"
                placeholder="4번째 수를 작성하세요."
                value={val.numbers[3]}
                onChange={(e) => onChangeNumber(e, i, 3)}
              />
            </div>
            <div>
              <label htmlFor="">5</label>
              <input
                type="text"
                placeholder="5번째 수를 작성하세요."
                value={val.numbers[4]}
                onChange={(e) => onChangeNumber(e, i, 4)}
              />
            </div>
            <div>
              <label htmlFor="">6</label>
              <input
                type="text"
                placeholder="6번째 수를 작성하세요."
                value={val.numbers[5]}
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
