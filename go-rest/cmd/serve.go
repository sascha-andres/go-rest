// Copyright © 2018 Sascha Andres <sascha.andres@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/sascha-andres/go-rest"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run rest api service",
	Long:  `Initializes and runs the rest api server`,
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := go_rest.NewServer(viper.GetString("serve.listen"), viper.GetInt("serve.port"))
		if err != nil {
			panic(err)
		}

		signals := make(chan os.Signal, 1)
		done := make(chan bool, 1)

		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			err = srv.ServeHTTP()
			if err != nil {
				panic(err)
			}
		}()

		go func() {
			sig := <-signals
			fmt.Println()
			fmt.Println(sig)
			done <- true
		}()
		<-done
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringP("listen", "l", "127.0.0.1", "ip binding, empty = listen on all")
	serveCmd.PersistentFlags().IntP("port", "p", 10880, "port to listen on")

	if err := viperBind(serveCmd, "listen"); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error binding listen flag: %s\n", err)
		os.Exit(1)
	}

	if err := viperBind(serveCmd, "port"); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error binding port flag: %s\n", err)
		os.Exit(1)
	}
}
