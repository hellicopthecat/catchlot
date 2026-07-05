import {ReactNode} from "react";

export default function BorderBtn({children}: {children: ReactNode}) {
  return (
    <button className="border border-(--stark-hair) rounded-md p-2 hover:border-(--stark-hair-dark) transition-colors duration-200 ease-in-out">
      {children}
    </button>
  );
}
