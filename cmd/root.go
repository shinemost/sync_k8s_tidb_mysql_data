/*
Copyright © 2024 shinemost
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"sync_k8s_tidb_mysql_data/service"
)

var allTables = []string{"produce", "produce_in", "produce_param"}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sync",
	Short: "测试环境TIDB数据同步工具",
	Long:  "将测试环境TIDB数据库相关表数据清除，并导入存放的历史数据",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var clearCmd = &cobra.Command{
	Use:   "clear table1 table2 …… [optional] empty args says all",
	Short: "清理数据",
	Long:  `清理数据库数据`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("开始清理所有表，请稍后！")
			args = allTables
		} else {
			log.Printf("开始清理【%s】，请稍后！", strings.Join(args, ","))
		}
		err := service.Clear(args)
		if err != nil {
			log.Fatalf("清理报错！%v", err)
			return
		}
		log.Printf("【%s】清理成功！", strings.Join(args, ","))
	},
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "导入数据",
	Long:  `导入数据进数据库`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("开始导入所有表，请稍后！")
			args = allTables
		} else {
			log.Printf("开始导入【%s】，请稍后！", strings.Join(args, ","))
		}
		err := service.Insert(args)
		if err != nil {
			log.Fatalf("导入报错！%v", err)
			return
		}
		log.Printf("【%s】导入成功！", strings.Join(args, ","))
	},
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "梭哈",
	Long:  `先清数据再导入`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("开始处理【%s】，请稍后！", strings.Join(allTables, ","))
		err := service.All(allTables)
		if err != nil {
			log.Fatalf("处理报错！%v", err)
			return
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(clearCmd, importCmd, allCmd)
	cobra.OnInitialize(initConfig)
}

// initConfig reads in util file and ENV variables if set.
func initConfig() {

	// Find home directory.
	//home, err := os.UserHomeDir()
	//cobra.CheckErr(err)

	// Search util in home directory with name ".sync_k8s_tidb_mysql_data" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a util file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using util file:", viper.ConfigFileUsed())
	}

}
