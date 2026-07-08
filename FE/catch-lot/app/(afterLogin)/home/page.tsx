"use client";

import {useRouter} from "next/navigation";

export default function Page() {
  const router = useRouter();
  const logOut = async () => {
    try {
      const res = await fetch(`http://localhost:4000/api/user/logout`, {
        method: "GET",
        credentials: "include",
      });
      const result = await res.json();
      if (!res.ok) {
        console.log("Not OK :: ", result);
      }
      router.push("/");
    } catch (error) {
      const err = error as Error;
      console.error("Logout Error :: ", err);
    }
  };
  return (
    <div>
      <button onClick={logOut}>logout</button>
    </div>
  );
}
