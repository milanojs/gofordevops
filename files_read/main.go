package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

/*Commented struct in order to user viper*/
/*
type Response struct {
	URL    string            `json:"url"`
	Header map[string]string `json:"headers"`
	origin string            `json:"origin"`
}
*/

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln("Unable to get response!")
	}
	defer resp.Body.Close()
	viper.SetConfigType("json")
	if err := viper.ReadConfig(resp.Body); err != nil {
		log.Fatalln("Unable to read Body")
	}
	fmt.Println(viper.AllSettings())
	m := viper.GetStringMapString("headers")
	fmt.Printf("==\n")

	for k, v := range m {
		fmt.Printf("%s: %s\n", k, v)
	}
	/*Commented in order to user viper*/
	/*
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("Unable to read Content!")
		}
		respContent := Response{}
		json.Unmarshal(content, &respContent)
		fmt.Printf("From %s\n", respContent.URL)
		for k, v := range respContent.Header {
			fmt.Printf("%s: %s\n", k, v)
		}
		fmt.Printf("From %s\n", respContent.origin)

		//fmt.Println(string(content))
	*/
}
