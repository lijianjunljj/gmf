/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gmf/src/common/config"
	"gmf/src/common/config/parser"
	"gmf/src/common/server"
	"gmf/src/common/utils"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		fmt.Println(args)
		if len(args) > 0 {
			serverName := args[0]
			configFile := args[1]
			if !utils.FileIsExisted(configFile) {
				fmt.Println("配置文件不存在")
				return
			}
			yamlParser := parser.NewYamlParser(configFile)
			yamlParser.Parse()
			config := config.NewConfig(yamlParser)
			config.InitMysql()
			config.InitEtcd()
			config.InitWeb()
			if serverName == "all" {
				server.StartAll(config)
			} else {
				server.Start(serverName, config)
			}
		} else {
			fmt.Println("please input server name")
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	startCmd.Flags().BoolP("name", "n", false, "input server name")
	startCmd.Flags().BoolP("config", "c", false, "input config file")
}
