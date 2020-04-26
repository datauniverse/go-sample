package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Game struct {
	Meta    Metadata            `yaml:"meta"`
	Info    Information         `yaml:"info"`
	Innings []map[string]Inning `yaml:"innings"`
}

type Metadata struct {
	DataVersion float32 `yaml:"data_version"`
	Created     string  `yaml:"created"`
	Revision    int16   `yaml:"revision"`
}

type Information struct {
	City        string `yaml:"city"`
	Gender      string `yaml:"gender"`
	MatchType   string `yaml:"match_type"`
	Venue       string `yaml:"venue"`
	Competition string `yaml:"competition"`

	PlayerOfMatch []string `yaml:"player_of_match"`
	Teams         []string `yaml:"teams"`
	Umpires       []string `yaml:"umpires"`
	Dates         []string `yaml:"dates"`

	NeutralVenue int16 `yaml:"neutral_venue"`
	Overs        int16 `yaml:"overs"`

	MatchTypeNumber int32 `yaml:"match_type_number"`

	Supersubs map[string]string   `yaml:"supersubs`
	BowlOut   []map[string]string `yaml:"bowl_out"`

	Outcome Outcome `yaml:"outcome"`
	Toss    Toss    `yaml:"toss"`
}

type Toss struct {
	Decision string `yaml:"decision"`
	Winner   string `yaml:"winner"`
}

type Outcome struct {
	BowlOut    string `yaml:"bowl_out"`
	Eliminator string `yaml:"eliminator"`
	Result     string `yaml:"result"`
	Method     string `yaml:"method"`
	Winner     string `yaml:"winner"`
	By         By     `yaml:"by"`
}

type By struct {
	Innings string `yaml:"innings"`
	Runs    int16  `yaml:"runs"`
	Wickets int16  `yaml:"wickets"`
}

type Inning struct {
	Team        string                `yaml:"team"`
	Declared    string                `yaml:"declared"`
	AbsentHurt  []string              `yaml:"absent_hurt"`
	PenaltyRuns PenaltyRuns           `yaml:"penalty_runs"`
	Deliveries  []map[string]Delivery `yaml:"deliveries`
}

type PenaltyRuns struct {
	Pre  int16 `yaml:"pre"`
	Post int16 `yaml:"post"`
}

type Delivery struct {
	Batsman      string       `yaml:"batsman"`
	Bowler       string       `yaml:"bowler"`
	NonStriker   string       `yaml:"non_striker"`
	Extras       Extras       `yaml:"extras"`
	Runs         Runs         `yaml:"runs"`
	Wicket       Wicket       `yaml:"wicket"`
	Replacements Replacements `yaml:"replacements"`
}

type Replacements struct {
	Role  []map[string]string `yaml:"role"`
	Match []map[string]string `yaml:"match"`
}

type Runs struct {
	Batsman     int16 `yaml:"batsman"`
	Extras      int16 `yaml:"extras"`
	Total       int16 `yaml:"total"`
	NonBoundary int16 `yaml:"non_boundary"`
}

type Wicket struct {
	Fielders  []string `yaml:"fielders"`
	Kind      string   `yaml:"kind"`
	PlayerOut string   `yaml:"player_out"`
}

type Extras struct {
	Wides   int `yaml:"wides"`
	Byes    int `yaml:"byes"`
	LegByes int `yaml:"legbyes"`
	NoBalls int `yaml:"noballs"`
	Penalty int `yaml:"penalty"`
}

func main() {
	dirName := "data/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yaml") {
			filename := fmt.Sprintf("%v%v", dirName, f.Name())

			file, err := os.Open(filename)
			if err != nil {
				fmt.Println(err)
				break
			}

			b, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
				break
			}

			g := Game{}

			err = yaml.UnmarshalStrict(b, &g)
			if err != nil {
				fmt.Println(filename)
				fmt.Println(err)
				break
			}

			// fmt.Println(g)
		}
	}
}
