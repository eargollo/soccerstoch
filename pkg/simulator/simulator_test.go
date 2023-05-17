package simulator_test

import (
	"reflect"
	"testing"

	"github.com/eargollo/soccrtstoch/pkg/simulator"
	"github.com/stretchr/testify/mock"
)

type TestProvider struct {
	mock.Mock
}

func (o *TestProvider) Teams() ([]simulator.Team, error) {
	args := o.Called()
	return args.Get(0).([]simulator.Team), args.Error(1)
}

func (o *TestProvider) PlayedMatches() ([]simulator.Match, error) {
	args := o.Called()
	return args.Get(0).([]simulator.Match), args.Error(1)
}

func (o *TestProvider) PendingMatches() ([]simulator.Match, error) {
	args := o.Called()
	return args.Get(0).([]simulator.Match), args.Error(1)
}

// func Test_simple3(t *testing.T) {
// 	expected := simulator.SimlationResult{
// 		Probabilities: map[simulator.Team][]float64{
// 			{ID: "1", Name: "A"}: {
// 				33.33,
// 				33.33,
// 				33.33,
// 			},
// 			{ID: "2", Name: "B"}: {
// 				33.33,
// 				33.33,
// 				33.33,
// 			},
// 			{ID: "3", Name: "C"}: {
// 				33.33,
// 				33.33,
// 				33.33,
// 			},
// 		},
// 	}

// 	testSimple3 := new(TestProvider)
// 	testSimple3.On("Teams").Return(
// 		[]simulator.Team{
// 			{ID: "1", Name: "A"},
// 			{ID: "2", Name: "B"},
// 			{ID: "3", Name: "C"},
// 		}, nil,
// 	)
// 	testSimple3.On("PlayedMatches").Return(
// 		[]simulator.Match{}, nil,
// 	)
// 	testSimple3.On("PendingMatches").Return(
// 		[]simulator.Match{}, nil,
// 	)

// 	sim := simulator.New(simulator.WithChampionshipProvider(testSimple3))

// 	res, err := sim.Run()
// 	if err != nil {
// 		t.Errorf("Error simulating: %v", err)
// 	}

// 	if !reflect.DeepEqual(res, expected) {
// 		t.Errorf("simulation with 3 teams = %v, want %v", res, expected)
// 	}
// }

func TestTeamsStats_AddMatch(t *testing.T) {
	tests := []struct {
		name    string
		ts      *simulator.TeamsStats
		arg     simulator.Match
		wantErr bool
		wantTs  *simulator.TeamsStats
	}{
		{
			name: "Home wins",
			ts: &simulator.TeamsStats{
				"A": simulator.TeamStat{},
				"B": simulator.TeamStat{},
			},
			arg:     simulator.Match{LocalID: "A", AwayID: "B", LocalScore: 3, AwayScore: 2},
			wantErr: false,
			wantTs: &simulator.TeamsStats{
				"A": simulator.TeamStat{HomeWins: 1, HomeGoals: 3, HomeGoaled: 2},
				"B": simulator.TeamStat{AwayLosses: 1, AwayGoals: 2, AwayGoaled: 3},
			},
		},
		{
			name: "Draw",
			ts: &simulator.TeamsStats{
				"A": simulator.TeamStat{},
				"B": simulator.TeamStat{},
			},
			arg:     simulator.Match{LocalID: "A", AwayID: "B", LocalScore: 8, AwayScore: 8},
			wantErr: false,
			wantTs: &simulator.TeamsStats{
				"A": simulator.TeamStat{HomeDraws: 1, HomeGoals: 8, HomeGoaled: 8},
				"B": simulator.TeamStat{AwayDraws: 1, AwayGoals: 8, AwayGoaled: 8},
			},
		},
		{
			name: "Away wins",
			ts: &simulator.TeamsStats{
				"A": simulator.TeamStat{},
				"B": simulator.TeamStat{},
			},
			arg:     simulator.Match{LocalID: "A", AwayID: "B", LocalScore: 3, AwayScore: 5},
			wantErr: false,
			wantTs: &simulator.TeamsStats{
				"A": simulator.TeamStat{HomeLosses: 1, HomeGoals: 3, HomeGoaled: 5},
				"B": simulator.TeamStat{AwayWins: 1, AwayGoals: 5, AwayGoaled: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ts.AddMatch(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("TeamsStats.AddMatch() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.ts, tt.wantTs) {
				t.Errorf("simulation with 3 teams = %v, want %v", tt.ts, tt.wantTs)
			}
		})
	}
}
