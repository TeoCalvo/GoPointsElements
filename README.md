# GoPointsElements

[![go-elements.png](https://i.postimg.cc/F15LFyWT/go-elements.png)](https://postimg.cc/xN5CtzSm)

Programa em [GoLang](https://golang.org/) para realizar adição ou remoção de pontos do StreamElements por usuário ou uma lista de usuários. A ideia é facilitar a vida do Streamer que precisa realizar alterações de pontuação de forma massiva ou até mesmo individual.

## Setup

Para utilizar o programa é necessário configurar duas variáveis de ambiente:

```
CHANNEL_ID = "SEU_ID_DE_CANAL"
STREAMELEMENTS_TOKEN = "SEU_TOKEN_DO_STREAMELEMENTS"
```

Voê pode utilizar o arquivo ```.env``` ou até mesmo fazer uso do ```EXPORT``` em seu próprio ambiente.

## Uso

### Usuário 
Para adicionar pontos a um determinado usuário, basta executar o seguinte comando:

```bash
$go run main.go --username nome_usuario_twitch --points quantidade_pontos
```
Assim, para adicionar 100 pontos para o usuário teomewhy, tem-se:

```sh
$go run main.go --username teomewhy --points 100
```

### Lista de Usuários

Já para uma lista de usuários, deve-se criar um arquivo contendo uma lista de `usernames`, onde cada linha deste arquivo corresponde a um `username`.

Imagine então que se queira realizar a adição de 100 pontos para os usuários `teomewhy` e `artesanah`. Assim, cria-se o arquivo: `~/names.txt` com duas linhas:

```
teomewhy
artesanah
```

Agora basta executar o programa da seguinte maneira:

```bash
$go run main.go --filepath ~/names.txt --points 100
```