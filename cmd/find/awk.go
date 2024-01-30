package find

import (
	"fmt"
	"github.com/alexperezortuno/toolbox/cmd/common"
	"github.com/spf13/cobra"
	"strings"
)

var (
	cmdStr    string
	columns   string
	delimiter string
	separator string
	fromStr   string
)

// awkCmd represents the awk command
var awkCmd = &cobra.Command{
	Use:   "text",
	Short: "awk command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var s []string
		var params []string
		var sep string
		var from string

		cmdStr = strings.TrimSpace(cmdStr)
		if cmdStr != "" {
			s = append(s, strings.Split(cmdStr, " ")...)
		}

		s = append(s, "|")
		s = append(s, "awk")
		if delimiter != "" {
			s = append(s, fmt.Sprintf("-F \"%s\"", delimiter))
		}

		if columns != "" {
			col := strings.Split(columns, ",")

			for i := 0; i < len(col); i++ {
				params = append(params, fmt.Sprintf("$%s", fmt.Sprintf("%s", col[i])))
			}
		}

		if separator == "" {
			sep = "\" \""
		} else {
			sep = fmt.Sprintf("\"%s\"", separator)
		}

		if fromStr != "" {
			from = fmt.Sprintf("NR>%s", fromStr)
		}

		s = append(s, fmt.Sprintf("'%s{print %s}'", from, strings.Join(params, sep)))
		fmt.Printf(common.LineStr, strings.Join(s, " "))

		out, err := common.LaunchCommand([]string{"bash", "-c", strings.Join(s, " ")})
		if err != nil {
			fmt.Printf(common.CmdFailed, err)
		}
		fmt.Printf(common.LineStr, out)
	},
}

func init() {
	awkCmd.Flags().StringVarP(&cmdStr, "command", "c", "", "Command to execute")
	awkCmd.Flags().StringVarP(&columns, "columns", "n", "", "Columns")
	awkCmd.Flags().StringVarP(&delimiter, "delimiter", "d", "", "Delimiter")
	awkCmd.Flags().StringVarP(&separator, "separator", "s", "", "Separator")
	awkCmd.Flags().StringVarP(&fromStr, "from", "f", "", "From row")
}
