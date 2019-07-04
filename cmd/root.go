package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Version    string
	cubeConfig string
)

func init() {

	RootCmd.PersistentFlags().StringVar(&cubeConfig, "cube-config", "", "config file (default is $HOME/.kubeconfig)")

}

var RootCmd = &cobra.Command{
	Use:   "some",
	Short: "Some tests",
	Long:  `Some tests`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
