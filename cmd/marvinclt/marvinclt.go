package main

import (
	"fmt"
	"os"

	"github.com/joaoh82/marvinblockchain/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "marvinctl",
	Short:   "marvinclt is a cli interface for the Marvin Blockchain",
	Long:    `marvinclt is a cli interface for the Marvin Blockchain`,
	Example: "Usage: marvinctl [command] [flags] [args]",
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of marvinclt",
	Long:  `All software has versions. This is marvinclt's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("marvinclt v0.0.1 -- HEAD")
	},
}

// versionCmd represents the version command
var tempCmd = &cobra.Command{
	Use:   "temp",
	Short: "Prints a temp message",
	Long:  `Prints a temp message for testing purposes`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.TempFunc()
	},
}

// init is called before the command is executed
func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(tempCmd)
}

// Execute is the entry point for the command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	Execute()
}
