Gopportunities

Uma API em Go para cadastro, listagem e gestão de vagas de emprego e estágio, construída com as bibliotecas/go frameworks Gin, GORM e SQLite.

🧩 Descrição

Gopportunities é uma API backend simples e funcional, ideal para quem quer aprender ou usar uma estrutura leve para gerenciar vagas de emprego ou estágio. Com ela, usuários podem cadastrar vagas, buscar por oportunidades e (se quiser) expandir com autenticação, filtros, persistência com SQLite, etc. A ideia é oferecer um ponto de partida claro para quem está aprendendo Go ou desenvolvendo um microserviço backend para ofertas de emprego.

💡 Funcionalidades

CRUD de vagas (criar, listar, atualizar, deletar)

Persistência usando SQLite com ORM via GORM

Rotas e endpoints definidos com Gin

Estrutura organizada para separar rotas (router), handlers, schemas/configurações

🚀 Começando / Como usar
Pré-requisitos

Go instalado (versão compatível conforme go.mod)

Git

Instalação e execução (local)
# Clone o repositório
git clone https://github.com/felipematheus1337/Gopportunities.git

cd Gopportunities

# Baixe dependências
go mod download

# Execute a aplicação
go run main.go


Em seguida, a API estará rodando (por padrão, em http://localhost:8080 — ajuste conforme estiver configurado).

Exemplos de uso

GET /vagas — lista todas as vagas

POST /vagas — cria uma nova vaga (envie JSON com os dados da vaga)
(adicione mais exemplos conforme seus endpoints reais)

🛠️ Tecnologias e ferramentas

Go
— linguagem principal

Gin
— framework HTTP para Go

GORM
— ORM para Go

[SQLite] — banco de dados leve, integrado localmente

✅ Status do projeto

Em estágio inicial — funcionando para operações básicas de cadastro/listagem, ideal para novos projetos ou estudos.

🤝 Como contribuir

Contribuições são bem-vindas! Se quiser adicionar funcionalidades (autenticação, filtros, front-end, testes, etc), abra uma issue ou pull request.

📄 Licença MIT