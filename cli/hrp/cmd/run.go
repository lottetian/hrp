package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/lottetian/hrp"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run $path...",
	Short: "run API test",
	Long:  `run yaml/json testcase files for API test`,
	Example: `  $ hrp run demo.json	# run specified json testcase file
  $ hrp run demo.yaml	# run specified yaml testcase file
  $ hrp run examples/	# run testcases in specified folder`,
	Args: cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		setLogLevel(logLevel)
	},
	Run: func(cmd *cobra.Command, args []string) {
		var paths []hrp.ITestCase
		for _, arg := range args {
			paths = append(paths, &hrp.TestCasePath{Path: arg})
		}
		runner := hrp.NewRunner(nil).
			SetFailfast(!continueOnFailure).
			SetSaveTests(saveTests)
		if genHTMLReport {
			runner.GenHTMLReport()
		}
		if proxyUrl != "" {
			runner.SetProxyUrl(proxyUrl)
		}
		err := runner.Run(paths...)
		if err != nil {
			os.Exit(1)
		}
	},
}

var (
	continueOnFailure bool
	silentFlag        bool
	proxyUrl          string
	saveTests         bool
	genHTMLReport     bool
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVar(&continueOnFailure, "continue-on-failure", false, "continue running next step when failure occurs")
	runCmd.Flags().BoolVarP(&silentFlag, "silent", "s", false, "disable logging request & response details")
	runCmd.Flags().StringVarP(&proxyUrl, "proxy-url", "p", "", "set proxy url")
	runCmd.Flags().BoolVar(&saveTests, "save-tests", false, "save tests summary")
	runCmd.Flags().BoolVarP(&genHTMLReport, "gen-html-report", "r", false, "generate html report")
}
