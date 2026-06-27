import { useEffect, useState } from "react";
import Layout from "../../components/Layout";
import api from "../../services/api";
import "./style.css";

const estadoInicial = {
  nome: "",
  tipo: "",
  status: "OCIOSO",
  setor_id: "",
  fornecedor_id: "",
  fabricante_id: "",
};

export default function Bens() {
  const [bens, setBens] = useState([]);

  const [setores, setSetores] = useState([]);
  const [fabricantes, setFabricantes] = useState([]);
  const [fornecedores, setFornecedores] = useState([]);

  const [form, setForm] = useState(estadoInicial);

  const [editando, setEditando] = useState(null);

  useEffect(() => {
    carregarDados();
  }, []);

  async function carregarDados() {
    const [bensRes, setoresRes, fabricantesRes, fornecedoresRes] =
      await Promise.all([
        api.get("/bens"),
        api.get("/setores"),
        api.get("/fabricantes"),
        api.get("/fornecedores"),
      ]);

    console.log("BENS:", bensRes.data);

    setBens(bensRes.data);
    setSetores(setoresRes.data);
    setFabricantes(fabricantesRes.data);
    setFornecedores(fornecedoresRes.data);
  }

  function alterarCampo(e) {
    setForm({
      ...form,
      [e.target.name]:
        e.target.value === ""
          ? null
          : e.target.name.includes("_id")
            ? Number(e.target.value)
            : e.target.value,
    });
  }

  async function salvar(e) {
    e.preventDefault();

    if (editando) {
      await api.put(`/bens/${editando}`, form);
    } else {
      await api.post("/bens", form);
    }

    setEditando(null);
    setForm(estadoInicial);
    carregarDados();
  }

  function editar(bem) {
    setEditando(bem.ID);

    setForm({
      nome: bem.Nome,
      tipo: bem.Tipo,
      status: bem.Status.Valid ? bem.Status.String : "OCIOSO",
      setor_id: bem.SetorID.Valid ? bem.SetorID.Int32 : "",
      fornecedor_id: bem.FornecedorID.Valid ? bem.FornecedorID.Int32 : "",
      fabricante_id: bem.FabricanteID.Valid ? bem.FabricanteID.Int32 : "",
    });
  }

  async function excluir(id) {
    if (!window.confirm("Excluir bem?")) return;

    await api.delete(`/bens/${id}`);

    carregarDados();
  }

  return (
    <Layout>
      <h2>Bens</h2>

      <form className="form-bem" onSubmit={salvar}>
        <input
          name="nome"
          placeholder="Nome"
          value={form.nome}
          onChange={alterarCampo}
        />

        <input
          name="tipo"
          placeholder="Tipo"
          value={form.tipo}
          onChange={alterarCampo}
        />

        <select name="status" value={form.status} onChange={alterarCampo}>
          <option value="OCIOSO">OCIOSO</option>
          <option value="EM_USO">EM USO</option>
          <option value="MANUTENCAO">MANUTENÇÃO</option>
        </select>

        <select
          name="setor_id"
          value={form.setor_id ?? ""}
          onChange={alterarCampo}
        >
          <option value="">Setor</option>

          {setores.map((setor) => (
            <option key={setor.ID} value={setor.ID}>
              {setor.Nome}
            </option>
          ))}
        </select>

        <select
          name="fabricante_id"
          value={form.fabricante_id ?? ""}
          onChange={alterarCampo}
        >
          <option value="">Fabricante</option>

          {fabricantes.map((f) => (
            <option key={f.ID} value={f.ID}>
              {f.Nome}
            </option>
          ))}
        </select>

        <select
          name="fornecedor_id"
          value={form.fornecedor_id ?? ""}
          onChange={alterarCampo}
        >
          <option value="">Fornecedor</option>

          {fornecedores.map((f) => (
            <option key={f.ID} value={f.ID}>
              {f.Nome}
            </option>
          ))}
        </select>

        <button>{editando ? "Atualizar" : "Salvar"}</button>
      </form>

      <table>
        <thead>
          <tr>
            <th>Nome</th>
            <th>Tipo</th>
            <th>Status</th>
            <th>Ações</th>
          </tr>
        </thead>

        <tbody>
          {bens.map((bem) => (
            <tr key={bem.ID}>
              <td>{bem.Nome}</td>
              <td>{bem.Tipo}</td>
              <td>{bem.Status.Valid ? bem.Status.String : "-"}</td>
              <td>
                <button onClick={() => editar(bem)}>Editar</button>

                <button onClick={() => excluir(bem.ID)}>Excluir</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </Layout>
  );
}
