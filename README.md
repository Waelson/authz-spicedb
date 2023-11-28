# Serviço de Autorização baseado no SpiceDB (Google Zanzibar)
Esse projeto demonstra a integração entre uma aplicação web e o mecanismo de autorização SpiceDB desenvolvido pela Authzed baseado no paper Zanzibar da Google

## Pré requisites
- Golang 1.21
- Docker
## Instalação
Após assegur que o Docker está instalado em sua máquina. Exlsecute o comado abaixo para executar o SpiceDB em sua máquina. Atente-se para o parametro `--grpc-preshared-key` o valor dele será utilizado para conectar sua aplicação ao SpiceDB.

`
docker run --rm -p 50051:50051 authzed/spicedb serve --grpc-preshared-key "somerandomkeyhere"
`

