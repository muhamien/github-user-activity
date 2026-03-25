package github

import (
	"encoding/json"
	"fmt"
	"github-user-activity/internal/github/domain"
	"net/http"
)

type groupedByTypeAndRepo struct {
	Type string
	Repo string
}

func GetUserActivity(username string) ([]domain.GithubActivity, error) {
	resp, err := http.Get("https://api.github.com/users/" + username + "/events")
	if err != nil {
		println("Error")
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("username not found")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %d", resp.StatusCode)
	}

	var activities []domain.GithubActivity
	if err := json.NewDecoder(resp.Body).Decode(&activities); err != nil {
		return nil, err
	}

	return activities, nil

}

func FormatActivity(events []domain.GithubActivity) error {
	if len(events) == 0 {
		return fmt.Errorf("no activity found")
	}

	grouped := make(map[groupedByTypeAndRepo][]domain.GithubActivity)

	for _, event := range events {
		key := groupedByTypeAndRepo{
			Type: event.Type,
			Repo: event.Repo.Name,
		}
		grouped[key] = append(grouped[key], event)
	}

	for _, event := range events {
		var action string
		switch event.Type {
		case "PushEvent":
			commitCount := len(grouped[groupedByTypeAndRepo{Type: "PushEvent", Repo: event.Repo.Name}])
			if commitCount == 1 {
				action = fmt.Sprintf("Pushed a commit to %s", event.Repo.Name)
			} else {
				action = fmt.Sprintf("Pushed %d commits to %s", commitCount, event.Repo.Name)
			}
		case "IssuesEvent":
			action = fmt.Sprintf("Opened a new issue in %s", event.Repo.Name)
		case "WatchEvent":
			action = fmt.Sprintf("Starred %s", event.Repo.Name)
		case "ForkEvent":
			action = fmt.Sprintf("Forked %s", event.Repo.Name)
		case "CreateEvent":
			action = fmt.Sprintf("Created %s in %s", event.Payload.Ref, event.Repo.Name)
		case "PullRequestEvent":
			action = fmt.Sprintf("Pull Request %s in %s", event.Payload.Ref, event.Repo.Name)
		case "IssueCommentEvent":
			action = fmt.Sprintf("Add Comment %s on %s", event.Payload.Ref, event.Repo.Name)
		default:
			action = fmt.Sprintf("%s in %s", event.Type, event.Repo.Name)
		}
		fmt.Println("- " + action)
	}
	return nil
}
