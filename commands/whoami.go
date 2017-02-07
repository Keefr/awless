package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wallix/awless/cloud/aws"
)

func init() {
	RootCmd.AddCommand(whoamiCmd)
}

var whoamiCmd = &cobra.Command{
	Use:                "whoami",
	Aliases:            []string{"who"},
	PersistentPreRun:   applyHooks(initAwlessEnvHook, initCloudServicesHook, checkStatsHook),
	PersistentPostRunE: saveHistoryHook,
	Short:              "Show the caller identity",

	Run: func(cmd *cobra.Command, args []string) {
		resp, err := aws.SecuAPI.GetCallerIdentity(nil)
		exitOn(err)
		fmt.Println(resp)
	},
}