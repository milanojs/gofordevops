package main

import (
	"fmt"
	//"errors"
	
)

func checkempty(tagvalue, tagname string) (string, error) {
	if tagvalue == "" {
		//message := fmt.Errorf("Unable to create Deploy, *%v* not specified in the manifest file", tagvalue)
		//return "", errors.New(message)
		return "", fmt.Errorf("Unable to create Deploy, * %v * not specified in the manifest file", tagname)

	}
	return tagvalue + " " + tagname, nil
}

func checktagserr(tagname, tagvalue string) string {

	value, err := checkempty(tagname, tagvalue)
	if err != nil {
		fmt.Println(err)
	}
	return value
}


func main() {
	a := ""
	endingvalue := checktagserr(a, "Business")
	fmt.Println(endingvalue)
}
