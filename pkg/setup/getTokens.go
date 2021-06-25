package setup

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func GetTelegramAPItoken() (string, error) {
	var token string
	configFileContent, err := ioutil.ReadFile("C:\\Users\\ragon\\projects\\anusRipper\\api")
	if err != nil {
		panic(err)
	}

	statementTelegramAPItoken := regexp.MustCompile("TelegramAPI_token(.\\s)=(.\\s)(.*)\n")
	token = statementTelegramAPItoken.FindString(string(configFileContent))

	fmt.Println(configFileContent)

	return token, nil
}
