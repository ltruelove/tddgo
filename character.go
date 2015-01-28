package main

import (
	"math"
	//"fmt"
)

const ALIGN_GOOD string = "Good"
const ALIGN_EVIL string = "Evil"
const ALIGN_NEUTRAL string = "Neutral"

type Character struct {
	Name       string
	Alignment  string
	ArmorClass int
	HitPoints  int
	Experience int
	Level      int
	Abilities  Abilities
	Class      Class
}

func NewCharacter() *Character {
	newCharacter := Character{
		"",
		"",
		10,
		HIT_POINTS_PER_LEVEL,
		0,
		1,
		NewAbilities(),
		DefaultClass{}}
	return &newCharacter
}

func GainExperience(character *Character, amountToIncrement int) {
	currentLevel := character.CalculatedLevel()
	character.Experience += amountToIncrement
	newLevel := character.CalculatedLevel()
	if newLevel > currentLevel {
		character.LevelUp()
	}
}

func (character *Character) LevelUp() {
	character.HitPoints = character.CalculatedHitPoints()
}

func (character *Character) CalculatedArmorClass(attacker *Character) int {
	armorModifier := 0
	if character.Class.AddsWisdomModifierToArmorClass() {
		armorModifier += MaxInt(0, GetAbilityModifier(character.Abilities.Wisdom))
	}

	dexterityModifier := GetAbilityModifier(character.Abilities.Dexterity)
	if attacker.Class.IgnoreDexterityModWhenAttacking() {
		if dexterityModifier < 0 {
			armorModifier += dexterityModifier
		}
	} else {
		armorModifier += dexterityModifier
	}

	return character.ArmorClass + armorModifier
}

const EXPERIENCE_STEP float64 = 1000.0

func (character *Character) CalculatedLevel() int {
	experienceIncrementor := math.Floor(float64(character.Experience) / EXPERIENCE_STEP)
	return character.Level + int(experienceIncrementor)
}

func (character *Character) CalculatedHitPoints() int {
	constitutionModifier := GetAbilityModifier(character.Abilities.Constitution)
	hitPointsMultiplier := character.Class.HitPointsPerLevel() + constitutionModifier
	combinedHitPoints := (hitPointsMultiplier * character.CalculatedLevel())
	return MaxInt(1, combinedHitPoints)
}

const USE_PARAMETER = -1
const CRIT_NATURAL_ROLL = 20
const TEMPORARY_ROLL_VALUE = 10

func (character *Character) CalculatedRoll(returnValue int) (int, bool) {
	roll, isCritical := RollDie(returnValue, character)
	incrementor := 0
	if character.Class.IsRollIncrementedBy1() {
		incrementor = character.CalculatedLevel()
	} else {
		incrementor = int(math.Floor(float64(character.CalculatedLevel()) / 2))
	}
	roll += incrementor
	return roll, isCritical
}

func (character *Character) TakeDamage(damageAmount int) {
	character.HitPoints -= damageAmount
}

func (character *Character) IsAlive() bool {
	return character.HitPoints > 0
}
