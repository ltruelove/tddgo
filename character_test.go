package main

import (
	"testing"
)

func TestCharacterName(t *testing.T) {
	myChar := NewCharacter()
	myChar.Name = "Ted"

	if myChar.Name != "Ted" {
		t.Errorf("Character name: '%s' does not match when testing Character Name.", myChar.Name)
	} else {
		t.Log("Character name matches")
	}
}

func TestCharacterAlignment(t *testing.T) {
	myChar := NewCharacter()
	myChar.Alignment = ALIGN_NEUTRAL

	if myChar.Alignment == "Neutral" {
		t.Log("Character alignment matches")
	} else {
		t.Error("Character alignment does not match")
	}

}

func TestCharacterDefaultArmor(t *testing.T) {
	myChar := NewCharacter()

	if myChar.ArmorClass == 10 {
		t.Log("Default armor class test passes")
	} else {
		t.Error("Default armor class test failed")
	}
}

func TestCharacterDefaultHitPoints(t *testing.T) {
	myChar := NewCharacter()

	if myChar.HitPoints == 5 {
		t.Log("Default hit points test passes")
	} else {
		t.Error("Default hit points test failed")
	}
}

func TestCharacterRollWithHit(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	dieRoll, isCritical := RollDie(20)

	wasAbleToAttack := attacker.Attack(&defender, dieRoll, isCritical)

	if wasAbleToAttack {
		t.Log("Attack succeeded")
	} else {
		t.Error("Attack expected to succeed but did not")
	}
}

func TestCharacterRollWithHitDoesDamage(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	dieRoll, isCritical := RollDie(10)

	attacker.Attack(&defender, dieRoll, isCritical)

	if defender.HitPoints == 4 {
		t.Log("Attack succeeded, hit points deducted")
	} else {
		t.Errorf("Attack expected to succeed but did not. Hit points remain at %d", defender.HitPoints)
	}
}

func TestCharacterRollWithCriticalHit(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	dieRoll, isCritical := RollDie(20)

	attacker.Attack(&defender, dieRoll, isCritical)

	if defender.HitPoints == 3 {
		t.Log("Attack succeeded, double hit points deducted")
	} else {
		t.Errorf("Attack expected to succeed but did not. Hit points: %d", defender.HitPoints)
	}
}

func TestCharacterIsAliveUntilZero(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	dieRoll, isCritical := RollDie(20)

	attacker.Attack(&defender, dieRoll, isCritical)
	attacker.Attack(&defender, dieRoll, isCritical)
	attacker.Attack(&defender, dieRoll, isCritical)

	if defender.IsAlive() {
		t.Errorf("Defender still lives with %d hit points", defender.HitPoints)
	} else {
		t.Log("Defender is dead")
	}
}

func TestCharacterAttackWithStrengthModifier(t *testing.T) {
	attacker := NewCharacter()
	attacker.Abilities.Strength = 15

	defender := NewCharacter()
	dieRoll, isCritical := RollDie(10)
	expectedHP := 2

	attacker.Attack(&defender, dieRoll, isCritical)

	if defender.HitPoints != expectedHP {
		t.Errorf("Defender hit points incorrect. Expected %d, and got %d", expectedHP, defender.HitPoints)
	} else {
		t.Log("Defender hit points correctly calculated after an attack with strength modifier")
	}
}

func TestCharacterAttackCritWithStrengthModifier(t *testing.T) {
	attacker := NewCharacter()
	attacker.Abilities.Strength = 15

	defender := NewCharacter()
	dieRoll, isCritical := RollDie(20)
	expectedHP := -1

	attacker.Attack(&defender, dieRoll, isCritical)

	if defender.HitPoints != expectedHP {
		t.Errorf("Defender hit points incorrect. Expected %d, and got %d", expectedHP, defender.HitPoints)
	} else {
		t.Log("Defender hit points correctly calculated after a critical attack with strength modifier")
	}
}

func TestCharacterAttackMinimumDamage(t *testing.T) {
	attacker := NewCharacter()
	attacker.Abilities.Strength = 1

	defender := NewCharacter()
	dieRoll, isCritical := RollDie(15)
	expectedHP := 4

	attacker.Attack(&defender, dieRoll, isCritical)

	if defender.HitPoints != expectedHP {
		t.Errorf("Defender hit points incorrect. Expected %d, and got %d", expectedHP, defender.HitPoints)
	} else {
		t.Log("Defender hit points correctly calculated after a negative strength modifier and a roll high enough to attack")
	}
}
