# card-transactions
A card transactions emulation software

Software escrito em Go que simula algumas operações de uma API de administração de pagamentos por cartão.

Em estado alpha.

# Pode ser executado via CLI com (a porta exposta é a 8000):
```bash
go run main.go
```
# Ou utilize o Docker:
```bash
git clone https://github.com/andrepereira/card-transactions.git
cd card-transactions
#Construção de uma imagem Docker com a tag "card-transactions"
docker build . -t card-transactions
#Executa um container com a imagem Docker gerada anteriormente e expõe a porta 8000
docker run -p 8000:8000 card-transactions
```

A esta altura, você terá um container chamado card-transaction rodando com o 
software devidamente compilado e executando, escutando na porta 8000.

O SO do container é um Alpine Linux 64 bits x86.


Os endpoints expostos:

--------------------------------------------------------------------
POST /accounts (criação de uma conta)

Request Body:
{
"document_number": "12345678900"
}

--------------------------------------------------------------------

GET /accounts/:accountId (consulta de informações de uma conta)

Response Body:
{
"account_id": 1,
"document_number": "12345678900"
}

--------------------------------------------------------------------

POST /transactions (criação de uma transação)

Request Body:
{
"account_id": 1,
"operation_type_id": 4,
"amount": 123.45
}

--------------------------------------------------------------------

# Testes:
```bash
cd models
go test -v
```