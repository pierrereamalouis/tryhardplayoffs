package mysportsfeeds

// MySportsFeeds Optional parameters
type QueryParams struct {
	Player   string `url:"player"`
	Position string `url:"position"`
	Country  string `url:"country"`
	Team     string `url:"team"`
	Date     string `url:"date"`
	Stats    string `url:"stats"`
	Offset   string `url:"offset"`
	Limit    string `url:"limit"`
	Force    string `url:"force"`
}

type Endpoints string

const (
	Players       Endpoints = "https://api.mysportsfeeds.com/v2.1/pull/nhl/players.json"
	PlayoffsStats Endpoints = "https://api.mysportsfeeds.com/v2.1/pull/nhl/2024-playoff/player_stats_totals.json"
)
