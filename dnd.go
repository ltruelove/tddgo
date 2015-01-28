package main

import (
	"fmt"
)

func main() {
	myChar := NewCharacter()
	myChar.Name = "Whatever"
	fmt.Println(myChar.Name)
}

const CRITICAL_MULTIPLIER = 2
const MINIMUM_DAMAGE = 1
const EXPERIENCE_INCREMENTOR = 10

func Attack(attacker *Character, defender *Character, dieRoll int, isCritical bool) bool {
	strengthModifier := GetStrengthModifier(attacker, isCritical)
	damageAmount := GetDamageAmount(isCritical, strengthModifier, attacker)

	dieRoll += strengthModifier

	canHit := false
	if dieRoll >= defender.CalculatedArmorClass(attacker) {
		canHit = true
		defender.TakeDamage(MaxInt(MINIMUM_DAMAGE, damageAmount))
		GainExperience(attacker, EXPERIENCE_INCREMENTOR)
	}

	return canHit
}

func GetStrengthModifier(character *Character, isCritical bool) int {
	abilityToUse := 0
	if character.Class.UseDexterityForStrengthModifier() {
		abilityToUse = character.Abilities.Dexterity
	} else {
		abilityToUse = character.Abilities.Strength
	}

	strengthModifier := GetAbilityModifier(abilityToUse)
	if isCritical {
		strengthModifier = strengthModifier * CRITICAL_MULTIPLIER
	}
	return strengthModifier
}

func GetDamageAmount(isCritical bool, strengthModifier int, attacker *Character) int {
	damageAmount := attacker.Class.MinimumDamage()
	if isCritical {
		if attacker.Class.IsTripleDamageOnCrit() {
			damageAmount = damageAmount * 3
		}
		damageAmount = damageAmount * CRITICAL_MULTIPLIER
	}
	damageAmount += strengthModifier
	return damageAmount
}

func RollDie(returnValue int, roller *Character) (int, bool) {
	isCritical := false
	if returnValue == CRIT_NATURAL_ROLL {
		isCritical = true
	}

	if roller.Class.IncreaseAttackRollBy1EverySecondAndThirdLevel() {
		level := roller.CalculatedLevel()
		if level >= 3 {
			returnValue += 2
		} else if level == 2 {
			returnValue += 1
		}
	}
	//for now, users must pass in a -1 if they don't want to specify the value they want back
	if returnValue != USE_PARAMETER {
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
