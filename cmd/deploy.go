package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your agent on Kubernetes",
	Long: `Run your agent on Kubernetes.

Examples:
    hive-cli deploy -f agent.py`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return deployAgent("")
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deployAgent(source string) error {
	fmt.Println("Coming soon: Deploying an agent to Kubernetes...")
	return nil
}
