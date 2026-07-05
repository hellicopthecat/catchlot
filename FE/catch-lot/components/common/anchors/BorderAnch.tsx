import {ReactNode} from "react";

export default function BorderAnch({children}: {children: ReactNode}) {
  return (
    <a
      href="#"
      className="text-center border border-(--stark-hair) rounded-md p-2 hover:border-(--stark-hair-dark) transition-colors duration-200 ease-in-out"
    >
      {children}
    </a>
  );
}
