package wiki

import (
	"fmt"
	"log"
	"pingcode-client/internal/pkg/sdk"

	"github.com/spf13/cobra"
)

// Run executes the logic for wiki
func Run(client *sdk.Client) {
	log.Println("Listing wiki spaces from Wiki...")
	spaces, err := client.ListWikiSpaces(0, 30)
	if err != nil {
		log.Fatalf("Error listing wiki spaces: %v", err)
	}

	fmt.Println("Wiki Spaces:")
	for _, s := range spaces {
		fmt.Printf("- %s (ID: %s)\n", s.Name, s.ID)
	}
}

func Help(cmd *cobra.Command, args []string) {
	// ctx := cmd.Context()
	cmd.Help()
}
