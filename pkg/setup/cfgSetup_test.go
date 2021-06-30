package setup

import (
	"fmt"
	"testing"
)

func TestCfgParser(t *testing.T) {
	const input = `
TelegramAPI_token = 1488:ABCD76F6

host        = 1host1
user        = 1user1
password    = 1password1
dbname      = 1dbname1
sslmode     = 1sslmode1
`

	fmt.Println("Testing config parser...")

	var isOk = true
	cfgSet := initCfgStructure()
	cfgSet = configParser([]byte(input), cfgSet)

	//for key, val := range cfgSet.ConfigMap {
	//	fmt.Printf("%s   :   %s\n", key, val)
	//}

	if cfgSet.ConfigMap["TelegramAPI_token"] != "1488:ABCD76F6" {
		isOk = false
	}
	if cfgSet.ConfigMap["host"] != "1host1" {
		isOk = false
	}
	if cfgSet.ConfigMap["user"] != "1user1" {
		isOk = false
	}
	if cfgSet.ConfigMap["password"] != "1password1" {
		isOk = false
	}
	if cfgSet.ConfigMap["dbname"] != "1dbname1" {
		isOk = false
	}
	if cfgSet.ConfigMap["sslmode"] != "1sslmode1" {
		isOk = false
	}

	if !isOk {
		t.Errorf("test failed - Parsing error. Expected: %v\n", input)
	}
}
