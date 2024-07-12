// Package commands содержит логику обработки команд пользователя.
package commands

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/al-kirpichenko/gophkeeper/internal/client/commands/auth"
	"github.com/al-kirpichenko/gophkeeper/internal/client/commands/keeper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "welcome to client for a gophkeeper server",
}

func init() {
	rootCmd.AddCommand(auth.Cmd, keeper.Cmd)
	rootCmd.PersistentFlags().StringP("server", "a", "http://localhost:8080", "host addr")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
