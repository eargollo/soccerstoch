package simulator

type ChampionshipProvicer interface {
	Teams() ([]Team, error)
	PlayedMatches() ([]Match, error)
	PendingMatches() ([]Match, error)
}

type Team struct {
	ID   string
	Name string
}

type Match struct {
	LocalID    string
	AwayID     string
	LocalScore int
	AwayScore  int
}
