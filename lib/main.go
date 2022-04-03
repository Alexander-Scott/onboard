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

	// 4) Get the destination repo URL
	destination_domain := prompt_destination_domain()
	destination_org := prompt_destination_org()
	destination_repo := prompt_destination_repo()
	fmt.Printf("Destination repo is %q/%q/%q\n", destination_domain, destination_org, destination_repo)

	// 5) Write the plan to the destination repo
	check_repo_access(destination_domain, destination_org, destination_repo)

	return nil
}
