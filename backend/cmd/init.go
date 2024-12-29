package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Automatic initialization of cobra.
// This function setup initial function for CMDs,
// and set some default flags.
func init() {
	cobra.OnInitialize(initConfig)

	// set config file flag
	// 默认配置文件为config.yaml，在/backend目录下
	rootCmd.PersistentFlags().StringP("config", "c", "config.yaml", "config file (default is ./config.yaml)")
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

// initConfig reads in config file and ENV variables if set.
// This function is called automatically by cobra.OnInitialize() before rootCmd.Execute()
func initConfig() {
	// set config file
	viper.SetConfigFile(viper.GetString("config"))
	fmt.Println("config file:", viper.GetString("config"))

	// read config file
	err := viper.ReadInConfig()
	if err != nil {
		//fmt.Println(err)
		fmt.Println("config file not found")
	}
}
