package main

type DefaultClass struct {
}

func (defaultClass DefaultClass) IsRollIncrementedBy1() bool {
	return false
}

func (defaultClass DefaultClass) HitPointsPerLevel() int {
	return HIT_POINTS_PER_LEVEL
}
