/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"github-user-activity/internal/github"

	"github.com/spf13/cobra"
)

// githubActivityCmd represents the githubActivity command
var githubActivityCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		err := RunGithubActivity(args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	},
}

func RunGithubActivity(args []string) error {
	if len(args) < 1 {
		return errors.New("please provide github username")
	}
	username := args[0]
	activities, err := github.GetUserActivity(username)
	if err != nil {
		return err
	}
	fmt.Println("Output:")
	return github.FormatActivity(activities)
}

func init() {
	rootCmd.AddCommand(githubActivityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// githubActivityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// githubActivityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
