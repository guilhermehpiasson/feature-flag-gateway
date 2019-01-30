#       feature-flag-gateway

Projeto com um gateway minimalista para consulta de Feature Flags em uma base Redis

## O Projeto

O projeto é feito Golang, com o objetivo de prover um uma rota paara consulta de estado de featureflag.
Esse estado é persisitdo em uma base Redis.
A plataforma de desenvolvimento é o Go.

> Para maiores informações sobre o Pattern Feature Toggle, acesse: https://martinfowler.com/articles/feature-toggles.html

### Modulos Principais

* [Redigo](https://github.com/gomodule/redigo)
* [Gin](https://github.com/gin-gonic/gin)

### Rotas

> api/v1/ff/:id - Rota  ser chamada para consulta de status de uma FeatureFlag*

Observação: 
1. Essa feature flag deve ter sido ja armazenada no Redis, conforme exemplo abaixo:
```
set ff:ff-test "true"
```
2. Exemplo de URL para consulta de uma FeatureFlag -> http://localhost:8080/api/v1/ff/ff-test

### Teste

Para testar a respectiva aplicação, basta disparar o seguinte comando:
```
go test
```
Nessa ideia, sera executado um teste de integração, onde será chamada a rota api/v1/ff/ff-test e validado o código de retorno. No caso, se for igual a 200 a execução é aceita, do contrário a execução é rejeitada.s

