# 🚀 PostGo API

API RESTful escrita em **Go (Golang)** usando o framework **Gin** com integração ao banco de dados via **GORM**, seguindo arquitetura MVC. A aplicação será consumida por um frontend feito em **Angular**.

---

## 📦 Funcionalidades

- 🔎 Listar todos os posts
- 📄 Visualizar post por ID
- ✍️ Criar novo post
- 🛠 Atualizar post existente
- ❌ Deletar post
- 📑 Documentação Swagger

---

## 🛠 Tecnologias Utilizadas

- Go 1.21+
- Gin
- GORM
- Swagger (Swaggo)
- SQL Server / MySQL / PostgreSQL (configurável)
- Angular (como frontend externo)

---

## 📁 Estrutura do Projeto

PostGo/
│
├── config/ # Configuração de conexão com o banco
├── controllers/ # Camada de controle (handlers)
├── services/ # Lógica de negócios
├── models/ # Estruturas e entidades do banco
├── routes/ # Rotas da API
├── utils/ # Funções utilitárias e helpers
├── docs/ # Swagger gerado automaticamente
├── main.go # Ponto de entrada
└── go.mod / go.sum # Gerenciamento de dependências


---

## ▶️ Como Executar o Projeto

### 1. Clone o repositório

```bash
git clone https://github.com/seu-usuario/PostGo.git
cd PostGo
```
3. Instale as dependências
```
go mod tidy
```

4. Gere a documentação Swagger
```

go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

5. Rode o projeto
```
go run main.go
```

A API estará disponível em: http://localhost:3001

📘 Documentação Swagger
Após rodar o projeto, acesse:

http://localhost:3001/swagger/index.html
