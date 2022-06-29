package cmd

import (
	"github.com/codexlynx/brutemq/pkg/bruteforcer"
	"github.com/codexlynx/brutemq/pkg/etcd"
	"github.com/spf13/cobra"
	"log"
)

var (
	etcdEndpoint string
	etcdUser     string
)

var etcdCmd = &cobra.Command{
	Use:     "etcd",
	Short:   "Bruteforce etcdv3 service endpoint",
	Aliases: []string{"e", "et", "etc"},
	Run: func(cmd *cobra.Command, args []string) {
		bruteEtcd, err := etcd.NewBruteEtcd(etcdEndpoint, etcdUser)
		if err != nil {
			log.Println(err)
		}
		defer bruteEtcd.Close()

		brute := bruteforcer.NewBruterforcerFile(bruteEtcd.TryPassword, threads, dictionary)
		log.Println("etcdv3 endpoint:", etcdEndpoint)
		log.Println("etcdv3 user:", etcdUser)

		brute.Start()
	},
}

func init() {
	etcdCmd.PersistentFlags().StringVarP(&etcdEndpoint, "endpoint", "e", "127.0.0.1:2379", "etcdv3 endpoint")
	etcdCmd.PersistentFlags().StringVarP(&etcdUser, "user", "u", "root", "etcdv3 username")
	rootCmd.AddCommand(etcdCmd)
}
