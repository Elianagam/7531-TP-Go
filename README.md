### 7531-TP-Go

# API de Búsqueda de tweets de los candidatos presidenciales 2019
Correr ejecutando `go run twitter/twitter.go`

## Búscar en todos los tweets
*GET* `localhost:8080/tweets/search?query=chanta`

## Búscar en todos los tweets de un candidato
*GET* `localhost:8080/tweets/search/:userId?query=chanta`

Candidatos disponibles: "alferdez", "mauriciomacri", "jlespert", "NicolasdelCano", "juanjomalvinas".

Están disponibles los últimos 2000 tweets de cada uno aproximadamente.

### Informe:
https://docs.google.com/document/d/1D4BASyrIlgwWW0Wlr7OcXWjungn3-o4pnPUjqdLvomE/edit?usp=sharing

### Presentación:
https://docs.google.com/presentation/d/1T-CI9V9Mc9CrgydfcGEKrLJZNXLh-vqCE92cRYaoYPk/edit
