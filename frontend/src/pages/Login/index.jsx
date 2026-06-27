import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../../services/api";

import "./style.css";

export default function Login() {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);

  async function fazerLogin(e) {
    e.preventDefault();

    try {
      setLoading(true);

      const { data } = await api.post("/login", {
        email,
        password,
      });

      localStorage.setItem("@SCI:token", data.token);

      navigate("/bens");
    } catch (error) {
      if (error.response?.data) {
        alert(error.response.data);
      } else {
        alert("Erro ao realizar login.");
      }
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="login-container">
      <form className="login-box" onSubmit={fazerLogin}>
        <h1>Sistema Patrimonial</h1>

        <input
          type="email"
          placeholder="E-mail"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Senha"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        <button disabled={loading}>{loading ? "Entrando..." : "Entrar"}</button>
      </form>
    </div>
  );
}
