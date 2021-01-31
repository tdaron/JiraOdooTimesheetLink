package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"jira-timesheet/jira"
	odoo "jira-timesheet/odoo"
	"jira-timesheet/teamsnotifier"
	"jira-timesheet/utils"
	"os"
	"strings"
)

func ErrorResponse(e error) events.APIGatewayProxyResponse {
	teamsnotifier.Notify(e)

	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       e.Error(),
	}
}

func preventOdooCrash(){
	if !odooCredsOk{
		teamsnotifier.Notify(errors.New("Odoo connection error , please check your creds"))
	}
}

var odooCredsOk = false


func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jiraRequest, err := jira.NewWorklogRequest(request.Body)
	if err != nil {
		return ErrorResponse(err), nil
	}

	var odoo_whitelist = os.Getenv("ODOO_WHITELIST")
	if  odoo_whitelist != "" {
		var white_projects = strings.Split(odoo_whitelist,",")
		if !utils.SliceContains(white_projects, jiraRequest.Issue.Fields.TimesheetCode) {
			return ErrorResponse(errors.New(jiraRequest.Issue.Fields.TimesheetCode+" is not in the whitelist ( "+odoo_whitelist+")")), nil
		}
	}



	var lastWorklog = jiraRequest.GetLastWorklog()

	defer preventOdooCrash() ////odoo panic if authentication fails, we have to notify user of bad creds

	err = odoo.CreateTimesheetLine(
		jiraRequest.Issue.Key,
		jiraRequest.User.Email,
		jiraRequest.Issue.Fields.TimesheetCode,
		lastWorklog.GetHours(),
		lastWorklog.GetDate())
	odooCredsOk = true

	if err != nil {
		return ErrorResponse(err), nil
	}

	teamsnotifier.Notify(errors.New("SUCCESS"))
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil


}


func main() {
	lambda.Start(Handler)
}