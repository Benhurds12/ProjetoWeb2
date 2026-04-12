# Proposta de Projeto: SCI - Sistema de Controle de Inventário

Este documento apresenta a proposta do Projeto Integrador para a disciplina de **Desenvolvimento de Sistemas Web II** (DIM0547) do IMD/UFRN.

---

## 1. Visão do Produto

**Para** gestores de TI e administradores de patrimônio  
**Que** sofrem com o controle descentralizado em planilhas e a perda de rastreabilidade de ativos  
**O SCI** é uma API REST de controle de inventário e bens  
**Que** centraliza o ciclo de vida do ativo, desde o fornecedor até a atribuição de local e/ou setor  
**Diferente de** planilhas manuais ou ERPs genéricos e complexos  
**Nosso produto** oferece controle estruturado, rastreabilidade completa e integração facilitada via API para emissão de relatórios e prestação de contas.

---

## 2. Definição do MVP (Minimum Viable Product)

### Problema Central
Falta de controle eficiente sobre bens de TI, causando perda de ativos, dificuldade de rastreamento e inconsistência de informações para auditorias.

### Hipótese de Valor
Acreditamos que gestores de TI vão utilizar o sistema para controlar os ativos porque ele centraliza informações e facilita o rastreamento de bens e responsáveis, garantindo eficiência na prestação de contas.

### Funcionalidades Essenciais (no MVP)

| Funcionalidade | Descrição |
|----------------|-------------|
| CRUD de bens | Cadastro, edição, listagem e remoção de equipamentos |
| CRUD de setores | Definir locais/departamentos |
| CRUD de fabricantes | Registrar fabricantes dos bens |
| CRUD de fornecedores | Registrar fornecedores |
| Vinculação de bens | Associar bem a setor, fabricante e fornecedor |
| Identificação única | Cada entidade com ID único |
| Listagem e filtros | Buscar bens por setor, fabricante ou fornecedor |
| Autenticação JWT | Login e controle de acesso |

### Fora do Escopo do MVP
- Dashboard com gráficos
- Notificações automáticas
- Controle de manutenção
- Upload de imagens
- Sistema de permissões avançado
- Integração com outros sistemas

---

## 📋 Product Backlog Priorizado (Sprints 1-5)

| Prioridade | User Story | Critérios de Aceitação | Pontos | Sprint |
| :---: | --- | --- | :---: | :---: |
| **P1** | Como usuário, quero me autenticar no sistema | Login via JWT e proteção de rotas privadas (Middleware). | **Médio** | 1 |
| **P1** | Como gestor, quero cadastrar Setores e Locais | Nome único por setor e vinculação a local físico. | **Simples** | 1 |
| **P1** | Como gestor, quero cadastrar Bens/Ativos | Nome, tipo e ID único (UUID). Persistência via SQLC. | **Médio** | 1 |
| **P2** | Como gestor, quero cadastrar Fabricantes e Fornecedores | Validação de CNPJ e campos de contato obrigatórios. | **Simples** | 2 |
| **P2** | Como gestor, quero vincular um bem a um setor e fornecedor | Relacionamento correto no DB (Chaves estrangeiras). | **Simples** | 2 |
| **P2** | Como usuário, quero filtrar bens por setor e fabricante | Filtros funcionais via query strings na API. | **Médio** | 2 |
| **P3** | Como gestor, quero registrar a transferência de um bem | Log de histórico (Setor Origem -> Setor Destino). | **Médio** | 3 |
| **P3** | Como gestor, quero **gerar o documento de transferência** | Endpoint que gera um PDF com dados para assinatura. | **Difícil** | 3 |
| **P3** | Como usuário, quero listar bens com paginação | Retorno de meta-dados (page, total_pages, limit). | **Médio** | 3 |
| **P4** | Como gestor, quero **anexar o documento de transferência assinado** | Upload de PDF vinculado à movimentação do bem. | **Difícil** | 4 |
| **P4** | Como gestor, quero atualizar ou remover um bem | Implementação de Soft Delete (remoção lógica). | **Médio** | 4 |
| **P4** | Como desenvolvedor, quero implementar testes automatizados | Cobertura mínima de 70% da lógica de negócio. | **Difícil** | 4 |
| **P5** | Como administrador, quero um desligamento seguro (*Graceful*) | Tratamento de sinais do SO para encerrar conexões. | **Médio** | 5 |
| **P5** | Como desenvolvedor, quero configurar o CI/CD e Deploy | GitHub Actions configurado para Build e Deploy (Docker). | **Difícil** | 5 |

---

### 📝 Legenda de Complexidade (Pontos)
* **Simples:** CRUD básico, poucas regras e sem dependências externas.
* **Médio:** Exige lógica de negócio, autenticação ou manipulação de banco de dados.
* **Difícil:** Envolve bibliotecas de terceiros (PDF), persistência de arquivos (Upload) ou infraestrutura/automação.

### 🎯 Definição de Prioridades
* **P1 - Crítico:** Funcionalidades base sem as quais o sistema não existe.
* **P2 - Necessário:** Relacionamentos e buscas que dão sentido ao inventário.
* **P3 - Valor de Negócio:** Onde o sistema resolve o problema da transferência.
* **P4 - Segurança e Qualidade:** Gestão de arquivos e garantia de estabilidade.
* **P5 - Operação:** Requisitos não-funcionais para colocar o projeto em produção.
---

## 4. Stack Tecnológica

### Backend (Mandatório)
* **Linguagem:** [Go (Golang)](https://go.dev/)
* **Roteamento:** [Chi Router](https://github.com/go-chi/chi)
* **Banco de Dados:** [PostgreSQL](https://www.postgresql.org/)
* **Gerador de Código SQL:** [sqlc](https://sqlc.dev/)

### Frontend (Bônus)
* **Framework:** [React](https://react.dev/) (JavaScript/TypeScript)

### Ferramentas de Suporte
* **Containerização:** Docker & Docker Compose
* **Documentação:** [Ainda avaliando]
* **Testes de API:** Postman / Insomnia
* **Versionamento:** Git & GitHub

**Justificativa:** A stack escolhida atende aos requisitos obrigatórios da disciplina, garantindo alta performance, segurança de tipos no backend e uma interface moderna e reativa no frontend.

---

## 5. Equipe

| Nome | Matrícula | Papel |
| :--- | :--- | :--- |
| Alvaro Soares dos Santos | 20260001348 | Full Stack |
| Ivis de Moura Nobre | 20220028454 | Full Stack |
| José Ben Hur Nascimento de Oliveira | 20240078121 | Full Stack |
---

## 🔗 Links

- **Repositório do projeto:** [https://github.com/Benhurds12/ProjetoWeb2](https://github.com/Benhurds12/ProjetoWeb2)
- **Vídeo Sprint 0:** *(link a ser adicionado)*
- **Documento de proposta (PDF):** *(anexo na entrega do SIGAA)*

---

## 📄 Licença

Este projeto é desenvolvido para fins educacionais na disciplina **DIM0547 - Desenvolvimento de Sistemas Web II com Go** (2026.1), sob orientação do Prof. Fernando Figueira.
