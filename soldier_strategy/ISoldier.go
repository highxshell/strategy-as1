package soldier_strategy

type ISoldier interface {
	Info()
}

type SoldierBehavior struct {
	SB ISoldier
}

func (soldier SoldierBehavior) DisplayInfo() {
	soldier.SB.Info()
}
