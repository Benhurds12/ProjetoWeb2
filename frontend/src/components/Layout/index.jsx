import Sidebar from "../Sidebar";
import Navbar from "../Navbar";

import "./style.css";

export default function Layout({ children }) {
  return (
    <>
      <Sidebar />
      <Navbar />

      <main className="content">{children}</main>
    </>
  );
}
