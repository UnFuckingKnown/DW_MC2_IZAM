package main

import (
	"fmt"
)

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func main() {

	profile := MakeProfile("naruto")
	fmt.Println("Name :", profile.Name)
	fmt.Println("Health :", profile.Health)
	fmt.Println("Power :", profile.Power)
	fmt.Println("Exp :", profile.Exp)
	powerUp := PowerUp(profile, 2)
	fmt.Println("Name :", powerUp.Name)
	fmt.Println("Health :", powerUp.Health)
	fmt.Println("Power :", powerUp.Power)
	fmt.Println("Exp :", powerUp.Exp)
}

func MakeProfile(name string) Profile {
	profile := Profile{
		Name:   name,
		Health: 100,
		Power:  200,
		Exp:    90,
	}
	return profile
}

func PowerUp(profile Profile, multi int) Profile {
	b := Profile{
		Name:   profile.Name,
		Health: profile.Health + (profile.Health * multi),
		Power:  profile.Power + (profile.Power * multi),
		Exp:    profile.Exp + (profile.Exp * multi),
	}
	return b
}
