package pkg

import (
	"fmt"
	randquotev1 "github.com/mashmorsik/rand-quote-grpc/api/proto/github.com/mashmorsik/rand-quote-grpc"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
)

const dataPath = "./hp_quotes.yaml"

type Data struct {
	Quotes map[string][]string `yaml:"Quotes"`
}

func ReturnQuote(enum randquotev1.Name) (string, error) {
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

	nameStr, err := EnumMatcher(enum)

	for character, _ := range q.Quotes {
		if character == nameStr {
			randQuote = q.Quotes[character][rand.Intn(len(q.Quotes[character]))]
		}
	}

	return randQuote, nil
}

func ReturnQuotesList(enum randquotev1.Name) ([]string, error) {
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

	nameStr, err := EnumMatcher(enum)

	for character, quotes := range q.Quotes {
		if character == nameStr {
			quotesList = quotes
		}
	}

	return quotesList, nil
}

func EnumMatcher(enum randquotev1.Name) (string, error) {
	enumName := map[randquotev1.Name]string{
		randquotev1.Name_ALBUS_DUMBLEDORE:     "Albus Dumbledore",
		randquotev1.Name_HERMIONE_GRANGER:     "Hermione Granger",
		randquotev1.Name_RON_WEASLEY:          "Ron Weasley",
		randquotev1.Name_SEVERUS_SNAPE:        "Severus Snape",
		randquotev1.Name_PROFESSOR_MCGONAGALL: "Professor McGonagall",
		randquotev1.Name_LUNA_LOVEGOOD:        "Luna Lovegood",
		randquotev1.Name_HAGRID:               "Hagrid",
	}

	var name string

	for e, _ := range enumName {
		if e == enum {
			name = enumName[e]
		}
	}
	return name, nil
}
