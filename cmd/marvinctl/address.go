package main

import (
	"fmt"

	"github.com/joaoh82/marvinblockchain/crypto"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip39"
)

// addressCmd represents the address command
var addressCmd = &cobra.Command{
	Use:                   "address",
	Short:                 "Manage addresses",
	Long:                  `Manage addresses for the Marvin Blockchain`,
	DisableFlagsInUseLine: true,
	Example:               "Usage: marvinctl address [command] [flags] [args]",
}

var addressCreateCmd = &cobra.Command{
	Use:   "create",
	Short: `Create a new address`,
	Long:  `Create a new address`,
	Run: func(cmd *cobra.Command, args []string) {
		entropy, err := bip39.NewEntropy(128)
		if err != nil {
			fmt.Println("Error generating entropy:", err)
			return
		}
		mnemonic := crypto.GetMnemonicFromEntropy(entropy)

		privateKey := crypto.NewPrivateKeyfromMnemonic(mnemonic)
		publicKey := privateKey.PublicKey()
		address := publicKey.Address()
		fmt.Println("mnemonic:", mnemonic)
		fmt.Println("address:", address)
	},
}

var mnemonicAddressRestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore an address from a mnemonic",
	Long:  `Restore an address from a mnemonic`,
	// Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mnemonic := cmd.Flag("mnemonic")
		if mnemonic == nil {
			fmt.Println("mnemonic flag is required")
			return
		}

		if !bip39.IsMnemonicValid(mnemonic.Value.String()) {
			fmt.Println("Invalid mnemonic")
			return
		}

		// mnemonic := args[0]
		privateKey := crypto.NewPrivateKeyfromMnemonic(mnemonic.Value.String())
		publicKey := privateKey.PublicKey()
		address := publicKey.Address()
		fmt.Println("address:", address)
	},
}

func init() {
	mnemonicAddressRestoreCmd.Flags().String("mnemonic", "", "The mnemonic to restore the address from")
}
