package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var username, password string

func useCredentials(username, password string) {
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
