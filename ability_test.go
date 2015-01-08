package main

import (
	"testing"
)

func TestAbilityDefaults(t *testing.T) {
	abilities := NewAbilities()
	valuesCorrect := true

	if abilities.Strength != 10 {
		valuesCorrect = false
	}

	if abilities.Dexterity != 10 {
		valuesCorrect = false
	}

	if abilities.Constitution != 10 {
		valuesCorrect = false
	}

	if abilities.Wisdom != 10 {
		valuesCorrect = false
	}

	if abilities.Intelligence != 10 {
		valuesCorrect = false
	}

	if abilities.Charisma != 10 {
		valuesCorrect = false
	}

	if valuesCorrect {
		t.Log("Default ability values are correct")
	} else {
		t.Error("Default ability values are not correct")
	}
}

func TestGetModifiers(t *testing.T) {
	valuesCorrect := true

	if GetAbilityModifier(1) != -5 {
		valuesCorrect = false
	}

	if GetAbilityModifier(5) != -3 {
		valuesCorrect = false
	}

	if GetAbilityModifier(10) != 0 {
		valuesCorrect = false
	}

	if GetAbilityModifier(15) != 2 {
		valuesCorrect = false
	}

	if GetAbilityModifier(20) != 5 {
		valuesCorrect = false
	}

	if valuesCorrect {
		t.Log("Ability modifier values are correct")
	} else {
		t.Error("Ability modifier values are not correct")
	}
}
