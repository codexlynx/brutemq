package cmd

import (
	"fmt"
	"github.com/codexlynx/brutemq/pkg/amqp"
	"github.com/codexlynx/brutemq/pkg/bruteforcer"
	"github.com/spf13/cobra"
	"log"
)

var (
	amqpEndpoint string
	amqpUser     string
)

var amqpCmd = &cobra.Command{
	Use:     "amqp",
	Short:   "Bruteforce AMQP PLAIN SASL service endpoint",
	Aliases: []string{"a", "am", "amq"},
	Run: func(cmd *cobra.Command, args []string) {
		bruteAmqp := amqp.BruteAmqpPlain{
			Endpoint: amqpEndpoint,
			User:     amqpUser,
		}

		metadata := fmt.Sprintf("amqp://%s@%s", amqpUser, amqpEndpoint)
		brute := bruteforcer.NewBruterforcerFile(bruteAmqp.TryPassword, threads, metadata, dictionary)
		log.Println("AMQP PLAIN SASL endpoint:", amqpEndpoint)
		log.Println("AMQP PLAIN SASL user:", amqpUser)

		brute.Start()
	},
}

func init() {
	amqpCmd.PersistentFlags().StringVarP(&amqpEndpoint, "endpoint", "e", "localhost:5672/vhost", "AMQP PLAIN SASL endpoint")
	amqpCmd.PersistentFlags().StringVarP(&amqpUser, "user", "u", "admin", "AMQP PLAIN SASL username")
	rootCmd.AddCommand(amqpCmd)
}
