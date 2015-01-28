package main

import (
	"testing"
	//"fmt"
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
	dieRoll, isCritical := RollDie(20, attacker)

	wasAbleToAttack := Attack(attacker, defender, dieRoll, isCritical)

	if wasAbleToAttack {
		t.Log("Attack succeeded")
	} else {
		t.Error("Attack expected to succeed but did not")
	}
}

func TestCharacterRollWithHitDoesDamage(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	dieRoll, isCritical := RollDie(10, attacker)

	Attack(attacker, defender, dieRoll, isCritical)

	if defender.HitPoints == 4 {
		t.Log("Attack succeeded, hit points deducted")
	} else {
		t.Errorf("Attack expected to succeed but did not. Hit points remain at %d", defender.HitPoints)
	}
}

func TestCharacterRollWithCriticalHit(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	dieRoll, isCritical := RollDie(20, attacker)

	Attack(attacker, defender, dieRoll, isCritical)

	if defender.HitPoints == 3 {
		t.Log("Attack succeeded, double hit points deducted")
	} else {
		t.Errorf("Attack expected to succeed but did not. Hit points: %d", defender.HitPoints)
	}
}

func TestCharacterIsAliveUntilZero(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	dieRoll, isCritical := RollDie(20, attacker)

	Attack(attacker, defender, dieRoll, isCritical)
	Attack(attacker, defender, dieRoll, isCritical)
	Attack(attacker, defender, dieRoll, isCritical)

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
	dieRoll, isCritical := RollDie(10, attacker)
	expectedHP := 2

	Attack(attacker, defender, dieRoll, isCritical)

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
	dieRoll, isCritical := RollDie(20, attacker)
	expectedHP := -1

	Attack(attacker, defender, dieRoll, isCritical)

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
	dieRoll, isCritical := RollDie(15, attacker)
	expectedHP := 4

	Attack(attacker, defender, dieRoll, isCritical)

	if defender.HitPoints != expectedHP {
		t.Errorf("Defender hit points incorrect. Expected %d, and got %d", expectedHP, defender.HitPoints)
	} else {
		t.Log("Defender hit points correctly calculated after a negative strength modifier and a roll high enough to attack")
	}
}

func TestDexterityModifierOnArmorClass(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	defender.Abilities.Dexterity = 15
	expectedArmorValue := 12
	armorClass := defender.CalculatedArmorClass(attacker)

	if armorClass != expectedArmorValue {
		t.Errorf("Defender armor class value incorrect. Expected %d, and got %d", expectedArmorValue, armorClass)
	} else {
		t.Log("Attacker armor class calculated correctly")
	}
}

func TestConstitutionModifierOnHitPoints(t *testing.T) {
	attacker := NewCharacter()
	attacker.Abilities.Constitution = 15
	expectedHitPoints := 7
	hitPoints := attacker.CalculatedHitPoints()

	if hitPoints != expectedHitPoints {
		t.Errorf("Attacker hit points value incorrect. Expected %d, and got %d", expectedHitPoints, hitPoints)
	} else {
		t.Log("Attacker hit points calculated correctly")
	}
}

func TestConstitutionModifierMinimumHitPoints(t *testing.T) {
	attacker := NewCharacter()
	attacker.Abilities.Constitution = 1
	expectedHitPoints := 1
	hitPoints := attacker.CalculatedHitPoints()

	if hitPoints != expectedHitPoints {
		t.Errorf("Attacker hit points value incorrect. Expected %d, and got %d", expectedHitPoints, hitPoints)
	} else {
		t.Log("Attacker hit points calculated correctly")
	}
}

func TestExperiencePointsGainedAfterAttack(t *testing.T) {
	attacker := NewCharacter()
	defender := NewCharacter()
	dieRoll, isCritical := RollDie(15, attacker)
	expectedExperience := 10

	Attack(attacker, defender, dieRoll, isCritical)

	if attacker.Experience != expectedExperience {
		t.Errorf("Attacker experience incorrect. Expected %d, and got %d", expectedExperience, attacker.Experience)
	} else {
		t.Log("Defender hit points correctly calculated after a negative strength modifier and a roll high enough to attack")
	}
}

func TestLevelDefaultsToOne(t *testing.T) {
	attacker := NewCharacter()
	expectedLevel := 1
	if attacker.Level != expectedLevel {
		t.Errorf("Character level was %d, test expected %d", attacker.Level, expectedLevel)
	} else {
		t.Log("Character level defaults correctly")
	}
}

func TestLevelUpWithEnoughExperience(t *testing.T) {
	attacker := NewCharacter()
	attacker.Experience = 1000
	expectedLevel := 2
	calculatedLevel := attacker.CalculatedLevel()

	if calculatedLevel != expectedLevel {
		t.Errorf("Character level was %d, test expected %d", attacker.Level, expectedLevel)
	} else {
		t.Log("Character level increments correctly")
	}
}

func TestLevelUpWithEnoughExperienceHigherExperience(t *testing.T) {
	attacker := NewCharacter()
	attacker.Experience = 5000
	expectedLevel := 6
	calculatedLevel := attacker.CalculatedLevel()

	if calculatedLevel != expectedLevel {
		t.Errorf("Character level was %d, test expected %d", attacker.Level, expectedLevel)
	} else {
		t.Log("Character level increments correctly")
	}
}

func TestHitPointsAfterLevelUpWithConstitutionBump(t *testing.T) {
	attacker := NewCharacter()
	attacker.Abilities.Constitution = 15
	attacker.Experience = 990
	defender := NewCharacter()
	dieRoll, isCritical := RollDie(15, attacker)
	expectedHitPoints := 14

	Attack(attacker, defender, dieRoll, isCritical)
	calculatedHitPoints := attacker.CalculatedHitPoints()

	if calculatedHitPoints != expectedHitPoints {
		t.Errorf("Attacker hit points were %d, test expected %d", calculatedHitPoints, expectedHitPoints)
	} else {
		t.Log("Attacker hit points incremented correctly with a level up")
	}
}

func TestAttackRollAdds1ForEveryEvenLevel(t *testing.T) {
	attacker := NewCharacter()
	attacker.Experience = 1000
	dieRoll, _ := attacker.CalculatedRoll(15)
	expectedRoll := 16

	if dieRoll != expectedRoll {
		t.Errorf("Attacker roll was %d, test expected %d", dieRoll, expectedRoll)
	} else {
		t.Log("Even levels add 1 to the die roll")
	}
}

func TestAttackRollAdds1ForEveryEvenLevelHigherLevel(t *testing.T) {
	attacker := NewCharacter()
	attacker.Experience = 5000
	dieRoll, _ := attacker.CalculatedRoll(15)
	expectedRoll := 18

	if dieRoll != expectedRoll {
		t.Errorf("Attacker roll was %d, test expected %d", dieRoll, expectedRoll)
	} else {
		t.Log("Even levels add 1 to the die roll even with higher levels")
	}
}
