package points

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Amount struct {
	Channel       string `json:"channel"`
	Username      string `json:"username"`
	Points        int    `json:"points"`
	PointsAlltime int    `json:"pointsAlltime"`
	Watchtime     int    `json:"watchtime"`
	Rank          int    `json:"rank"`
}

// Addpoints adiciona pontos para o usuário
func Addpoints(username string, points int) {

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

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}

//GetAmount obtem a informação de saldo de pontos do usuário
func GetAmount(username string) int {

	channelID := os.Getenv("CHANNEL_ID")
	channelToken := os.Getenv("STREAMELEMENTS_TOKEN")

	client := &http.Client{}
	url := "https://api.streamelements.com/kappa/v2/points/%s/%s"
	url = fmt.Sprintf(url, channelID, username)
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+channelToken)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	amount := Amount{}
	json.NewDecoder(resp.Body).Decode(&amount)
	return amount.Points
}

// ReadNames lê lista de nomes
func ReadNames(filepath string) []string {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("Erro ao ler o arquivo")
	}
	names := string(data)
	return strings.Split(names, "\n")
}
