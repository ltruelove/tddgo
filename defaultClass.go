package main

type DefaultClass struct {
}

func (defaultClass DefaultClass) IsRollIncrementedBy1() bool {
	return false
}

func (defaultClass DefaultClass) HitPointsPerLevel() int {
	return HIT_POINTS_PER_LEVEL
}

func (defaultClass DefaultClass) IsTripleDamageOnCrit() bool {
	return false
}

func (defaultClass DefaultClass) IgnoreDexterityModWhenAttacking() bool {
	return false
}

func (defaultClass DefaultClass) UseDexterityForStrengthModifier() bool {
	return false
}

func (defaultClass DefaultClass) MinimumDamage() int {
	return MINIMUM_DAMAGE
}

func (defaultClass DefaultClass) AddsWisdomModifierToArmorClass() bool {
	return false
}

func (defaultClass DefaultClass) IncreaseAttackRollBy1EverySecondAndThirdLevel() bool {
	return false
}
