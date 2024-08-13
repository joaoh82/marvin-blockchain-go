package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "marvinctl",
	Short:   "marvinclt is a cli interface for the Marvin Blockchain",
	Long:    `marvinclt is a cli interface for the Marvin Blockchain`,
	Example: "Usage: marvinctl [command] [flags] [args]",
	Version: "0.0.1",
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

// init is called before the command is executed
func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(addressCmd)
	rootCmd.SetVersionTemplate("marvinclt v0.0.1 -- HEAD")
	addressCmd.AddCommand(addressCreateCmd)
	addressCmd.AddCommand(mnemonicAddressRestoreCmd)
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
