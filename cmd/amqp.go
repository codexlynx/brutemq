package cmd

import (
	"github.com/codexlynx/brutemq/pkg/amqp"
	"github.com/codexlynx/brutemq/pkg/bruteforcer"
	"github.com/spf13/cobra"
	"log"
)

var endpoint string

var amqpCmd = &cobra.Command{
	Use:     "amqp",
	Short:   "Bruteforce AMQP service",
	Aliases: []string{"a", "am", "amq"},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting brutemq", version, "...")
		log.Println("AMQP endpoint:", endpoint)
		log.Println("AMQP user:", user)

		bruteAmqp := amqp.BruteAmqp{
			Endpoint: endpoint,
			User:     user,
		}

		log.Println("Attacking...")
		bruteforcer.StartBruteforcerWithFile(bruteAmqp.TryPassword, threads, dictionary)
	},
}

func init() {
	amqpCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "localhost:5672/vhost", "AMQP service endpoint")
	rootCmd.AddCommand(amqpCmd)
}
