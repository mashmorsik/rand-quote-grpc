package internal

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
)

const dataPath = "./hp_quotes.yaml"

type Data struct {
	Quotes map[string][]string `yaml:"Quotes"`
}

func ReturnQuote(name string) (string, error) {
	q := Data{}
	var randQuote string

	quotesData, err := os.ReadFile(dataPath)
	if err != nil {
		return err.Error(), fmt.Errorf("can't unmarshal data file: %s", dataPath)
	}

	err = yaml.Unmarshal(quotesData, &q)
	if err != nil {
		return err.Error(), fmt.Errorf("can't unmarshal data file: %s", dataPath)
	}

	for character, _ := range q.Quotes {
		if character == name {
			randQuote = q.Quotes[character][rand.Intn(len(q.Quotes[character]))]
		}
	}

	return randQuote, nil
}
