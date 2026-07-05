import BorderAnch from "@/components/common/anchors/BorderAnch";
import ColorAnch from "@/components/common/anchors/ColorAnchBtn";
import BorderBtn from "@/components/common/buttons/BorderBtn";
import ColorButton from "@/components/common/buttons/ColorBtn";
import Card from "@/components/common/Card";

export default function Login() {
  return (
    <div className="">
      <Card>
        <h1>CATCH LOT</h1>
        <ColorButton>Zoltraak</ColorButton>
        <BorderBtn>Strak</BorderBtn>
        <ColorAnch>Fern</ColorAnch>
        <BorderAnch>Zahen</BorderAnch>
      </Card>
    </div>
  );
}
