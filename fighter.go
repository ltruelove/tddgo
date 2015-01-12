package main

type Fighter struct {
}

func (fighter Fighter) IsRollIncrementedBy1() bool {
	return true
}

func (fighter Fighter) HitPointsPerLevel() int {
	return HIT_POINTS_PER_LEVEL + 5
}
