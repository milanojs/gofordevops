// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func CheckWhiteSpacesInTags(Valstring string) bool {
	CheckWhiteSpace := regexp.MustCompile(`\s`)
	matched := CheckWhiteSpace.MatchString(Valstring)
	if matched == true {
		fmt.Printf("There was an error, Tag provided must not have any spaces, Please Checkit *%v*\n", Valstring)
		return true
	} else {
		return false
	}
	CheckIfLowerCase := strings.ToLower(Valstring) == Valstring
	if CheckIfLowerCase == true {
		return true
	} else {
		fmt.Printf("There was an error, Tag provided must not have all character in lowercase, Please Checkit *%v*\n", Valstring)
		return false
	}

}
func ContainsTagsProposed(arr []string, str string) (bool, error) {
	for _, a := range arr {
		if a == str {
			return true, nil
		}
	}
	message := "The Proposed Tags can't be used, Please check the proposed values *%s*"
	return false, fmt.Errorf(message, str)
}

func main() {

	TagToTestValue := "nxt"

	CheckWhiteSpacesInTags(TagToTestValue)
	b, err := ContainsTagsProposed(arrayProposedTagsBu, TagToTestValue)
	if err == nil {
		fmt.Println(b)
	} else {
		fmt.Println(err)
	}
}
