package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goodvs/internal/controller"
	"goodvs/internal/dao"
	"goodvs/internal/service"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GoodVS",
	Short: "GoodVS price comparison",
	Long:  `GoodVS price comparison`,

	RunE: func(cmd *cobra.Command, args []string) error {
		// init db
		dao.InitDB()

		//// start server
		fmt.Println("Hello, GoodVS!")
		//service.SendEmail("3220100886@zju.edu.cn")
		service.InitES()
		return controller.StartServer()

		//return nil
	},
}

// Execute is the entry point of the program
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
