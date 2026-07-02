import "./Sidebar.css";

import {
  FaBoxes,
  FaBuilding,
  FaIndustry,
  FaTruck,
  FaSignOutAlt,
} from "react-icons/fa";

import { Link, useNavigate } from "react-router-dom";
import api from "../../services/api";

export default function Sidebar() {
  const navigate = useNavigate();

  async function logout() {
    try {
      const refreshToken = localStorage.getItem("@SCI:refreshToken");

      if (refreshToken) {
        await api.post("/logout", {
          refresh_token: refreshToken,
        });
      }
    } catch (error) {
      console.error("Erro ao realizar logout:", error);
    } finally {
      localStorage.removeItem("@SCI:token");
      localStorage.removeItem("@SCI:refreshToken");

      navigate("/");
    }
  }

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

      <button onClick={logout}>
        <FaSignOutAlt />
        <span>Sair</span>
      </button>
    </aside>
  );
}
