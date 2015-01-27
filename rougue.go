package main

type Rogue struct {
}

func (rogue Rogue) IsRollIncrementedBy1() bool {
	return false
}

func (rogue Rogue) HitPointsPerLevel() int {
	return HIT_POINTS_PER_LEVEL
}

func (rogue Rogue) IsTripleDamageOnCrit() bool {
	return true
}

func (rogue Rogue) IgnoreDexterityModWhenAttacking() bool {
	return true
}
