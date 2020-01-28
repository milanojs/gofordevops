package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "cobraintro",
	Short: "This tools gets a URL with basic auth",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("Must set URL!")
		}
		client := http.Client{}
		req, err := http.NewRequest("GET", args[0], nil)

		if err != nil {
			log.Fatalln("Unable to get request")
		}
		if username != "" && password != "" {
			req.SetBasicAuth(username, password)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln("Unable to get response")
		}
		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("Unable to read body")
		}
		fmt.Println(string(content))

	},
}

var username, password string

func main() {
	cmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username for credentials")
	cmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password for credentials")
	cmd.Execute()
}
