package main

type Abilities struct {
	Strength     int
	Dexterity    int
	Constitution int
	Wisdom       int
	Intelligence int
	Charisma     int
}

func NewAbilities() Abilities {
	abilities := Abilities{10, 10, 10, 10, 10, 10}

	return abilities
}

func GetAbilityModifier(score int) int {
	modifier := 0
	n := score - 10

	if n%2 != 0 {
		n -= 1
	}

	modifier = n / 2

	return modifier
}
