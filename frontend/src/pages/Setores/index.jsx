import { useEffect, useState } from "react";
import Layout from "../../components/Layout";
import api from "../../services/api";
import "./style.css";

const estadoInicial = {
  nome: "",
  local: "",
};

export default function Setores() {
  const [setores, setSetores] = useState([]);
  const [form, setForm] = useState(estadoInicial);
  const [editando, setEditando] = useState(null);

  useEffect(() => {
    carregarDados();
  }, []);

  async function carregarDados() {
    const { data } = await api.get("/setores");

    console.log("SETORES:", data);

    setSetores(data);
  }

  function alterarCampo(e) {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  }

  async function salvar(e) {
    e.preventDefault();

    if (editando) {
      await api.put(`/setores/${editando}`, form);
    } else {
      await api.post("/setores", form);
    }

    setEditando(null);
    setForm(estadoInicial);

    carregarDados();
  }

  function editar(setor) {
    setEditando(setor.ID);

    setForm({
      nome: setor.Nome,
      local: setor.Local,
    });
  }

  async function excluir(id) {
    if (!window.confirm("Excluir setor?")) return;

    await api.delete(`/setores/${id}`);

    carregarDados();
  }

  return (
    <Layout>
      <div className="pagina-setores">
        <h2>Setores</h2>

        <form className="form-setor" onSubmit={salvar}>
          <input
            name="nome"
            placeholder="Nome"
            value={form.nome}
            onChange={alterarCampo}
          />

          <input
            name="local"
            placeholder="Local"
            value={form.local}
            onChange={alterarCampo}
          />

          <button>{editando ? "Atualizar" : "Salvar"}</button>
        </form>

        <table>
          <thead>
            <tr>
              <th>Nome</th>
              <th>Local</th>
              <th>Ações</th>
            </tr>
          </thead>

          <tbody>
            {setores.map((setor) => (
              <tr key={setor.ID}>
                <td>{setor.Nome}</td>
                <td>{setor.Local}</td>

                <td>
                  <button onClick={() => editar(setor)}>Editar</button>

                  <button onClick={() => excluir(setor.ID)}>Excluir</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </Layout>
  );
}
