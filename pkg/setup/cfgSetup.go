package setup

/*
Пакет для обработки конфигурационного файла и инициализации переменных проекта.

func initCfgStructure() ConfigSet	---	инициализирует структуру ключами и пустыми значениями


func configParser(configFileContent []byte, cfgSet ConfigSet) ConfigSet	  ---    парсит содержимое конфига


func Init(configPath string) ConfigSet   ---   заполняет структуру с переменными и возвращает ее


func GetCfgSet() ConfigSet   ---   возвращает ConfigSet. !!Паника при использовании до инициализации!!
*/

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var cfgSet ConfigSet

type ConfigSet struct {
	ConfigMap map[string]string //Ключ - поле, значение - значение
	IsOk      bool              //Валидность полей
}

func initCfgStructure() ConfigSet {
	var cfgSet ConfigSet

	cfgSet.ConfigMap = make(map[string]string)

	cfgSet.ConfigMap["host"] = ""
	cfgSet.ConfigMap["user"] = ""
	cfgSet.ConfigMap["password"] = ""
	cfgSet.ConfigMap["dbname"] = ""
	cfgSet.ConfigMap["port"] = ""
	cfgSet.ConfigMap["sslmode"] = ""
	cfgSet.ConfigMap["TelegramAPI_token"] = ""

	return cfgSet
}

func configParser(configFileContent []byte, cfgSet ConfigSet) ConfigSet {
	var isOk bool = true

	for key := range cfgSet.ConfigMap {

		var configLine string

		regexpString := fmt.Sprintf("%s*(\\s)=*(\\s)(.*)", key)
		statementConfigLine, err := regexp.Compile(regexpString)
		if err == nil {
			configLine = statementConfigLine.FindString(string(configFileContent))
		} else {
			fmt.Printf("Error during reading: %s\n", key)
			fmt.Println("Skipping....")
			isOk = false
			continue
		}

		if idx := strings.Index(configLine, "="); idx != -1 {
			configLine = configLine[idx:]
		} else {
			fmt.Printf("Wrong config file format: %s\n", configLine)
		}

		cfgSet.ConfigMap[key] = regexp.MustCompile("(\\w+)(:*)(\\w+)").FindString(configLine)

	}

	cfgSet.IsOk = isOk
	return cfgSet
}

func Init(configPath string) ConfigSet {

	if cfgSet.IsOk {
		return cfgSet
	}
	configFileContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		cfgSet.IsOk = false
		return cfgSet
	}
	cfgSet = initCfgStructure()
	cfgSet = configParser(configFileContent, cfgSet)

	return cfgSet
}

func GetCfgSet() ConfigSet {
	if cfgSet.IsOk {
		return cfgSet
	} else {
		panic("ConfigSet may not be initiated or smt is wrong...")
	}
}
