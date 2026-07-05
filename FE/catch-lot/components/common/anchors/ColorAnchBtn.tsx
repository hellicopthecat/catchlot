import {ReactNode} from "react";

export default function ColorAnch({children}: {children: ReactNode}) {
  return (
    <a
      href="#"
      className="text-center bg-primary rounded-md border border-(--fern-hair-light) p-2 hover:bg-(--fern-hair-dark) transition-colors duration-200 ease-in-out"
    >
      {children}
    </a>
  );
}
