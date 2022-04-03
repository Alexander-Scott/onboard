package lib

import (
	"context"
	"fmt"

	"github.com/google/go-github/v43/github"
	"golang.org/x/oauth2"
)

func write_plan(plan_data map[interface{}]interface{}) {
	destination_domain := prompt_destination_domain()
	destination_org := prompt_destination_org()
	destination_repo := prompt_destination_repo()
	fmt.Printf("Destination repo is %q/%q/%q\n", destination_domain, destination_org, destination_repo)

	github_client := create_client()

	repository := get_repository(github_client, destination_domain, destination_org, destination_repo)
	if repository != nil {
		issue := &github.IssueRequest{
			Title: github.String("test"),
		}
		github_client.Issues.Create(context.Background(), destination_org, destination_repo, issue)
	}
}

func create_client() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(ctx, ts)

	// `NewEnterpriseClient` for GHE
	client := github.NewClient(tc)
	return client
}

func get_repository(client *github.Client, domain string, org string, repo string) *github.Repository {
	repository, _, err := client.Repositories.Get(context.Background(), org, repo)
	if err != nil {
		fmt.Printf("Request failed %v\n", err)
		return nil
	}
	fmt.Print("Repo access check success!")
	return repository
}
