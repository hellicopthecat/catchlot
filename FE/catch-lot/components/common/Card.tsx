import {ReactNode} from "react";

export default function Card({children}: {children: ReactNode}) {
  return (
    <div className="flex flex-col rounded-xl border border-(--color-border) p-5">
      {children}
    </div>
  );
}
