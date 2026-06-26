import Sidebar from "../Sidebar";
import Navbar from "../Navbar";

export default function Layout({ titulo, children }) {
  return (
    <div className="page">
      <Sidebar />

      <main className="content">
        <Navbar titulo={titulo} />

        {children}
      </main>
    </div>
  );
}