package main

type Fighter struct {
}

func (fighter Fighter) IsRollIncrementedBy1() bool {
	return true
}

func (fighter Fighter) HitPointsPerLevel() int {
	return HIT_POINTS_PER_LEVEL + 5
}

func (fighter Fighter) IsTripleDamageOnCrit() bool {
	return false
}

func (fighter Fighter) IgnoreDexterityModWhenAttacking() bool {
	return false
}

func (fighter Fighter) UseDexterityForStrengthModifier() bool {
	return false
}
