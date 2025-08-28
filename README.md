# Worker API

API assíncrona para processamento de pedidos utilizando filas SQS e worker em Go.

## Sobre o Projeto

Este projeto implementa uma API REST para processamento assíncrono de pedidos, utilizando uma arquitetura baseada em filas. Quando um pedido é criado, ele é enviado para uma fila SQS (Simple Queue Service) e posteriormente processado por um worker, que atualiza o status do pedido para "paid".

### Tecnologias Utilizadas

- Go 1.21.3
- Echo Framework (Web Framework)
- MySQL (Banco de dados)
- AWS SQS (Sistema de filas)
- Docker & Docker Compose
- GORM (ORM)


## Como Executar

### Pré-requisitos

- Docker
- Docker Compose
- Git

### Instalação

1. Clone o repositório:
```bash
git clone https://github.com/neiltonrodriguez/worker-golang
cd worker-golang
```

2. Configure as variáveis de ambiente:
```bash
cp .env-example .env
```

3. Inicie os containers:
```bash
docker-compose up -d --build
```

Nota: Se houver conflito com a porta 3306 do MySQL, você pode:
- Parar o serviço MySQL local: `sudo service mysql stop`
- Ou alterar a porta no arquivo `docker-compose.yaml`

## Funcionalidades

- Criação assíncrona de pedidos
- Processamento em background via worker
- Sistema de filas com AWS SQS
- CRUD completo de produtos
- Paginação de resultados
- Logs estruturados

## API Endpoints

- `POST /orders` - Criar novo pedido
- `GET /orders/:id` - Buscar pedido específico

## Desenvolvedor

Desenvolvido por Neilton Rodrigues