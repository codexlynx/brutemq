package cmd

import (
	"github.com/codexlynx/brutemq/pkg/bruteforcer"
	"github.com/codexlynx/brutemq/pkg/vault"
	"github.com/spf13/cobra"
	"log"
)

var (
	vaultEndpoint string
	vaultUser     string
)

var vaultCmd = &cobra.Command{
	Use:     "vault",
	Short:   "Bruteforce HashiCorp Vault Userpass auth",
	Aliases: []string{"v", "va", "vau", "vaul"},
	Run: func(cmd *cobra.Command, args []string) {
		bruteVault := vault.BruteVaultUserPass{
			Endpoint: vaultEndpoint,
			User:     vaultUser,
		}

		brute := bruteforcer.NewBruterforcerFile(bruteVault.TryPassword, threads, dictionary)
		log.Println("HashiCorp Vault endpoint:", vaultEndpoint)
		log.Println("HashiCorp Vault user:", vaultUser)

		brute.Start()
	},
}

func init() {
	vaultCmd.PersistentFlags().StringVarP(&vaultEndpoint, "endpoint", "e", "http://127.0.0.1:8200", "Hashicorp Vault endpoint")
	vaultCmd.PersistentFlags().StringVarP(&vaultUser, "user", "u", "root", "HashiCorp Vault username")
	rootCmd.AddCommand(vaultCmd)
}
