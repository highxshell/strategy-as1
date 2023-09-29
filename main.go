package main

import "strategy-as1/soldier_strategy"

func main() {
	soldier := soldier_strategy.GetConfig()
	singletonSoldier := soldier_strategy.GetConfig()
	BasicSoldier := soldier
	BasicSoldier1 := singletonSoldier
	bowSoldier := soldier_strategy.BowSoldier{}
	shieldSoldier := soldier_strategy.ShieldSoldier{}

	soldier1 := soldier_strategy.SoldierBehavior{SB: BasicSoldier}
	soldier2 := soldier_strategy.SoldierBehavior{SB: bowSoldier}
	soldier3 := soldier_strategy.SoldierBehavior{SB: shieldSoldier}
	soldier4 := soldier_strategy.SoldierBehavior{SB: BasicSoldier1}

	soldier1.DisplayInfo()
	soldier2.DisplayInfo()
	soldier3.DisplayInfo()
	soldier4.DisplayInfo()
}
