package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	filepath := flag.String("filepath", "", "Endereço do arquivo com nome dos usuário para ganhar pontos")
	username := flag.String("username", "", "Nome do usuário a receber pontos")
	points := flag.Int("points", 0, "Quantidade de pontos a ser adicionada para usuário")
	flag.Parse()

	if *username != "" {
		addpoints(*username, *points)
	} else if *filepath != "" {
		names := readNames(*filepath)
		for _, n := range names {
			addpoints(n, *points)
		}
	}
}

func addpoints(username string, points int) {

	channelID := os.Getenv("CHANNEL_ID")
	channelToken := os.Getenv("STREAMELEMENTS_TOKEN")

	client := &http.Client{}
	url := "https://api.streamelements.com/kappa/v2/points/%s/%s/%s"
	url = fmt.Sprintf(url, channelID, username, strconv.Itoa(points))
	fmt.Println(url)

	req, _ := http.NewRequest(http.MethodPut, url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+channelToken)

	resp, _ := client.Do(req)
	fmt.Println(resp)
	resp.Body.Close()
}

func readNames(filepath string) []string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("Erro ao ler o arquivo")
	}
	names := string(data)
	return strings.Split(names, "\n")
}
