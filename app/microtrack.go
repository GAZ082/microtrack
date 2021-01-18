package microtrack

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

type RecordResumenMensual struct {
	AccCent       int32
	Accelerations int32
	Braking       int32
	Identificador string
	InfHigh       int32
	InfLight      int32
	InfMedium     int32
	InfTotal      int32
	Kms           float64
	Score         float64
	ScoreAP       float64
	ScoreRR       float64
	ScoreZU       float64
	Period        int32
}

var (
	config = readConfig("microtrack")
)

func readConfig(fileName string) (config *viper.Viper) {
	config = viper.New()
	config.SetConfigName(fileName)
	config.SetConfigType("yaml")
	config.AddConfigPath("$HOME/.config")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return
}

func GetApiToken(user, password, clientID string) (token string) {
	//Generado por POSTMAN
	url := "https://mtkb2c.b2clogin.com/tfp/mtkb2c.onmicrosoft.com/B2C_1_ropc/oauth2/v2.0/token"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	writer.WriteField("client_id", clientID)
	writer.WriteField("grant_type", "password")
	writer.WriteField("username", user)
	writer.WriteField("password", password)
	writer.WriteField("response_type", "token id_token")
	writer.WriteField("scope", "openid https://mtkb2c.onmicrosoft.com/WebAPIProactivo/user_impersonation")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	return gjson.GetBytes(body, "access_token").String()

}

// getResumenMensual. Period format YYYYMM
func getResumenMensual(period int, IDPlantilla string, version int) []byte {
	mesTime, _ := time.Parse("200601", strconv.Itoa(period))
	desde := mesTime.Format("2006-01-02")
	hasta := mesTime.AddDate(0, 1, -1).Format("2006-01-02")
	url := fmt.Sprintf("https://webapiproactivo-clientes.azurewebsites.net/api/%v/Ranking/%v/ResumenMensual?desde=%v&hasta=%v", version, IDPlantilla, desde, hasta)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	token := GetApiToken(config.GetString("app.user"), config.GetString("app.password"),
		config.GetString("app.clientID"))
	req.Header.Add("Authorization", "Bearer "+token)
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

func processResumenMensual(input []byte, period int32) (output []RecordResumenMensual) {
	for _, v := range gjson.ParseBytes(input).Array() {
		record := RecordResumenMensual{
			AccCent:       int32(v.Get("accCent").Int()),
			Accelerations: int32(v.Get("accelerations").Int()),
			Braking:       int32(v.Get("braking").Int()),
			Identificador: v.Get("identificador").String(),
			InfHigh:       int32(v.Get("infHigh").Int()),
			InfLight:      int32(v.Get("infLight").Int()),
			InfMedium:     int32(v.Get("infMedium").Int()),
			InfTotal:      int32(v.Get("infTotal").Int()),
			Kms:           v.Get("kms").Float(),
			Score:         v.Get("score").Float(),
			ScoreAP:       v.Get("scoreAP").Float(),
			ScoreRR:       v.Get("scoreRR").Float(),
			ScoreZU:       v.Get("scoreZU").Float(),
			Period:        period,
		}
		output = append(output, record)
	}
	return
}
