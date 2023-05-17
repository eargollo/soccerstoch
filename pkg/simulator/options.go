package simulator

type SimulatorOption func(*Simulator)

var defaults = []SimulatorOption{
	WithPointSystem(3, 1, 0),
}

func WithChampionshipProvider(prov ChampionshipProvicer) SimulatorOption {
	return func(s *Simulator) {
		s.provider = prov
	}
}

func WithPointSystem(win int, draw int, loss int) SimulatorOption {
	return func(s *Simulator) {
		s.points.win = win
		s.points.draw = draw
		s.points.loss = loss
	}
}
