package cmd

import (
	"fmt"
	"github.com/Rankgice/api-gen/cmd/api"
	"github.com/Rankgice/api-gen/cmd/rpc"
	"github.com/spf13/cobra"
	"os"
)

// 定义根命令
var rootCmd = &cobra.Command{
	Use:   "api-gen",
	Short: "A tool for generating Go code",
}

// 定义 api 子命令
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Generate API definition from Mysql",
	Run:   api.RunApi,
}

// 定义 rpc 子命令
var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "Generate RPC definition from Mysql",
	Run:   rpc.RunRpc,
}

func init() {
	// 添加子命令到根命令
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(rpcCmd)

	// 为 api 子命令定义参数
	apiCmd.Flags().StringP("url", "u", "", "The data source of database,like \"root:password@tcp(127.0.0.1:3306)/database\" (required)")
	apiCmd.Flags().StringP("table", "t", "", "The name of the database table (required)")
	apiCmd.Flags().StringP("dir", "d", "./desc", "The target dir")
	apiCmd.Flags().StringP("prefix", "p", "/prefix", "This is the API file routing prefix")
	apiCmd.Flags().StringP("home", "", "", "The apiGen home path of the template")
	apiCmd.Flags().StringP("service", "s", "", "This is the service name of the API file")
	_ = apiCmd.MarkFlagRequired("url")
	_ = apiCmd.MarkFlagRequired("table")

	// 为 rpc 子命令定义参数
	rpcCmd.Flags().StringP("url", "u", "", "The data source of database,like \"root:password@tcp(127.0.0.1:3306)/database\" (required)")
	rpcCmd.Flags().StringP("table", "t", "", "The name of the database table (required)")
	rpcCmd.Flags().StringP("dir", "d", "./pb", "The target dir")
	rpcCmd.Flags().StringP("package", "p", "", "This is the name of the generated PB file package")
	rpcCmd.Flags().StringP("home", "", "", "The apiGen home path of the template")
	_ = rpcCmd.MarkFlagRequired("url")
	_ = rpcCmd.MarkFlagRequired("table")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
