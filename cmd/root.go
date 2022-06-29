package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "brutemq",
	Short: "brutemq - An exotic service bruteforce tool",
}

var (
	dictionary string
	threads    int
)

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&dictionary, "dictionary", "d", "", "dictionary file path")
	rootCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 100, "threads number")

	err := rootCmd.MarkPersistentFlagRequired("dictionary")
	if err != nil {
		log.Println(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
