package soldier_strategy

import "fmt"

type ShieldSoldier struct {
}

func (s ShieldSoldier) Info() {
	fmt.Println("I'm a Soldier with Shield")
}
