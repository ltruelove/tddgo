package main

import (
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
	Abilities  Abilities
}

func NewCharacter() Character {
	newCharacter := Character{"", "", 10, 5, NewAbilities()}
	return newCharacter
}

func (c Character) Attack(defender *Character, dieRoll int, isCritical bool) bool {
	canHit := false
	damageAmount := 1

	strengthModifier := GetAbilityModifier(c.Abilities.Strength)

	if isCritical {
		damageAmount = damageAmount * 2
		strengthModifier = strengthModifier * 2
	}

	damageAmount += strengthModifier
	dieRoll += strengthModifier

	if dieRoll >= defender.ArmorClass {
		canHit = true
		defender.TakeDamage(MaxInt(1, damageAmount))
	}

	return canHit
}

func RollDie(returnValue int) (int, bool) {
	isCritical := false
	if returnValue != -1 {
		//if it's a 'natural' 20 then it's a critical hit
		if returnValue == 20 {
			isCritical = true
		}
		return returnValue, isCritical
	}

	//@todo return a random value between 1 and 20 and test for crit
	return 10, isCritical
}

func (c *Character) TakeDamage(damageAmount int) {
	c.HitPoints -= damageAmount
}

func (c Character) IsAlive() bool {
	return c.HitPoints > 0
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