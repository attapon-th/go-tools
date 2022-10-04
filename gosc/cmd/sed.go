package cmd

import (
	"regexp"

	"github.com/bitfield/script"
	"github.com/spf13/cobra"
)

// sedCmd represents the sed command
var sedCmd = &cobra.Command{
	Use:   "sed [search-regex] [replace-string]",
	Short: "Replace string from STDIN with MatchRegexp",
	Long:  `Replace string from STDIN with MatchRegexp`,
	Args:  cobra.ExactValidArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		re := regexp.MustCompile(args[0])
		script.Stdin().ReplaceRegexp(re, args[1]).Stdout()
	},
}

func init() {
	rootCmd.AddCommand(sedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
