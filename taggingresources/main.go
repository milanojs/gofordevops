package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go/aws/awserr"

)

const (
	// ProviderName is the cloud provider providing loadbalancing functionality
	ProviderName = "aws"
)

// ELB is the struct implementing the lbprovider interface
type ELB struct {
	client            *elb.ELB
	elbResourceName   string
	resourceapiClient *resourcegroupstaggingapi.ResourceGroupsTaggingAPI
}
type Value struct{
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}


// NewELB is the factory method for ELB
func NewELB(id string, secret string, region string) (*ELB, error) {
	awsConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(id, secret, ""),
	}

	awsConfig = awsConfig.WithCredentialsChainVerboseErrors(true)
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("Unable to initialize AWS session: %v", err)
	}

	return &ELB{
		client:            elb.New(sess),
		resourceapiClient: resourcegroupstaggingapi.New(sess),
		elbResourceName:   "elasticloadbalancing",
	}, nil
}

// GetLoadbalancersByTag gets the loadbalancers by tag
func (e *ELB) GetLoadbalancersByTag(tagKey string, tagValue string) ([]string, error) {
	tagFilters := &resourcegroupstaggingapi.TagFilter{}
	tagFilters.Key = aws.String(tagKey)
	tagFilters.Values = append(tagFilters.Values, aws.String(tagValue))

	getResourcesInput := &resourcegroupstaggingapi.GetResourcesInput{}
	getResourcesInput.TagFilters = append(getResourcesInput.TagFilters, tagFilters)
	getResourcesInput.ResourceTypeFilters = append(
		getResourcesInput.ResourceTypeFilters,
		aws.String(e.elbResourceName),
	)

	resources, err := e.resourceapiClient.GetResources(getResourcesInput)
	if err != nil {
		return nil, err
	}

	elbs := []string{}
	for _, resource := range resources.ResourceTagMappingList {
		elbs = append(elbs, strings.Split(*resource.ResourceARN, "/")[1])
	}
	return elbs, nil
}
func main(){
	creds := credentials.NewEnvCredentials()

	// Retrieve the credentials value
	credValue, err := creds.Get()
	if err != nil {
		// handle error
	}

}

var Credentials = new AWS.Credentials

var (
	// ErrAccessKeyIDNotFound is returned when the AWS Access Key ID can't be
	// found in the process's environment.
	//
	// @readonly
	ErrAccessKeyIDNotFound = awserr.New("EnvAccessKeyNotFound", "AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment", nil)

	// ErrSecretAccessKeyNotFound is returned when the AWS Secret Access Key
	// can't be found in the process's environment.
	//
	// @readonly
	ErrSecretAccessKeyNotFound = awserr.New("EnvSecretNotFound", "AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY not found in environment", nil)
)

// A EnvProvider retrieves credentials from the environment variables of the
// running process. Environment credentials never expire.
//
// Environment variables used:
//
// * Access Key ID:     AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY
// * Secret Access Key: AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY
type EnvProvider struct {
	retrieved bool
}

// NewEnvCredentials returns a pointer to a new Credentials object
// wrapping the environment variable provider.
func NewEnvCredentials() *Credentials {
	return NewCredentials(&EnvProvider{})
}

// Retrieve retrieves the keys from the environment.
func (e *EnvProvider) Retrieve() (Value, error) {
	e.retrieved = false

	id := os.Getenv("AWS_ACCESS_KEY_ID")
	if id == "" {
		id = os.Getenv("AWS_ACCESS_KEY")
	}

	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secret == "" {
		secret = os.Getenv("AWS_SECRET_KEY")
	}

	if id == "" {
		return Value{}, ErrAccessKeyIDNotFound
	}

	if secret == "" {
		return Value{}, ErrSecretAccessKeyNotFound
	}

	e.retrieved = true
	return Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
		SessionToken:    os.Getenv("AWS_SESSION_TOKEN"),
	}, nil
}

// IsExpired returns if the credentials have been retrieved.
func (e *EnvProvider) IsExpired() bool {
	return !e.retrieved
}
