# GitHub User Activity

## Overview

A command-line tool written in Go that fetches and displays a GitHub user's recent public activity directly in the terminal. Built as a practical exercise in working with REST APIs, JSON parsing, and CLI development in Go.

## What It Does

Given a GitHub username, the tool queries the GitHub Events API and prints a human-readable summary of the user's latest actions — such as commits pushed, issues opened, repositories starred or forked, and pull requests submitted.

## Technical Highlights

- **Language:** Go
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra) — the same library used by tools like kubectl and Hugo
- **Architecture:** Clean separation between the CLI layer (`cmd/`) and the business logic (`internal/github/`) following Go project conventions
- **API Integration:** Consumes the public GitHub REST API (`/users/{username}/events`) with proper HTTP status handling
- **Data Grouping:** Events are grouped by type and repository before formatting, so repeated actions (e.g. multiple commits to the same repo) are aggregated into a single summary line

## Project Structure

```
.
├── main.go                          # Entry point
├── cmd/
│   ├── root.go                      # Root Cobra command
│   └── githubActivity.go            # github-activity subcommand
└── internal/
    └── github/
        ├── UserActivity.go          # API fetch + output formatting
        └── domain/
            ├── GithubActivities.go  # Activity model
            └── User.go              # User model
```

## Skills Demonstrated

- Building structured CLI applications in Go with Cobra
- Consuming and deserializing REST API responses
- Organizing Go code with the `internal/` package pattern
- Handling HTTP errors and edge cases (user not found, empty activity)

## Inspiration

This project is based on the [GitHub User Activity](https://roadmap.sh/projects/github-user-activity) challenge from roadmap.sh — a practical project for developers learning to work with external APIs from the command line.
