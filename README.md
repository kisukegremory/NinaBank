# NinaBank

This projects seeks to build a fake bank webapp, where the user after its authentication will be able to:
- account create
- account login
- bank deposit
- cash withdrawal
- bank balance
- cash transfer between ninaccounts
- bank statement

The interface will be implemented using htmx, we will use golang with gin and sqlc/sqlx to be our server side render and backend, kafka will be used for message queues, mysql will be used as our database, those choices were made to learn some of these tecnologies that i never been able to use in my last projects, in specific htmx, kafka and sqlc


These project has the purpose of being an opportunity to design a scalable system that could upscaled horinzontally, the only point that is not been considered here is the database reliability, if we were going to a production enviroment i would pick the AWS RDS to manage the database and apply the backups and replicas within different regions to be sure that will be persisted, but for our study purposes we are going to rely on docker images!

inspired by the devgym challenge: https://app.devgym.com.br/challenges/9af13172-e1fe-4c2e-ac10-cb6b0bcf2efc


# Todo

## Backend Service
- [ ] Fazer endpoints sem lógica interna (com versionamento incluso)
- [ ] Implementar Docker do mysql & Kafka
- [ ] Implementar o account create (REST)
    - [ ] modelo AccountUser
    - [ ] criptografar os dados sensíveis (salt and pepper)
    - [ ] sqlc
- [ ] Implementar o account login (REST)
    - [ ] Validador de dados
    - [ ] Endpoint
    - [ ] Mensagens de erro
- [ ] Adicionar Autenticação nos endpoints
- [ ] Implementar o cash withdrawal
    - [ ] model BankingOperations
    - [ ] model AccountBalance
    - [ ] fila ninabank.cash.withdrawal
    - [ ] Consumer (atomic operation)
- [ ] Implementar o bank deposit
    - [ ] fila ninabank.cash.deposit
    - [ ] Consumer (atomic operation)
- [ ] Implementar a consulta de saldo
    - [ ] sqlc
- [ ] Implementar a transferência entre contas
    - [ ] modelo BankTransfer
    - [ ] fila ninabank.cash.transfer
    - [ ] Consumer (atomic operations: AccountBalance + Cash withdrawal + Cash Deposit)
- [ ] Implementar o endpoint de bank statement
    - [ ] sqlc
    - [ ] modelo Modelo BankStatement: `List<Transactions>`

## Frontend Service
- [ ] estudar o htmx
- [ ] Coletar fotos da nina para status
- [ ] especificar telas para cada caso com figma (?)

* planejamento em desenvolvimento