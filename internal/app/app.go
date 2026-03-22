package app

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
)

// Run executes the main application logic
func Run() {
	log.Println("Executing application logic inside internal/app...")
	fmt.Println("Hello from your new project!")
}

func Help(cmd *cobra.Command, args []string) {
	// ctx := cmd.Context()
	cmd.Help()

}