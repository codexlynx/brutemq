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
	Short:   "Bruteforce AMQP Plain service endpoint",
	Aliases: []string{"a", "am", "amq"},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("AMQP Plain endpoint:", endpoint)
		log.Println("AMQP Plain user:", user)

		bruteAmqp := amqp.BruteAmqpPlain{
			Endpoint: endpoint,
			User:     user,
		}
		brute := bruteforcer.NewBruterforcerFile(bruteAmqp.TryPassword, threads, dictionary)
		brute.Start()
	},
}

func init() {
	amqpCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "localhost:5672/vhost", "AMQP Plain endpoint")
	rootCmd.AddCommand(amqpCmd)
}
