package cmd

import (
	"fmt"
	"os"
	"pingcode-client/info"
	"pingcode-client/internal/app"

	"pingcode-client/internal/pkg/sdk"

	"github.com/spf13/cobra"
)

var (
	clientID     string
	clientSecret string
	pcClient     *sdk.Client
)

var rootCmd = &cobra.Command{
	Use:     "pingcode-client",
	Short:   "PingCode Go Client CLI",
	Long:    "A CLI tool for interacting with PingCode API using Client Credentials.",
	Version: info.Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if clientID == "" {
			clientID = os.Getenv("PINGCODE_CLIENT_ID")
		}
		if clientSecret == "" {
			clientSecret = os.Getenv("PINGCODE_CLIENT_SECRET")
		}

		if clientID != "" && clientSecret != "" {
			pcClient = sdk.NewClient(clientID, clientSecret)
		}
	},
	Run: app.Help,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&clientID, "client-id", "", "PingCode Client ID (env: PINGCODE_CLIENT_ID)")
	rootCmd.PersistentFlags().StringVar(&clientSecret, "client-secret", "", "PingCode Client Secret (env: PINGCODE_CLIENT_SECRET)")
}

func GetClient() (*sdk.Client, error) {
	if pcClient == nil {
		return nil, fmt.Errorf("client not initialized. Please provide client-id and client-secret via flags or environment variables")
	}
	return pcClient, nil
}
