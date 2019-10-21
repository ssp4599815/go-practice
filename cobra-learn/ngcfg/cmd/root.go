package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "ngcfg",
	Short: " Get Nginx Config with Filebeat",
	Long:  " 通过 ansible 获取 nginx 配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ok")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
