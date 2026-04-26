import { Outlet } from "react-router-dom";

import Header from "../components/Header";

export default function AppLayout() {
  return (
    <>
      <Header />
      <a
        className="text-[#000] text-[10px] bg-gradient-to-r from-slate-200 "
        href="https://www.flaticon.com/free-icons/shrink"
        title="shrink icons"
      >
        Shrink icons created by Freepik - Flaticon
      </a>
      <Outlet />
    </>
  );
}
