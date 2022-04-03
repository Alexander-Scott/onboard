package lib

import (
	"fmt"
)

func Run() error {
	// 1) Get the location of the plan
	plan_location := prompt_plan_location()
	fmt.Printf("Plan location is %q\n", plan_location)

	// 2) Fetch the plan
	plan_data := get_plan(plan_location)
	validate_plan_contents(plan_data)

	// 3) Optionally select/deselect parts of the plan

	// 4) Write the plan to the destination repo
	write_plan(plan_data)

	return nil
}
