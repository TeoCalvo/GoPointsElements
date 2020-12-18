package main

import (
	"flag"

	"github.com/TeoCalvo/GoPointsElements/points"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	filepath := flag.String("filepath", "", "Endereço do arquivo com nome dos usuário para ganhar pontos")
	username := flag.String("username", "", "Nome do usuário a receber pontos")
	nPoints := flag.Int("points", 0, "Quantidade de pontos a ser adicionada para usuário")
	flag.Parse()

	if *username != "" {
		points.Addpoints(*username, *nPoints)
	} else if *filepath != "" {
		names := points.ReadNames(*filepath)
		for _, n := range names {
			points.Addpoints(n, *nPoints)
		}
	}
}
