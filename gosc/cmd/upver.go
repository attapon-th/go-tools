package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/bitfield/script"
	"github.com/spf13/cobra"
)

var (
	regexPatten = `(\b\d+\.\d+\.)(\d{0,2})\b`
)

// upverCmd represents the upver command
var upverCmd = &cobra.Command{
	Use:   "upver [filename]",
	Short: "Up version in file ",
	Long: `Increment Version from version number(ex. 1.2.3 --> 1.2.4) 
Example:
	VERSION: 0.1.9
	--- result --
	VERSION: 0.1.10`,
	Args: cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		re, err := regexp.Compile(regexPatten)
		checkErrAndExist(err)
		fs, err := os.OpenFile(args[0], os.O_RDONLY, os.ModePerm)
		defer fs.Close()
		checkErrAndExist(err)
		bs := bufio.NewScanner(fs)

		fileText := ""
		isSet := false
		for bs.Scan() {
			l := bs.Text()
			if !isSet && re.MatchString(l) {
				ss := re.FindStringSubmatch(l)
				if len(ss) == 3 {
					n, _ := strconv.Atoi(ss[2])
					newVer := fmt.Sprintf("%s%d", ss[1], n+1)

					l = re.ReplaceAllString(l, newVer)
					script.Echo(fmt.Sprintf("Saved file: %s\nOld Version: %s\nNew Version: %s\n", args[0], ss[0], newVer)).Stdout()
					isSet = true
				}
			}
			fileText += fmt.Sprintln(l)
		}
		checkErrAndExist(bs.Err())

		fw, err := os.OpenFile(args[0], os.O_WRONLY, os.ModePerm)
		checkErrAndExist(err)
		defer fw.Close()
		fw.Write([]byte(fileText))

	},
}

func init() {
	rootCmd.AddCommand(upverCmd)

	upverCmd.PersistentFlags().StringVarP(&regexPatten, "regex", "r", regexPatten, "Set reject search version")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
