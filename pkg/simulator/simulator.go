package simulator

type Simulator struct {
	provider ChampionshipProvicer
	points   PointsMap
}

type PointsMap struct {
	Win  int
	Draw int
	Loss int
}

func New(options ...SimulatorOption) *Simulator {
	sim := Simulator{}

	options = append(defaults, options...)

	for _, opt := range options {
		opt(&sim)
	}

	return &sim
}

type SimlationResult struct {
	Probabilities map[Team][]float64
}

type TeamStat struct {
	HomeWins   int
	HomeDraws  int
	HomeLosses int
	AwayWins   int
	AwayDraws  int
	AwayLosses int
	HomeGoals  int
	AwayGoals  int
	HomeGoaled int
	AwayGoaled int
}

type TeamsStats map[string]TeamStat

type TeamRank struct {
	Rank   int
	ID     string
	Name   string
	Points int
	TeamStat
}

type ChampionshipRank []TeamRank

func NewRank(stats TeamsStats, points PointsMap) *ChampionshipRank {
	rank := ChampionshipRank{}

	for id, stat := range stats {
		tr := TeamRank{}
		tr.TeamStat = stat
		tr.ID = id
		tr.Rank = 1
		tr.Points = (stat.HomeWins+stat.AwayWins)*points.Win +
			(stat.HomeDraws+stat.AwayDraws)*points.Draw +
			(stat.HomeLosses+stat.AwayLosses)*points.Loss

		rank = append(rank, tr)
	}
	return &rank
}

func (ts TeamsStats) AddMatch(m Match) error {
	local := ts[m.LocalID]
	away := ts[m.AwayID]

	local.HomeGoals += m.LocalScore
	away.AwayGoaled += m.LocalScore
	local.HomeGoaled += m.AwayScore
	away.AwayGoals += m.AwayScore

	// Either home win, draw or away win
	if m.LocalScore > m.AwayScore {
		// Home wins
		local.HomeWins++
		away.AwayLosses++
	}
	if m.LocalScore == m.AwayScore {
		// Draw
		local.HomeDraws++
		away.AwayDraws++
	}
	if m.LocalScore < m.AwayScore {
		// Away wins
		local.HomeLosses++
		away.AwayWins++
	}

	ts[m.LocalID] = local
	ts[m.AwayID] = away

	return nil
}

func (sim *Simulator) Run() (SimlationResult, error) {
	// played := sim.provider.PlayedMatches()

	return SimlationResult{}, nil
}
