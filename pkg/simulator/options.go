package simulator

type SimulatorOption func(*Simulator)

var defaults = []SimulatorOption{
	WithPointSystem(PointsMap{Win: 3, Draw: 1, Loss: 0}),
}

func WithChampionshipProvider(prov ChampionshipProvicer) SimulatorOption {
	return func(s *Simulator) {
		s.provider = prov
	}
}

func WithPointSystem(points PointsMap) SimulatorOption {
	return func(s *Simulator) {
		s.points = points
	}
}
