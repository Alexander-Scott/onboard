package lib

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func prompt_plan_location() string {
	prompt := promptui.Prompt{
		Label:   "Where is the plan located? (URL or local path)",
		Default: "/Users/q483830/Documents/GitHub/onboard/data/plan_1.yaml", // TODO: Remove default
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return result
}

func prompt_destination_domain() string {
	prompt := promptui.Prompt{
		Label:   "Where is the repo located?",
		Default: "github.com",
	}

	domain, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return domain
}

func prompt_destination_org() string {
	prompt := promptui.Prompt{
		Label:   "Which organisation/user is the repo in? (must have write access)",
		Default: "alexander-scott", // TODO: Remove default value
	}

	organisation, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return organisation
}

func prompt_destination_repo() string {
	prompt := promptui.Prompt{
		Label:   "What is the name of the repository? (must have write access)",
		Default: "test_plans", // TODO: Remove default value
	}

	repo, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return repo
}
