package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sfn"
)

func main() {
	flag.Parse()
	args := flag.Args()

	token := os.Getenv("TASK_TOKEN_ENV_VARIABLE")
	fmt.Printf("token: %s\n", token)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		fmt.Println(err)
	}

	sfnClient := sfn.New(sess)

	json := `
	{
	"field1": "test1",
	"field2": "test2"
	}
	`

	isSuccess := args[0]
	if isSuccess == "true" {
		_, err = sfnClient.SendTaskSuccess(&sfn.SendTaskSuccessInput{
			Output:    aws.String(json),
			TaskToken: aws.String(token),
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	_, err = sfnClient.SendTaskFailure(&sfn.SendTaskFailureInput{
		Error:     aws.String("task failure"),
		TaskToken: aws.String(token),
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
