package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmd = &cobra.Command{
	Use:   "cobraintro",
	Short: "This tools gets a URL with basic auth",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatalln("must Use a sub command")
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a URL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("Must set URL!")
		}
		client := http.Client{}
		req, err := http.NewRequest("GET", args[0], nil)

		username := viper.GetString("username")
		password := viper.GetString("password")

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

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Post a URL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatalln("Must set URL!")
		}
		client := http.Client{}
		var contentReader io.Reader
		content := viper.GetString("content")

		if content != "" {
			contentReader = bytes.NewReader([]byte(content))
		}
		req, err := http.NewRequest("POST", args[0], contentReader)

		if err != nil {
			log.Fatalln("Unable to post request")
		}

		username := viper.GetString("username")
		password := viper.GetString("password")

		if username != "" && password != "" {
			req.SetBasicAuth(username, password)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln("Unable to post response")
		}
		defer resp.Body.Close()

		respContent, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("Unable to read body")
		}
		fmt.Println(string(respContent))
	},
}

func main() {
	cmd.PersistentFlags().StringP("username", "u",
		viper.GetString("credentials.username"), "Username for credentials")
	cmd.PersistentFlags().StringP("password", "p",
		viper.GetString("credentials.username"), "password for credentials")
	viper.BindPFlag("username", cmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", cmd.PersistentFlags().Lookup("password"))
	postCmd.PersistentFlags().StringP("content", "c", "", "The content for post")
	viper.BindPFlag("content", cmd.PersistentFlags().Lookup("content"))
	cmd.AddCommand(getCmd)
	cmd.AddCommand(postCmd)
	cmd.Execute()
}
func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("cobra")
	viper.ReadInConfig()
}
