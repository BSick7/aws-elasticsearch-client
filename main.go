package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	awsauth "github.com/smartystreets/go-aws-auth"
	"gopkg.in/olivere/elastic.v5"
)

type allthethings struct{}

func (l *allthethings) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func main() {
	endpoint := os.Getenv("AWS_ES_ENDPOINT")

	signingTransport, err := getSigningTransport()
	if err != nil {
		log.Fatalf(err.Error())
	}
	signingClient := &http.Client{Transport: http.RoundTripper(signingTransport)}

	client, err := elastic.NewSimpleClient(
		elastic.SetSniff(false),
		elastic.SetURL(endpoint),
		elastic.SetHttpClient(signingClient),
		elastic.SetTraceLog(&allthethings{}),
		elastic.SetInfoLog(&allthethings{}),
		elastic.SetErrorLog(&allthethings{}),
	)
	if err != nil {
		log.Fatalf("could not create elasticsearch client: %s", err)
	}

	//"https://search-mdl-alpha-cfo-yt4buq3ihh32l7v2bf2kf54qza.us-east-1.es.amazonaws.com"
	log.Println(client.Ping(endpoint).Do(context.Background()))
}

func getSigningTransport() (AwsSigningTransport, error) {
	/*
		creds, err := assumeRole("arn:aws:iam::295854699460:role/mdl-alpha-cfo-poweruser", "golang-test-client")
		if err != nil {
			return nil, fmt.Errorf("error assuming role: %s", err)
		}
	*/

	return AwsSigningTransport{
		Credentials: awsauth.Credentials{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),     //*creds.AccessKeyId,
			SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"), //*creds.SecretAccessKey,
			SecurityToken:   os.Getenv("AWS_SESSION_TOKEN"),     //*creds.SessionToken,
		},
		HttpClient: http.DefaultClient,
	}, nil
}

func assumeRole(role string, sessionName string) (*sts.Credentials, error) {
	s := session.Must(session.NewSession())
	svc := sts.New(s)
	out, err := svc.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         aws.String(role),
		RoleSessionName: aws.String(sessionName),
	})
	if err != nil {
		return nil, err
	}
	return out.Credentials, nil
}

type AwsSigningTransport struct {
	Credentials awsauth.Credentials
	HttpClient  *http.Client
}

func (a AwsSigningTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := a.HttpClient.Do(awsauth.Sign4(req, a.Credentials))
	if res != nil && res.StatusCode >= 300 {
		raw, _ := ioutil.ReadAll(res.Body)
		log.Println(res.StatusCode, string(raw))
		res.Body.Close()
	}
	return res, err
}
