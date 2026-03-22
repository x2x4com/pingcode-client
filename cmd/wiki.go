package cmd

import (
	"pingcode-client/internal/app/wiki"

	"github.com/spf13/cobra"
)

var wikiCmd = &cobra.Command{
	Use:   "wiki",
	Short: "PingCode 知识管理 (Wiki) 模块",
	Long:  `用于列出和查看 PingCode Wiki 空间信息。`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := GetClient()
		if err != nil {
			panic(err)
		}
		wiki.Run(client)
	},
}

func init() {
	rootCmd.AddCommand(wikiCmd)

	// Here you will define your flags and configuration settings.
	// wikiCmd.PersistentFlags().String("foo", "", "A help for foo")
	// wikiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
