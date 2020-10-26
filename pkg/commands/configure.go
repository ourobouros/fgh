package commands

import (
	"fmt"

	"github.com/Matt-Gleich/fgh/pkg/commands/configure"
	"github.com/enescakir/emoji"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	Use:                   "configure",
	Short:                 fmt.Sprintf("%v  Configure fgh with an interactive prompt", emoji.Gear),
	Long:                  longDocStart + "https://github.com/Matt-Gleich/fgh#%EF%B8%8F-fgh-configure",
	Run: func(cmd *cobra.Command, args []string) {
		regularConfig := configure.AskQuestions()
		configure.WriteConfiguration(regularConfig)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
