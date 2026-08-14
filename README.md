# Mini Kanban

Aplicação fullstack de gerenciamento de tarefas no estilo Kanban, desenvolvida com React no frontend e Go no backend, com persistência em PostgreSQL via Docker.

---

## Tecnologias

| Camada | Tecnologia |
|--------|-----------|
| Frontend | React 19 + Vite |
| Backend | Go 1.26 + chi router |
| Banco de dados | PostgreSQL 16 |
| Infraestrutura | Docker + Docker Compose |

---

## Pré-requisitos

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Go 1.22+](https://golang.org/dl/)
- [Node.js 18+](https://nodejs.org/)

---

## Como rodar

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/mini-kanban.git
cd mini-kanban
```

### 2. Suba o banco de dados

```bash
docker compose up -d
```

> O PostgreSQL ficará disponível na porta `5433`. A tabela `tasks` é criada automaticamente na primeira execução.

### 3. Inicie o backend

```bash
cd backend
go run .
```

O servidor estará disponível em `http://localhost:8080`.

### 4. Inicie o frontend

Em outro terminal:

```bash
cd frontend
npm install
npm run dev
```

A aplicação estará disponível em `http://localhost:5173`.

---

## Endpoints da API

| Método | Rota | Descrição |
|--------|------|-----------|
| `GET` | `/tasks` | Lista todas as tarefas |
| `POST` | `/tasks` | Cria uma nova tarefa |
| `PUT` | `/tasks/{id}` | Atualiza título, descrição ou status |
| `DELETE` | `/tasks/{id}` | Remove uma tarefa |

### Exemplo de payload

```json
{
  "title": "Criar tela de login",
  "description": "Validação de e-mail obrigatória",
  "status": "todo"
}
```

**Status válidos:** `todo` · `in_progress` · `done`

---

## Estrutura do projeto

```
mini-kanban/
├── backend/
│   ├── main.go        # Entry point e configuração de rotas
│   ├── handlers.go    # Handlers HTTP (CRUD)
│   ├── models.go      # Struct Task e validações
│   └── db.go          # Conexão com PostgreSQL e criação da tabela
├── frontend/
│   └── src/
│       ├── components/
│       │   ├── Column.jsx     # Coluna do Kanban com formulário
│       │   └── TaskCard.jsx   # Card com ações de editar, mover e deletar
│       ├── services/
│       │   └── api.js         # Camada de comunicação com a API
│       └── App.jsx            # Estado global e lógica principal
├── docs/
│   ├── user-flow.jpeg
│   └── data-flow.jpeg
├── docker-compose.yml
└── README.md
```

---

## Decisões técnicas

**Go com chi router** — escolhido pela leveza e performance. O chi oferece roteamento com parâmetros de URL (`{id}`) sem a complexidade de frameworks maiores, mantendo compatibilidade com a interface padrão `net/http`.

**PostgreSQL via Docker** — garante ambiente reproduzível sem instalação local. A tabela é criada automaticamente pelo backend na primeira inicialização, eliminando a necessidade de migrations para o escopo do MVP.

**React com Vite** — substituição ao Create React App por ser significativamente mais rápido no desenvolvimento. O estado global das tarefas é mantido no `App.jsx` e distribuído via props, evitando a complexidade de um gerenciador de estado externo para o escopo do projeto.

**Sem ORM** — o banco é acessado diretamente via `database/sql`, tornando as queries explícitas e o comportamento previsível, o que facilita o aprendizado do fluxo de dados.

---

## Limitações conhecidas

- Sem autenticação — qualquer usuário pode criar, editar e deletar tarefas
- Sem paginação — todas as tarefas são carregadas de uma vez
- Sem drag and drop — movimentação feita pelos botões ← →
- CORS configurado apenas para `localhost:5173`
