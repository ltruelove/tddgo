package main

type Class interface {
	IsRollIncrementedBy1() bool
	HitPointsPerLevel() int
}

type Fighter struct {
}

func (fighter Fighter) IsRollIncrementedBy1() bool {
	return true
}

func (fighter Fighter) HitPointsPerLevel() int {
	return 10
}
