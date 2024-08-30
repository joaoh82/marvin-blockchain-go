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

// addressCreateCmd represents the create command
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
		mnemonic, err := crypto.GetMnemonicFromEntropy(entropy)
		if err != nil {
			fmt.Println("Error generating mnemonic:", err)
			return
		}

		privateKey, err := crypto.NewPrivateKeyfromMnemonic(mnemonic)
		if err != nil {
			fmt.Println("Error generating private key:", err)
			return
		}
		publicKey := privateKey.PublicKey()
		address := publicKey.Address()
		fmt.Println("mnemonic:", mnemonic)
		fmt.Println("address:", address)
	},
}

// mnemonicAddressRestoreCmd represents the restore command
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
		privateKey, err := crypto.NewPrivateKeyfromMnemonic(mnemonic.Value.String())
		if err != nil {
			fmt.Println("Error generating private key:", err)
			return
		}
		publicKey := privateKey.PublicKey()
		address := publicKey.Address()
		fmt.Println("address:", address)
	},
}

func init() {
	mnemonicAddressRestoreCmd.Flags().String("mnemonic", "", "The mnemonic to restore the address from")
}
