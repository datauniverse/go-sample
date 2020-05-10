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

	"./model/cricket"
)

func readDataDirectory(directoryPath string) []cricket.Game {
	var games []cricket.Game

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".yaml") {
			filename := fmt.Sprintf("%v%v", directoryPath, f.Name())

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

			g := cricket.Game{}
			g.GameID = strings.Split(f.Name(), ".")[0]

			err = yaml.UnmarshalStrict(b, &g)
			if err != nil {
				log.Println(filename)
				log.Println(err)
			}

			games = append(games, g)
		}
	}
	return games
}

func getGame(games []cricket.Game) [][]string {
	gs := [][]string{}
	gs = append(gs, []string{
		"GameId", "City", "StartDate", "Gender", "MatchType",
		"MatchTypeNumber", "NeutralVenue",
		"Winner", "Result", "OutcomeMethod", "Eliminator",
		"OutcomeByInnings", "OutcomeByRuns",
		"OutcomeByWickets", "OutcomeBowlOut",
		"Overs", "PlayerOfMatch[0]", "TeamOne", "TeamTwo",
		"TossDecision", "TossWinner", "UmpireOne", "UmpireTwo", "Venue",
	})

	for _, g := range games {
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

		gTemp := []string{
			g.GameID, g.Info.City, d, g.Info.Gender, g.Info.MatchType,
			strconv.Itoa(int(g.Info.MatchTypeNumber)), strconv.Itoa(int(g.Info.NeutralVenue)),
			g.Info.Outcome.Winner, g.Info.Outcome.Result, g.Info.Outcome.Method, g.Info.Outcome.Eliminator,
			g.Info.Outcome.By.Innings, strconv.Itoa(int(g.Info.Outcome.By.Runs)),
			strconv.Itoa(int(g.Info.Outcome.By.Wickets)), g.Info.Outcome.BowlOut,
			strconv.Itoa(int(g.Info.Overs)), pom, t1, t2,
			g.Info.Toss.Decision, g.Info.Toss.Winner, u1, u2, g.Info.Venue,
		}

		gs = append(gs, gTemp)
	}

	return gs
}

func getDeliveries(games []cricket.Game) [][]string {
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

	for _, game := range games {
		for _, i := range game.Innings {
			for _, in := range i {
				for _, d := range in.Deliveries {
					for _, del := range d {
						var wf string
						if len(del.Wicket.Fielders) > 0 {
							wf = del.Wicket.Fielders[0]
						}

						delivery := []string{
							game.GameID, del.Batsman, del.Bowler, del.NonStriker,
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

	return deliveries
}

func writeStringMatrixToCSVFile(filepath string, data [][]string) {
	file, err := os.Create(filepath)
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

	games := readDataDirectory("data/all/")
	gamesMatrix := getGame(games)
	deliveriesMatrix := getDeliveries(games)

	writeStringMatrixToCSVFile("data/games.csv", gamesMatrix)
	writeStringMatrixToCSVFile("data/deliveries.csv", deliveriesMatrix)

	end := time.Now()
	log.Println("Ended execution:", end)
	log.Println("Total execution time in seconds:", end.Sub(start).Seconds())
}
