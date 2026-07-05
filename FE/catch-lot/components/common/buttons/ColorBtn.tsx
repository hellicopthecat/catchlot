import {ReactNode} from "react";

export default function ColorButton({children}: {children: ReactNode}) {
  return (
    <button className="bg-primary rounded-md border border-(--fern-hair-light) p-2 hover:bg-(--fern-hair-dark) transition-colors duration-200 ease-in-out">
      {children}
    </button>
  );
}
