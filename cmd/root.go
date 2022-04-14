/*
* @Author: wangqilong
* @Description:
* @File: root
* @Date: 2021/9/18 8:04 下午
 */

package cmd

import (
	"github.com/spf13/cobra"
	"ippool_center/app"
	"os"
)

var RootCmd = &cobra.Command{
	Use:     "ip_pool",
	Short:   "ip池中心",
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
