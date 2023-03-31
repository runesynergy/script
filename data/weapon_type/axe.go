package weapon_type

import (
	"fmt"

	"runesynergy.dev/enum"
	"runesynergy.dev/enum/animation"
	"runesynergy.dev/enum/combat"
)

func init() {
	fmt.Println("e-z")
}

var Axe = enum.NewWeaponType(&enum.WeaponType{
	Ref: "axe",

	Animations: enum.WeaponAnimations{
		combat.ActionAttack: animation.Slash2H,
	},
})
