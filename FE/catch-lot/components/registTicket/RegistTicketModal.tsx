import Card from "../common/Card";

export default function RegistTicketModal() {
  return (
    <Card>
      <button>+</button>
      <button>-</button>
      <form>
        <fieldset>
          <legend></legend>
          <div>
            <label htmlFor=""></label>
            <input type="text" />
          </div>
          <div>
            <label htmlFor=""></label>
            <input type="text" />
          </div>
          <div>
            <label htmlFor=""></label>
            <input type="text" />
          </div>
          <div>
            <label htmlFor=""></label>
            <input type="text" />
          </div>
          <div>
            <label htmlFor=""></label>
            <input type="text" />
          </div>
          <button>초기화</button>
        </fieldset>
        <button>취소</button>
        <button>확인</button>
      </form>
    </Card>
  );
}
