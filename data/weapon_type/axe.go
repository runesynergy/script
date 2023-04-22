package weapon_type

import (
	"fmt"

	"common.runesynergy.dev/enum"
	"common.runesynergy.dev/enum/combat"
)

func init() {
	fmt.Println("e-z")
}

var Axe = enum.NewWeaponType(&enum.WeaponType{
	Ref: "axe",

	Animations: enum.WeaponAnimations{
		combat.ActionAttack: nil,
	},
})
