package weapon_type

import (
	"runesynergy.dev/enum"
	"runesynergy.dev/enum/animation"
	. "runesynergy.dev/enum/combat"
)

var Sword2H = enum.NewWeaponType(&enum.WeaponType{
	Ref: "sword_2h",
	Animations: enum.WeaponAnimations{
		ActionSlash:  animation.Slash2H,
		ActionCrush:  animation.Crush2H,
		ActionDefend: animation.Block2H,
	},
	Sounds: enum.WeaponSounds{
		ActionSlash: "2h_slash",
		ActionCrush: "2h_crush",
		ActionEquip: "equip_sword",
	},
	Options: []Option{
		{
			Name:   "Chop",
			Style:  StyleAccurate,
			Action: ActionSlash,
		},
		{
			Name:   "Slash",
			Style:  StyleAggressive,
			Action: ActionSlash,
		},
		{
			Name:   "Smash",
			Style:  StyleAggressive,
			Action: ActionCrush,
		},
		{
			Name:   "Block",
			Style:  StyleDefensive,
			Action: ActionSlash,
		},
	},
})
