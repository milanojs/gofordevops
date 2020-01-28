package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var username, password string

func useCredentials(username, password string) {
	fmt.Printf("Connecting to %s at %s\n",
		viper.GetString("webserver.name"),
		viper.GetString("webserver.endpoint"),
	)
	fmt.Printf("Your Username is: %v\n", username)
	fmt.Printf("Your Password is: %v\n", password)
}
func useCredentialsWithViper() {
	fmt.Printf("Your Username is: %s\n", viper.GetString("credentials.username"))
	fmt.Printf("Your Password is: %s\n", viper.GetString("credentials.username"))
}

func main() {
	flag.StringVar(&username, "user", "", "Insert the user name")
	flag.StringVar(&password, "password", "", "Insert the user password")
	flag.Parse()
	/*Setting default values for viper*/
	viper.SetDefault("webserver.endpoint", "https://httpbin.org/get")
	viper.SetDefault("webserver.name", "Testing Endpoint")
	viper.AddConfigPath(".")
	viper.SetConfigName("creds")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Unable to read config")
	}
	if username != "" {
		viper.Set("credentials.username", username)
	}
	if password != "" {
		viper.Set("credentials.password", password)
	}
	if !viper.IsSet("credentials.username") || !viper.IsSet("credentials.password") {
		log.Fatalln("No credentials supplied")
	}

	viper.Set("credentials.username", username)
	viper.Set("credentials.password", password)
	useCredentials(username, password)
	fmt.Println("Clean line")
	useCredentialsWithViper()

}
