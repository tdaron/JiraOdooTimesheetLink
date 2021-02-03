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
	"reflect"
	"strings"
)

func ErrorResponse(request jira.Request, e error) events.APIGatewayProxyResponse {
	var blank = jira.Request{}

	if !reflect.DeepEqual(request, blank) {
		teamsnotifier.Notify(teamsnotifier.NewOdooError(request, e))
	} else {
		teamsnotifier.Notify(teamsnotifier.NewJiraError(e))

	}

	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       e.Error(),
	}
}

func preventOdooCrash(request jira.Request){
	if !odooCredsOk{
		teamsnotifier.Notify(teamsnotifier.NewOdooError(request, errors.New("Odoo connection error , please check your creds")))
	}
}

var odooCredsOk = false


func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jiraRequest, err := jira.NewWorklogRequest(request.Body)
	if err != nil {
		return ErrorResponse(jira.Request{}, err), nil
	}

	var odoo_whitelist = os.Getenv("ODOO_WHITELIST")
	if  odoo_whitelist != "" {
		var white_projects = strings.Split(odoo_whitelist,",")
		if !utils.SliceContains(white_projects, jiraRequest.Issue.Fields.TimesheetCode) {
			return ErrorResponse(jiraRequest, errors.New(jiraRequest.Issue.Fields.TimesheetCode+" is not in the whitelist ( "+odoo_whitelist+")")), nil
		}
	}



	var lastWorklog = jiraRequest.GetLastWorklog()

	defer preventOdooCrash(jiraRequest) ////odoo panic if authentication fails, we have to notify user of bad creds

	err = odoo.CreateTimesheetLine(
		jiraRequest.Issue.Key,
		jiraRequest.User.Email,
		jiraRequest.Issue.Fields.TimesheetCode,
		lastWorklog.GetHours(),
		lastWorklog.GetDate())
	odooCredsOk = true

	if err != nil {
		return ErrorResponse(jiraRequest, err), nil
	}

	teamsnotifier.Notify(teamsnotifier.NewSuccess(jiraRequest))
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "OK",
	}, nil


}


func main() {
	lambda.Start(Handler)
}