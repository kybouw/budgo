package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type BudgPlan map[string]map[string]float64

func BudgTable(plan BudgPlan, amount float64) string {
	var ret string = ""
	for groupName, groupValue := range plan {
		ret += fmt.Sprintln(groupName)
		for itemName, itemValue := range groupValue {
			var itemRatio float64 = itemValue / 100
			var itemAmount float64 = amount * itemRatio
			var lineOut string = "  %-20s$%8.2f\n"
			ret += fmt.Sprintf(lineOut, itemName, itemAmount)
		}
	}
	return ret
}

func main() {
	planInputRaw, err := os.ReadFile("plan.toml")
	if err != nil {
		panic(err)
	}
	var planInput string = string(planInputRaw[:])

	var plan BudgPlan
	toml.Decode(planInput, &plan)

	var amount float64 = 1000
	fmt.Print(BudgTable(plan, amount))
}
