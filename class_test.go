package main

import (
	"testing"
)

func TestFighterIncreasesRollBy1ForEveryLevel(t *testing.T) {
	attacker := NewCharacter()
	attacker.Class = Fighter{}
	attacker.Experience = 5000
	dieRoll, _ := attacker.CalculatedRoll(15)
	expectedRoll := 21

	if dieRoll != expectedRoll {
		t.Errorf("Attacker roll was %d, test expected %d", dieRoll, expectedRoll)
	} else {
		t.Log("Even levels add 1 to the die roll even with higher levels")
	}
}

func TestFighterHitPointsByLevel(t *testing.T) {
	attacker := NewCharacter()
	attacker.Class = Fighter{}
	attacker.Experience = 5000
	expectedHitPoints := 60
	hitPoints := attacker.CalculatedHitPoints()

	if hitPoints != expectedHitPoints {
		t.Errorf("Hit points were %d, expected %d", hitPoints, expectedHitPoints)
	} else {
		t.Log("Hit points matched the expected value")
	}
}

func TestRogueTriplesDamageOnCriticalHits(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	attacker.Class = Rogue{}
	defender.HitPoints = 7
	expectedHitPoints := 1

	dieRoll, isCritical := RollDie(20)
	Attack(attacker, defender, dieRoll, isCritical)

	if defender.HitPoints != expectedHitPoints {
		t.Errorf("Hit points were %d, expected %d", defender.HitPoints, expectedHitPoints)
	} else {
		t.Log("Hit points matched the expected value")
	}
}

func TestRogueIgnoresDexterityModOnAttack(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	attacker.Class = Rogue{}
	defender.HitPoints = 7
	defender.Abilities.Dexterity = 20
	expectedHitPoints := 6

	dieRoll, isCritical := RollDie(10)
	Attack(attacker, defender, dieRoll, isCritical)

	if defender.HitPoints != expectedHitPoints {
		t.Errorf("Hit points were %d, expected %d", defender.HitPoints, expectedHitPoints)
	} else {
		t.Log("Hit points matched the expected value")
	}
}

func TestRogueAddsDexterityToAttack(t *testing.T) {
	attacker := NewCharacter()
	attacker.Class = Rogue{}
	attacker.Abilities.Dexterity = 20
	expectedStrengthModifier := 5

	strengthModifier := GetStrengthModifier(attacker, false)

	if strengthModifier != expectedStrengthModifier {
		t.Errorf("Strength modifier was %d, expected %d", strengthModifier, expectedStrengthModifier)
	} else {
		t.Log("Strength modifier matched the expected value")
	}
}
