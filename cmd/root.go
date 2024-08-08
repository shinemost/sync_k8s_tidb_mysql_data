/*
Copyright © 2024 shinemost
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync_k8s_tidb_mysql_data/entity"
	"sync_k8s_tidb_mysql_data/service"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Config entity.Config
var clearTableName string

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
			args = append(args, "proudce", "produce_in", "produce_param")
		} else {
			log.Printf("开始清理【%s】，请稍后！", strings.Join(args, ","))
		}
		err := service.Clear(args)
		if err != nil {
			log.Fatalf("清理报错！%v", err)
			return
		}
	},
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "导入数据",
	Long:  `导入数据进数据库`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("开始导入所有表，请稍后！")
			args = append(args, "proudce", "produce_in", "produce_param")
		} else {
			log.Printf("开始导入【%s】，请稍后！", strings.Join(args, ","))
		}
		err := service.Insert(args)
		if err != nil {
			log.Fatalf("插入报错！{}", err)
			return
		}
	},
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "梭哈",
	Long:  `一把梭哈，扔个核弹`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("核弹准备中，扔……")
		err := service.All()
		if err != nil {
			log.Fatalf("核弹威力太大，黑猪跑不掉，烧成黑炭了！{}", err)
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

	// 将配置文件映射到结构体
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

}
