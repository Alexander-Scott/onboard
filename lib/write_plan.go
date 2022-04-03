package lib

import (
	"context"
	"fmt"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

func check_repo_access(domain string, org string, repo string) bool {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(ctx, ts)

	// `NewEnterpriseClient` for GHE
	client := github.NewClient(tc)
	_, _, err := client.Repositories.Get(ctx, org, repo)
	if err != nil {
		fmt.Printf("Request failed %v\n", err)
		return false
	}
	fmt.Print("Repo access check success!")
	return true
}
