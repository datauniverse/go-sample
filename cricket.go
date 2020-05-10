package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// Game Defines a game of cricket
type Game struct {
	Meta    Metadata            `yaml:"meta"`
	Info    Information         `yaml:"info"`
	Innings []map[string]Inning `yaml:"innings"`
}

// Metadata Technical data related to a particular Game
type Metadata struct {
	DataVersion float32 `yaml:"data_version"`
	Created     string  `yaml:"created"`
	Revision    int16   `yaml:"revision"`
}

// Information Info related to a particular Game
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

	Supersubs map[string]string   `yaml:"supersubs"`
	BowlOut   []map[string]string `yaml:"bowl_out"`

	Outcome Outcome `yaml:"outcome"`
	Toss    Toss    `yaml:"toss"`
}

// Toss Toos details of the Game
type Toss struct {
	Decision string `yaml:"decision"`
	Winner   string `yaml:"winner"`
}

// Outcome Outcome of a particular Game
type Outcome struct {
	BowlOut    string `yaml:"bowl_out"`
	Eliminator string `yaml:"eliminator"`
	Result     string `yaml:"result"`
	Method     string `yaml:"method"`
	Winner     string `yaml:"winner"`
	By         By     `yaml:"by"`
}

// By More details related to Outcome of a particular Game
type By struct {
	Innings string `yaml:"innings"`
	Runs    int16  `yaml:"runs"`
	Wickets int16  `yaml:"wickets"`
}

// Inning Details of Innnings played in a particular Game
type Inning struct {
	Team        string                `yaml:"team"`
	Declared    string                `yaml:"declared"`
	AbsentHurt  []string              `yaml:"absent_hurt"`
	PenaltyRuns PenaltyRuns           `yaml:"penalty_runs"`
	Deliveries  []map[string]Delivery `yaml:"deliveries"`
}

// PenaltyRuns Penalties awarded in a particular Game
type PenaltyRuns struct {
	Pre  int16 `yaml:"pre"`
	Post int16 `yaml:"post"`
}

// Delivery Details of a Delivery that was bowled in a particular Game
type Delivery struct {
	Batsman      string       `yaml:"batsman"`
	Bowler       string       `yaml:"bowler"`
	NonStriker   string       `yaml:"non_striker"`
	Extras       Extras       `yaml:"extras"`
	Runs         Runs         `yaml:"runs"`
	Wicket       Wicket       `yaml:"wicket"`
	Replacements Replacements `yaml:"replacements"`
}

// Replacements made if any in a particular Game
type Replacements struct {
	Role  []map[string]string `yaml:"role"`
	Match []map[string]string `yaml:"match"`
}

// Runs scored off a Delivery in a particular Game
type Runs struct {
	Batsman     int16 `yaml:"batsman"`
	Extras      int16 `yaml:"extras"`
	Total       int16 `yaml:"total"`
	NonBoundary int16 `yaml:"non_boundary"`
}

// Wicket details of a Wicket taken on a particular Delivery in a particular Game
type Wicket struct {
	Fielders  []string `yaml:"fielders"`
	Kind      string   `yaml:"kind"`
	PlayerOut string   `yaml:"player_out"`
}

// Extras given off of a particular Delivery in a particular Game
type Extras struct {
	Wides   int `yaml:"wides"`
	Byes    int `yaml:"byes"`
	LegByes int `yaml:"legbyes"`
	NoBalls int `yaml:"noballs"`
	Penalty int `yaml:"penalty"`
}

func writeCSV(filename string, data [][]string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, d := range data {
		err := writer.Write(d)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	start := time.Now()
	log.Println("Started execution:", start)

	dirName := "data/all/"
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	games := [][]string{}
	games = append(games, []string{
		"GameId", "City", "StartDate", "Gender", "MatchType",
		"MatchTypeNumber", "NeutralVenue",
		"Winner", "Result", "OutcomeMethod", "Eliminator",
		"OutcomeByInnings", "OutcomeByRuns",
		"OutcomeByWickets", "OutcomeBowlOut",
		"Overs", "PlayerOfMatch", "TeamOne", "TeamTwo",
		"TossDecision", "TossWinner", "UmpireOne", "UmpireTwo", "Venue",
	})

	deliveries := [][]string{}
	deliveries = append(deliveries, []string{
		"GameId", "Batsman", "Bowler", "NonStriker",
		"RunsBatsman", "RunsExtras",
		"RunsNonBoundary", "RunsTotal",
		"ExtrasByes", "ExtrasLegByes",
		"ExtrasNoBalls", "ExtrasPenalty",
		"ExtrasWides",
		"WicketKind", "WicketFielders", "WicketPlayerOut",
	})

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

			var pom, u1, u2, d, t1, t2 string
			if len(g.Info.PlayerOfMatch) > 0 {
				pom = g.Info.PlayerOfMatch[0]
			}
			if len(g.Info.Umpires) > 0 {
				u1 = g.Info.Umpires[0]
			} else if len(g.Info.Umpires) == 2 {
				u2 = g.Info.Umpires[1]
			}
			if len(g.Info.Dates) > 0 {
				d = g.Info.Dates[0]
			}
			if len(g.Info.Teams) > 0 {
				t1 = g.Info.Teams[0]
			} else if len(g.Info.Teams) == 2 {
				t2 = g.Info.Teams[1]
			}

			game := []string{
				strings.Split(f.Name(), ".")[0], g.Info.City, d, g.Info.Gender, g.Info.MatchType,
				strconv.Itoa(int(g.Info.MatchTypeNumber)), strconv.Itoa(int(g.Info.NeutralVenue)),
				g.Info.Outcome.Winner, g.Info.Outcome.Result, g.Info.Outcome.Method, g.Info.Outcome.Eliminator,
				g.Info.Outcome.By.Innings, strconv.Itoa(int(g.Info.Outcome.By.Runs)),
				strconv.Itoa(int(g.Info.Outcome.By.Wickets)), g.Info.Outcome.BowlOut,
				strconv.Itoa(int(g.Info.Overs)), pom, t1, t2,
				g.Info.Toss.Decision, g.Info.Toss.Winner, u1, u2, g.Info.Venue,
			}
			games = append(games, game)

			for _, i := range g.Innings {
				for _, in := range i {
					for _, d := range in.Deliveries {
						for _, del := range d {
							var wf string
							if len(del.Wicket.Fielders) > 0 {
								wf = del.Wicket.Fielders[0]
							}

							delivery := []string{
								strings.Split(f.Name(), ".")[0], del.Batsman, del.Bowler, del.NonStriker,
								strconv.Itoa(int(del.Runs.Batsman)), strconv.Itoa(int(del.Runs.Extras)),
								strconv.Itoa(int(del.Runs.NonBoundary)), strconv.Itoa(int(del.Runs.Total)),
								strconv.Itoa(del.Extras.Byes), strconv.Itoa(del.Extras.LegByes),
								strconv.Itoa(del.Extras.NoBalls), strconv.Itoa(del.Extras.Penalty),
								strconv.Itoa(del.Extras.Wides),
								del.Wicket.Kind, wf, del.Wicket.PlayerOut,
							}

							deliveries = append(deliveries, delivery)
						}
					}
				}
			}
		}
	}
	writeCSV("data/games.csv", games)
	writeCSV("data/deliveries.csv", deliveries)

	end := time.Now()
	log.Println("Ended execution:", end)
	log.Println("Total execution time in seconds:", end.Sub(start).Seconds())
}
