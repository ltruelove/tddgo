package main

const HIT_POINTS_PER_LEVEL int = 5

type Class interface {
	IsRollIncrementedBy1() bool
	HitPointsPerLevel() int
	IsTripleDamageOnCrit() bool
	IgnoreDexterityModWhenAttacking() bool
	UseDexterityForStrengthModifier() bool
	MinimumDamage() int
	AddsWisdomModifierToArmorClass() bool
	IncreaseAttackRollBy1EverySecondAndThirdLevel() bool
}
