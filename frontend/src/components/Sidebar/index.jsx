import "./Sidebar.css";

import {
  FaBoxes,
  FaBuilding,
  FaIndustry,
  FaTruck,
  FaSignOutAlt,
} from "react-icons/fa";

import { Link } from "react-router-dom";

export default function Sidebar() {
  return (
    <aside className="sidebar">
      <h2>Patrimônio</h2>

      <nav>
        <Link to="/bens">
          <FaBoxes />
          <span>Bens</span>
        </Link>

        <Link to="/setores">
          <FaBuilding />
          <span>Setores</span>
        </Link>

        <Link to="/fabricantes">
          <FaIndustry />
          <span>Fabricantes</span>
        </Link>

        <Link to="/fornecedores">
          <FaTruck />
          <span>Fornecedores</span>
        </Link>
      </nav>

      <button>
        <FaSignOutAlt />

        <span>Sair</span>
      </button>
    </aside>
  );
}
