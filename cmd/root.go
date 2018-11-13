package cmd

import (
	"fmt"
	"github.com/codexlynx/brutemq/lib"
	"github.com/spf13/cobra"
	"os"
)

var (
	url       *string
	user      *string
	passwords *string
	threads   *int
)

var RootCmd = &cobra.Command{
	Use:   "brutemq",
	Short: "High performance RabbitMQ (amqp) Brute Force tool.",
	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Help()
		lib.Start(*url, *user, *passwords, *threads)
	},
}

func init() {
	url = RootCmd.PersistentFlags().String("url", "localhost:5672/vhost", "rabbitmq connection URL")
	user = RootCmd.PersistentFlags().StringP("user", "u", "guest", "username")
	passwords = RootCmd.PersistentFlags().StringP("file", "f", "passwords.txt", "load several passwords from file")
	threads = RootCmd.PersistentFlags().IntP("threads", "t", 200, "number of threads")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
