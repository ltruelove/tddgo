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

	dieRoll, isCritical := RollDie(20, attacker)
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

	dieRoll, isCritical := RollDie(10, attacker)
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

func TestWarMonkHitPointsByLevel(t *testing.T) {
	attacker := NewCharacter()
	attacker.Class = WarMonk{}
	attacker.Experience = 5000
	expectedHitPoints := 36
	hitPoints := attacker.CalculatedHitPoints()

	if hitPoints != expectedHitPoints {
		t.Errorf("Hit points were %d, expected %d", hitPoints, expectedHitPoints)
	} else {
		t.Log("Hit points matched the expected value")
	}
}

func TestWarMonkDamageAmount(t *testing.T) {
	attacker := NewCharacter()
	attacker.Class = WarMonk{}
	damageAmount := GetDamageAmount(false, 0, attacker)
	expectedDamage := 3

	if damageAmount != expectedDamage {
		t.Errorf("Damage amount was %d, expected %d", damageAmount, expectedDamage)
	} else {
		t.Log("Damage amount matched the expected value")
	}
}

func TestWarMonkWisdomModifier(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	defender.Class = WarMonk{}
	defender.Abilities.Wisdom = 20

	armorClass := defender.CalculatedArmorClass(attacker)
	expectedArmorClass := 15

	if armorClass != expectedArmorClass {
		t.Errorf("Armor class was %d, expected %d", armorClass, expectedArmorClass)
	} else {
		t.Log("Armor class matched the expected value")
	}
}

func TestWarMonkWisdomModifierAgainstRogue(t *testing.T) {
	attacker := NewCharacter()
	attacker.Class = Rogue{}
	defender := NewCharacter()
	defender.Class = WarMonk{}
	defender.Abilities.Wisdom = 20
	defender.Abilities.Dexterity = 20

	armorClass := defender.CalculatedArmorClass(attacker)
	expectedArmorClass := 15

	if armorClass != expectedArmorClass {
		t.Errorf("Armor class was %d, expected %d", armorClass, expectedArmorClass)
	} else {
		t.Log("Armor class matched the expected value")
	}
}

func TestWarMonkNegativeWisdomModifier(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	defender.Class = WarMonk{}
	defender.Abilities.Wisdom = -20

	armorClass := defender.CalculatedArmorClass(attacker)
	expectedArmorClass := 10

	if armorClass != expectedArmorClass {
		t.Errorf("Armor class was %d, expected %d", armorClass, expectedArmorClass)
	} else {
		t.Log("Armor class matched the expected value")
	}
}

func TestWarMonkAttackRollModByLevel(t *testing.T) {
	attacker := NewCharacter()
	attacker.Class = WarMonk{}
	GainExperience(attacker, 1001)
	dieRoll, _ := RollDie(10, attacker)
	expectedDieRoll := 11

	if dieRoll != expectedDieRoll {
		t.Errorf("Die roll was %d, expected %d", dieRoll, expectedDieRoll)
	} else {
		t.Log("Die roll for level 2 matched the expected value")
	}

	GainExperience(attacker, 1001)
	dieRoll, _ = RollDie(10, attacker)
	expectedDieRoll = 12

	if dieRoll != expectedDieRoll {
		t.Errorf("Die roll was %d, expected %d", dieRoll, expectedDieRoll)
	} else {
		t.Log("Die roll for level 3 matched the expected value")
	}

	GainExperience(attacker, 1001)
	dieRoll, _ = RollDie(10, attacker)
	expectedDieRoll = 12

	if dieRoll != expectedDieRoll {
		t.Errorf("Die roll was %d, expected %d", dieRoll, expectedDieRoll)
	} else {
		t.Log("Die roll for level 4 matched the expected value")
	}
}
