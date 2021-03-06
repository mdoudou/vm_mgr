package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoneyan/vm_mgr/controller/api"
	"github.com/yoneyan/vm_mgr/controller/server"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start controller server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		go api.VNCProxy()
		server.Server()
		fmt.Println("end")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

}
