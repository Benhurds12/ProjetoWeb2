-- ============================================
-- USUÁRIOS
-- Senha de todos: 123456
-- ============================================

INSERT INTO users (nome, email, cpf, password) VALUES
('Administrador', 'admin@projetoweb2.com', '11111111111', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy'),
('João Silva', 'joao@projetoweb2.com', '22222222222', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy'),
('Maria Souza', 'maria@projetoweb2.com', '33333333333', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy'),
('Carlos Oliveira', 'carlos@projetoweb2.com', '44444444444', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy'),
('Ana Lima', 'ana@projetoweb2.com', '55555555555', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy');


-- ============================================
-- SETORES
-- ============================================

INSERT INTO setores (nome, local) VALUES
('Tecnologia', 'Bloco A'),
('Financeiro', 'Bloco B'),
('Recursos Humanos', 'Bloco C'),
('Diretoria', 'Bloco D'),
('Almoxarifado', 'Bloco E');


-- ============================================
-- FABRICANTES
-- ============================================

INSERT INTO fabricantes (nome, cnpj) VALUES
('Dell', '11111111000101'),
('HP', '22222222000102'),
('Lenovo', '33333333000103'),
('Samsung', '44444444000104'),
('LG', '55555555000105');


-- ============================================
-- FORNECEDORES
-- ============================================

INSERT INTO fornecedores (nome, cnpj, contato) VALUES
('KaBuM', '66666666000106', 'contato@kabum.com'),
('Pichau', '77777777000107', 'vendas@pichau.com'),
('Terabyte', '88888888000108', 'suporte@terabyte.com'),
('Amazon Brasil', '99999999000109', 'comercial@amazon.com'),
('Magazine Luiza', '12345678000110', 'corporativo@magalu.com');


-- ============================================
-- BENS
-- ============================================

INSERT INTO bens
(id, nome, status, tipo, setor_id, fornecedor_id, fabricante_id)
VALUES

(gen_random_uuid(), 'Notebook Dell Latitude', 'EM_USO', 'Notebook', 1, 1, 1),

(gen_random_uuid(), 'Impressora HP LaserJet', 'OCIOSO', 'Impressora', 2, 2, 2),

(gen_random_uuid(), 'Notebook Lenovo ThinkPad', 'MANUTENCAO', 'Notebook', 3, 3, 3),

(gen_random_uuid(), 'Monitor Samsung 27"', 'EM_USO', 'Monitor', 4, 4, 4),

(gen_random_uuid(), 'Smart TV LG 55"', 'EM_USO', 'Televisor', 5, 5, 5);