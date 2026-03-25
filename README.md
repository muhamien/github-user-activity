# GitHub User Activity

A CLI tool to fetch and display recent GitHub activity for any user, built with Go.

[Challenge roadmap.sh](https://roadmap.sh/projects/github-user-activity)

## Features

- Fetch public GitHub events for any user via the GitHub API
- Groups and formats activity by event type and repository
- Supports common event types: pushes, issues, stars, forks, pull requests, comments, and more

## Installation

```bash
git clone https://github.com/your-username/github-user-activity.git
cd github-user-activity
go build -o github-user-activity .
```

## Usage

```bash
./github-user-activity github-activity <username>
```

**Example:**

```bash
./github-user-activity github-activity muhamien
```

**Output:**

```
Output:
- Pushed 3 commits to torvalds/linux
- Starred some/repo
- Opened a new issue in torvalds/linux
```

## Event Types

| Event | Description |
|-------|-------------|
| `PushEvent` | Pushed commits to a repository |
| `IssuesEvent` | Opened a new issue |
| `WatchEvent` | Starred a repository |
| `ForkEvent` | Forked a repository |
| `CreateEvent` | Created a branch or tag |
| `PullRequestEvent` | Opened or updated a pull request |
| `IssueCommentEvent` | Commented on an issue |

## Requirements

- Go 1.21+
- Internet access to reach the GitHub API

## Dependencies

- [cobra](https://github.com/spf13/cobra) — CLI framework

## License

See [LICENSE](LICENSE).
# github-user-activity
