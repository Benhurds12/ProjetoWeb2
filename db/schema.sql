CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    cpf TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE setores (
    id SERIAL PRIMARY KEY,
    nome TEXT UNIQUE NOT NULL,
    local TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);



CREATE TABLE fornecedores (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    cnpj TEXT UNIQUE NOT NULL,
    contato TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE fabricantes (
    id SERIAL PRIMARY KEY,
    nome TEXT NOT NULL,
    cnpj TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);


--A tabela bens precisa ter uma chave estrangeira para os setores, fornecedores e fabricantes
--Então ela precisa existir depois das tabelas acimas
CREATE TABLE bens (
    id UUID PRIMARY KEY, 
    nome TEXT NOT NULL,
    status TEXT DEFAULT 'OCIOSO',
    tipo TEXT NOT NULL,

    setor_id INT REFERENCES setores(id) ON DELETE SET NULL,
    fornecedor_id INT REFERENCES fornecedores(id) ON DELETE SET NULL,
    fabricante_id INT REFERENCES fabricantes(id) ON DELETE SET NULL,

    created_at TIMESTAMP DEFAULT NOW()
);