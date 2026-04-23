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

CREATE TABLE bens (
    id UUID PRIMARY KEY, 
    nome TEXT NOT NULL,
    status TEXT DEFAULT 'OCIOSO',
    tipo TEXT NOT NULL,
    setor_id INT REFERENCES setores(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT NOW()
);