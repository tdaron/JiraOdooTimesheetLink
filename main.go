package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"jira-timesheet/jira"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jiraRequest, err := jira.NewWorklogRequest(request.Body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	} else {
		respString, _ := json.Marshal(jiraRequest)
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(respString),
		}, nil
	}

}


func main() {
	lambda.Start(Handler)
}