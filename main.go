package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

var jsonData = flag.String("json", "", "JSON content to parse")
var jsonFile = flag.String("jfile", "", "File of JSON parse")

func init() {
	flag.Parse()
}

func main() {

	// fmt.Println(os.Args)

	json, err := getJSONParsed()
	var jsonArray []interface{}
	if err != nil {
		jsonArray, err = getJSONArrayNParsed()
		if err != nil {
			logrus.WithError(err).Error("invalid JSON")
			return
		}
	}

	output := ""
	if json != nil && jsonArray == nil {
		output = "Root: (object)\n"
		output += parse(json, 1)
	} else {
		output = "Root: (array)\n"
		output += parseArray(jsonArray, 1)

	}
	fmt.Println(output)
}

func parse(json map[string]interface{}, deep int) string {
	output := ""
	for key, value := range json {
		typeStr := getType(value)

		output += fmt.Sprintf("%s%s: (%s)\n", getTabSpace(deep), key, typeStr)
		if typeStr == "object" {
			output += parse(value.(map[string]interface{}), deep+1)
		} else if typeStr == "array" {
			output += parseArray(value.([]interface{}), deep+1)
		}
	}

	return output
}

func parseArray(json []interface{}, deep int) string {
	output := ""
	for _, value := range json {
		typeStr := getType(value)

		if typeStr == "object" {
			output += parse(value.(map[string]interface{}), deep+1)
		} else if typeStr == "array" {
			output += parseArray(value.([]interface{}), deep+1)
		}
	}

	return output
}

func getJSONParsed() (map[string]interface{}, error) {
	var result map[string]interface{}

	if *jsonData == "" && *jsonFile != "" {
		content, err := readJSON(*jsonFile)
		if err != nil {
			return nil, err
		}
		*jsonData = content
	}

	err := json.Unmarshal([]byte(*jsonData), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getJSONArrayNParsed() ([]interface{}, error) {
	var result []interface{}

	if *jsonData == "" && *jsonFile != "" {
		content, err := readJSON(*jsonFile)
		if err != nil {
			return nil, err
		}
		*jsonData = content
	}

	err := json.Unmarshal([]byte(*jsonData), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//readJSON reads data from the specified file
func readJSON(filePath string) (string, error) {
	//checks if the file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return "", nil
	}

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	return string(file), nil
}

func getTabSpace(deep int) string {
	if deep == 0 {
		return ""
	}

	// leftpad := ""
	// for i := 0; i < deep; i++ {
	// 	leftpad += "   "
	// }
	format := fmt.Sprintf("%%%ds", deep*2)
	return fmt.Sprintf(format, "|_")
}

func getType(value interface{}) string {
	switch value.(type) {
	case int:
		return "numeric"
	case int8:
		return "numeric"
	case int16:
		return "numeric"
	case int32:
		return "numeric"
	case int64:
		return "numeric"
	case byte:
		return "numeric"
	case uint:
		return "numeric"
	//case uint8: (byte)
	case uint16:
		return "numeric"
	case uint32:
		return "numeric"
	case uint64:
		return "numeric"
	case uintptr:
		return "numeric"
	case float64:
		return "float"
	case float32:
		return "float"
	case string:
		return "string"
	case []interface{}:
		return "array"
	case map[string]interface{}:
		return "object"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}
