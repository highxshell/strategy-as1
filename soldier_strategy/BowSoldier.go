package soldier_strategy

import "fmt"

type BowSoldier struct {
}

func (b BowSoldier) Info() {
	fmt.Println("I'm a Soldier with Bow")
}
