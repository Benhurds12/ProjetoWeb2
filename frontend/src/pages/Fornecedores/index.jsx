import { useEffect, useState } from "react";
import Layout from "../../components/Layout";
import api from "../../services/api";
import "./style.css";

const estadoInicial = {
  nome: "",
  cnpj: "",
  contato: "",
};

export default function Fornecedores() {
  const [fornecedores, setFornecedores] = useState([]);
  const [form, setForm] = useState(estadoInicial);
  const [editando, setEditando] = useState(null);

  useEffect(() => {
    carregarDados();
  }, []);

  async function carregarDados() {
    const { data } = await api.get("/fornecedores");

    console.log("FORNECEDORES:", data);

    setFornecedores(data);
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
      await api.put(`/fornecedores/${editando}`, form);
    } else {
      await api.post("/fornecedores", form);
    }

    setEditando(null);
    setForm(estadoInicial);

    carregarDados();
  }

  function editar(fornecedor) {
    setEditando(fornecedor.ID);

    setForm({
      nome: fornecedor.Nome,
      cnpj: fornecedor.Cnpj,
      contato: fornecedor.Contato,
    });
  }

  async function excluir(id) {
    if (!window.confirm("Excluir fornecedor?")) return;

    await api.delete(`/fornecedores/${id}`);

    carregarDados();
  }

  return (
    <Layout>
      <div className="pagina-fornecedores">
        <h2>Fornecedores</h2>

        <form className="form-fornecedor" onSubmit={salvar}>
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

          <input
            name="contato"
            placeholder="Contato"
            value={form.contato}
            onChange={alterarCampo}
          />

          <button>{editando ? "Atualizar" : "Salvar"}</button>
        </form>

        <table>
          <thead>
            <tr>
              <th>Nome</th>
              <th>CNPJ</th>
              <th>Contato</th>
              <th>Ações</th>
            </tr>
          </thead>

          <tbody>
            {fornecedores.map((fornecedor) => (
              <tr key={fornecedor.ID}>
                <td>{fornecedor.Nome}</td>
                <td>{fornecedor.Cnpj}</td>
                <td>{fornecedor.Contato}</td>

                <td>
                  <button onClick={() => editar(fornecedor)}>Editar</button>

                  <button onClick={() => excluir(fornecedor.ID)}>
                    Excluir
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </Layout>
  );
}
