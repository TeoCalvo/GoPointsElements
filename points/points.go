package points

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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

	resp, _ := client.Do(req)
	fmt.Println(resp)
	resp.Body.Close()
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
