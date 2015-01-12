package main

type DefaultClass struct {
}

func (defaultClass DefaultClass) IsRollIncrementedBy1() bool {
	return false
}

const HIT_POINTS_PER_LEVEL int = 5

func (defaultClass DefaultClass) HitPointsPerLevel() int {
	return HIT_POINTS_PER_LEVEL
}
