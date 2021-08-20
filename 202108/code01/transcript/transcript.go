package transcript

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetScores() {
	file, err := ioutil.ReadFile("transcript/transcript.yaml")
	if err != nil {
		panic(err)
	}
	result := make(map[string]map[string]string)
	yaml.Unmarshal(file, &result)
	for key, value := range result["成绩"] {
		fmt.Println(key, value)
	}
}
