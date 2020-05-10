package cricket

// Game Defines a game of cricket
type Game struct {
	GameID  string
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
