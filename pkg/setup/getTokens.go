package setup

import (
	"io/ioutil"
	"regexp"
)

func GetTelegramAPItoken() (string, error) {
	var token string
	configFileContent, err := ioutil.ReadFile(".\\api\\config.txt")
	if err != nil {
		panic(err)
	}

	statementTelegramAPItoken := regexp.MustCompile("TelegramAPI_token*(\\s)=*(\\s)(.*)")
	configLine := statementTelegramAPItoken.FindString(string(configFileContent))
	token = regexp.MustCompile("(\\d+):(.*)").FindString(configLine)

	return token, nil
}
