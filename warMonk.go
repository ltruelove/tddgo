package main

type WarMonk struct {
}

func (monk WarMonk) IsRollIncrementedBy1() bool {
	return false
}

func (monk WarMonk) HitPointsPerLevel() int {
	return HIT_POINTS_PER_LEVEL + 1
}

func (monk WarMonk) IsTripleDamageOnCrit() bool {
	return false
}

func (monk WarMonk) IgnoreDexterityModWhenAttacking() bool {
	return false
}

func (monk WarMonk) UseDexterityForStrengthModifier() bool {
	return false
}

func (monk WarMonk) MinimumDamage() int {
	return MINIMUM_DAMAGE + 2
}

func (monk WarMonk) AddsWisdomModifierToArmorClass() bool {
	return true
}

func (monk WarMonk) IncreaseAttackRollBy1EverySecondAndThirdLevel() bool {
	return true
}
