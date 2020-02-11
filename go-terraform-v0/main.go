package terramod

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/terraform-providers/terraform-provider-aws/aws"
)

var code string

const stateFilename = "simple.tfstate"

func main() {
	count := 0
	/*
		keyName := "demo"
	*/
	filestoread := checkExt()
	ReadFile(filestoread)
	/*
		log := log.New(os.Stderr, "", log.LstdFlags)
		logMiddleware := logger.NewMiddleware()
		defer logMiddleware.Close()
	*/

	platform, err := NewPlatform(code).
		AddProvider("aws", aws.Provider()).
		PersistStateToFile(stateFilename)

	if err != nil {
		log.Fatalf("Fail to create the platform using state file %s. %s", stateFilename, err)
	}

	terminate := (count == 0)
	if err := platform.Apply(terminate); err != nil {
		log.Fatalf("Fail to apply the changes to the platform. %s", err)
	}

}

func ReadFile(filesToRead []string) string {
	/* Read all terraform files */
	tfdirectory := "terraform/"
	// Read Variables Files
	var tfvarFile []byte

	var dataFile []byte

	var err error

	for _, file := range filesToRead {
		tfvarFile, err = ioutil.ReadFile(tfdirectory + file)
		if err != nil {
			log.Fatalf("Fail read file %v, %v", file, err)
		}
		dataFile = append(dataFile, tfvarFile...)
	}
	return string(dataFile)
}

func checkExt() []string {

	tfdirectory := "terraform/"
	tfExtension := ".tf"
	var files []string
	filepath.Walk(tfdirectory, func(path string, fileInDir os.FileInfo, _ error) error {
		if !fileInDir.IsDir() {
			if filepath.Ext(path) == tfExtension {
				files = append(files, fileInDir.Name())
			}
		}
		return nil
	})
	return files
}

func init() {
	code = `
  provider "aws" {
    region        = "us-west-2"
  }
  resource "aws_instance" "server" {
    instance_type = "t2.micro"
	ami           = "ami-0d1cd67c26f5fca19"
  }
`
}
