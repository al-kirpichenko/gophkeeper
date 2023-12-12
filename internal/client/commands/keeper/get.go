package keeper

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get command",
	Long:  `The get command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		title := cmd.Flag("title").Value.String()
		secret, err := keeperService.Get(title)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(secret)
		return nil
	},
}

func init() {
	Cmd.AddCommand(getCmd)
}
