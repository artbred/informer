package src

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type CallResponse struct {
	MediaSession string `json:"media_session_access_secure_url"`
}

var baseUrl string

func Call(data string) bool {
	base, err := url.Parse(baseUrl)
	if err != nil {
		logrus.Errorf("Voximplant parse base url, %s", err)
		return false
	}

	params := url.Values{}
	params.Add("script_custom_data", data)
	base.RawQuery += params.Encode()

	client := http.Client{}
	res, err := client.Post(base.String(), "", nil)
	if err != nil {
		logrus.Errorf("Voximplant get request, %s", err)
		return false
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			logrus.Errorf("Voximplant can't close body, %s", err)
		}
	}()

	body, err := ioutil.ReadAll(res.Body); if err != nil {
		logrus.Errorf("Voximplant can't read body")
	}


	_ = body
	return true
}

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Panic(err)
	}

	baseUrl = fmt.Sprintf("https://api.voximplant.com/platform_api/StartScenarios/?account_id=%s&api_key=%s&rule_id=%s&",
		os.Getenv("VOXIMPLANT_ACCOUNT_ID"), os.Getenv("VOXIMPLANT_API_KEY"), os.Getenv("VOXIMPLANT_RULE_ID"))
}
