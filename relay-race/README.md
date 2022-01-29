# Corrida de Revezamento
O objetivo deste trabalho é simular uma corrida de revezamento em que existe múltiplas equipes concorrendo e cada equipe possui quatro corredores

# Tecnologia Utilizada

* Go (Versão 1.17.6)

# A solução

A solução foi construída utilizando os seguintes conceitos:
* Goroutines: para simular o corredor de cada equipe;
* Channels sem buffer: para servir como canal de comunicação entre os corredores de cada equipe e passar o bastão;
* Channel com buffer: para armazenar a ordem de chegada de cada equipe;
* WaitGroup: para a corrida ser finalizada somente quando todas as equipes terminarem todo o percurso.

# Compilação

Para simplesmente executar um código fonte em Go sem gerar um arquivo executável, faça:

```
$ go run relayRace.go
```

Já para gerar um arquivo executável, faça:

```
go build -o relayRace relayRace.go
```

# Desenvolvedora

* Karina Maria [@karinamaria](https://github.com/karinamaria/)