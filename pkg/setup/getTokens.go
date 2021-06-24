package setup

import (
	"fmt"
	"io/ioutil"
	"regexp"
)


func GetTelegramAPItoken() (string, error) {
	var token string
	configFileContent, err := ioutil.ReadFile("api/config.txt")
	if err != nil {
		panic(err)
	}

	statementTelegramAPItoken := regexp.MustCompile("TelegramAPI_token(.\\s)=(.\\s)(.*)\n")
	token = statementTelegramAPItoken.FindString(string(configFileContent))

	fmt.Println(token)

	return "", nil
}