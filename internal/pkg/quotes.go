package pkg

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

	nameStr, err := EnumMatcher(name)

	for character, _ := range q.Quotes {
		if character == nameStr {
			randQuote = q.Quotes[character][rand.Intn(len(q.Quotes[character]))]
		}
	}

	return randQuote, nil
}

func ReturnQuotesList(name string) ([]string, error) {
	q := Data{}
	var quotesList []string

	quotesData, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal data file: %s", dataPath)
	}

	err = yaml.Unmarshal(quotesData, &q)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal data file: %s", dataPath)
	}

	nameStr, err := EnumMatcher(name)

	for character, quotes := range q.Quotes {
		if character == nameStr {
			quotesList = quotes
		}
	}

	return quotesList, nil
}

func EnumMatcher(enumStr string) (string, error) {
	enumName := map[string]string{
		"NAME_ALBUS_DUMBLEDORE":     "Albus Dumbledore",
		"NAME_HERMIONE_GRANGER":     "Hermione Granger",
		"NAME_RON_WEASLEY":          "Ron Weasley",
		"NAME_SEVERUS_SNAPE":        "Severus Snape",
		"NAME_PROFESSOR_MCGONAGALL": "Professor McGonagall",
		"NAME_LUNA_LOVEGOOD":        "Luna Lovegood",
		"NAME_HAGRID":               "Hagrid",
	}

	var nameStr string

	for enum, _ := range enumName {
		if enum == enumStr {
			nameStr = enumName[enum]
		}
	}
	return nameStr, nil
}
