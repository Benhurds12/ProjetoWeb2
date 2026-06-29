import { useEffect, useState } from "react";
import Layout from "../../components/Layout";
import api from "../../services/api";
import "./style.css";

const estadoInicial = {
  nome: "",
  cnpj: "",
};

export default function Fabricantes() {
  const [fabricantes, setFabricantes] = useState([]);
  const [form, setForm] = useState(estadoInicial);
  const [editando, setEditando] = useState(null);

  useEffect(() => {
    carregarDados();
  }, []);

  async function carregarDados() {
    const { data } = await api.get("/fabricantes");

    console.log("FABRICANTES:", data);

    setFabricantes(data);
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
      await api.put(`/fabricantes/${editando}`, form);
    } else {
      await api.post("/fabricantes", form);
    }

    setEditando(null);
    setForm(estadoInicial);

    carregarDados();
  }

  function editar(fabricante) {
    setEditando(fabricante.ID);

    setForm({
      nome: fabricante.Nome,
      cnpj: fabricante.Cnpj,
    });
  }

  async function excluir(id) {
    if (!window.confirm("Excluir fabricante?")) return;

    await api.delete(`/fabricantes/${id}`);

    carregarDados();
  }

  return (
    <Layout>
      <h2>Fabricantes</h2>

      <form className="form-fabricante" onSubmit={salvar}>
        <input
          name="nome"
          placeholder="Nome"
          value={form.nome}
          onChange={alterarCampo}
        />

        <input
          name="cnpj"
          placeholder="CNPJ"
          value={form.cnpj}
          onChange={alterarCampo}
        />

        <button>
          {editando ? "Atualizar" : "Salvar"}
        </button>
      </form>

      <table>
        <thead>
          <tr>
            <th>Nome</th>
            <th>CNPJ</th>
            <th>Ações</th>
          </tr>
        </thead>

        <tbody>
          {fabricantes.map((fabricante) => (
            <tr key={fabricante.ID}>
              <td>{fabricante.Nome}</td>
              <td>{fabricante.Cnpj}</td>

              <td>
                <button onClick={() => editar(fabricante)}>
                  Editar
                </button>

                <button onClick={() => excluir(fabricante.ID)}>
                  Excluir
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </Layout>
  );
}