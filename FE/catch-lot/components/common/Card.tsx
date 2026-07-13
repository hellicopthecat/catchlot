import {ReactNode} from "react";

export default function Card({children}: {children: ReactNode}) {
  return (
    <section className="flex flex-col rounded-xl border border-(--fern-eye) p-5">
      {children}
    </section>
  );
}
