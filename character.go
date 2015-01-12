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
	newCharacter := Character{"",
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

func (character *Character) CalculatedArmorClass() int {
	dexterityModifier := GetAbilityModifier(character.Abilities.Dexterity)
	return character.ArmorClass + dexterityModifier
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
	roll, isCritical := RollDie(returnValue)
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

const CRITICAL_MULTIPLIER = 2
const MINIMUM_DAMAGE = 1
const EXPERIENCE_INCREMENTOR = 10

func Attack(attacker *Character, defender *Character, dieRoll int, isCritical bool) bool {
	strengthModifier := GetStrengthModifier(attacker, isCritical)
	damageAmount := GetDamageAmount(isCritical, strengthModifier)

	dieRoll += strengthModifier

	canHit := false
	if dieRoll >= defender.ArmorClass {
		canHit = true
		defender.TakeDamage(MaxInt(MINIMUM_DAMAGE, damageAmount))
		GainExperience(attacker, EXPERIENCE_INCREMENTOR)
	}

	return canHit
}

func GetStrengthModifier(character *Character, isCritical bool) int {
	strengthModifier := GetAbilityModifier(character.Abilities.Strength)
	if isCritical {
		strengthModifier = strengthModifier * CRITICAL_MULTIPLIER
	}
	return strengthModifier
}

func GetDamageAmount(isCritical bool, strengthModifier int) int {
	damageAmount := MINIMUM_DAMAGE
	if isCritical {
		damageAmount = damageAmount * CRITICAL_MULTIPLIER
	}
	damageAmount += strengthModifier
	return damageAmount
}

func RollDie(returnValue int) (int, bool) {
	isCritical := false
	//for now, users must pass in a -1 if they don't want to specify the value they want back
	if returnValue != USE_PARAMETER {
		if returnValue == CRIT_NATURAL_ROLL {
			isCritical = true
		}
		return returnValue, isCritical
	}
	//@todo return a random value between 1 and 20 and test for crit
	return TEMPORARY_ROLL_VALUE, isCritical
}

func MaxInt(x int, y int) int {
	if x > y {
		return x
	} else if y > x {
		return y
	} else {
		return x
	}
}
