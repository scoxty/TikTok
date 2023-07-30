package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sunflower10086/TikTok/http/version"
)

var (
	vers bool
)

var RootCmd = &cobra.Command{
	Use:     "tiktok",
	Long:    "tiktok API后端",
	Short:   "tiktok API后端",
	Example: "tiktok API后端 commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return errors.New("no flags find")
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "domo version")
}
