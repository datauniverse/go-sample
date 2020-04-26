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
	Meta    Metadata          `yaml:"meta"`
	Info    Information       `yaml:"info"`
	Innings map[string]Inning `yaml:"innings"`
	/*
		Innings []struct {
			Inning struct {
				Team       string `yaml:"team"`
				Deliveries []struct {
					Delivery struct {
						Batsman string `yaml:"batsman"`
					}
				}
			}
		}
	*/
}

type Metadata struct {
	DataVersion float32 `yaml:"data_version"`
	Created     string  `yaml:"created"`
	Revision    int16   `yaml:"revision"`
}

type Information struct {
	City      string   `yaml:"city"`
	Dates     []string `yaml:"dates"`
	Gender    string   `yaml:"gender"`
	MatchType string   `yaml:"match_type"`
	Outcome   struct {
		By struct {
			Runs int16 `yaml:"runs"`
		}
		Winner string `yaml:"winner"`
	}
	Overs         int16    `yaml:"overs"`
	PlayerOfMatch []string `yaml:"player_of_match"`
	Teams         []string `yaml:"teams"`
	Toss          struct {
		Decision string `yaml:"decision"`
		Winner   string `yaml:"winner"`
	}
	Umpires []string `yaml:"umpires"`
	Venue   string   `yaml:"venue"`
}

type Inning struct {
}

func main() {
	dirName := "C:\\Users\\abhil\\OneDrive\\Desktop\\all\\"
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
				fmt.Println(err)
				break
			}

			fmt.Println(g)
		}
	}
}
