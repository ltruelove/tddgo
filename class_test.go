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