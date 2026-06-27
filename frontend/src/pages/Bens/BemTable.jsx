export default function BemTable(){

    const dados=[

        {
            id:1,
            nome:"Notebook Dell",
            patrimonio:"0001",
            setor:"TI",
            status:"Em uso"
        },

        {
            id:2,
            nome:"Monitor LG",
            patrimonio:"0002",
            setor:"RH",
            status:"Disponível"
        }

    ];

    return(

        <div className="card">

            <table>

                <thead>

                    <tr>

                        <th>Nome</th>

                        <th>Patrimônio</th>

                        <th>Setor</th>

                        <th>Status</th>

                        <th>Ações</th>

                    </tr>

                </thead>

                <tbody>

                    {

                        dados.map((bem)=>(

                            <tr key={bem.id}>

                                <td>{bem.nome}</td>

                                <td>{bem.patrimonio}</td>

                                <td>{bem.setor}</td>

                                <td>{bem.status}</td>

                                <td>

                                    <button>

                                        Editar

                                    </button>

                                    {" "}

                                    <button>

                                        Excluir

                                    </button>

                                </td>

                            </tr>

                        ))

                    }

                </tbody>

            </table>

        </div>

    );

}