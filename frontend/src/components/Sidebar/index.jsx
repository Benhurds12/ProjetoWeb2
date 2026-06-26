import { Link, useLocation } from "react-router-dom";
import "./style.css";

export default function Sidebar() {
  const location = useLocation();

  const menu = [
    { nome: "Bens", rota: "/bens", icone: "📦" },
    { nome: "Setores", rota: "/setores", icone: "🏢" },
    { nome: "Fabricantes", rota: "/fabricantes", icone: "🏭" },
    { nome: "Fornecedores", rota: "/fornecedores", icone: "🚚" },
  ];

  function sair() {
    localStorage.removeItem("@SCI:token");
    window.location.href = "/";
  }

  return (
    <aside className="sidebar">
      <h2>SCI</h2>

      <nav>
        {menu.map((item) => (
          <Link
            key={item.rota}
            to={item.rota}
            className={location.pathname === item.rota ? "active" : ""}
          >
            <span>{item.icone}</span>
            {item.nome}
          </Link>
        ))}
      </nav>

      <button onClick={sair}>Sair</button>
    </aside>
  );
}