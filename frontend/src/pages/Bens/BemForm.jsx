export default function BemForm() {

    return (

        <div className="card">

            <h2>Novo Bem</h2>

            <form className="form">

                <input
                    type="text"
                    placeholder="Nome do Bem"
                />

                <input
                    type="text"
                    placeholder="Número Patrimonial"
                />

                <select>

                    <option>Selecione o setor</option>

                </select>

                <select>

                    <option>Status</option>

                    <option>Disponível</option>

                    <option>Em uso</option>

                    <option>Manutenção</option>

                </select>

                <button>

                    Salvar

                </button>

            </form>

        </div>

    );

}