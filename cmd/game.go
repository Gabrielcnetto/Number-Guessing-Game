/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Gabrielcnetto/Number-Guessing-Game/services"
	"github.com/spf13/cobra"
)

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Guess a number",
	Long:  `Set a difficulty and play an lucky game`,
	RunE: func(cmd *cobra.Command, args []string) error {
		services.Game()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(gameCmd)

}
