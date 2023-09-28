package main

import "strategy-as1/soldier_strategy"

func main() {
	BasicSoldier := soldier_strategy.BasicSoldier{}
	bowSoldier := soldier_strategy.BowSoldier{}
	shieldSoldier := soldier_strategy.ShieldSoldier{}

	soldier1 := soldier_strategy.SoldierBehavior{SB: BasicSoldier}
	soldier2 := soldier_strategy.SoldierBehavior{SB: bowSoldier}
	soldier3 := soldier_strategy.SoldierBehavior{SB: shieldSoldier}

	soldier1.DisplayInfo()
	soldier2.DisplayInfo()
	soldier3.DisplayInfo()
}
