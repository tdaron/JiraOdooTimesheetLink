package jira

import "encoding/json"

//Types for JSON Deserialization
type JiraWorklog struct {
	TimeSpentSeconds int `json:"timeSpentSeconds"`
}

type JiraWorklogsField struct {
	Worklogs []JiraWorklog `json:"worklogs"`

}
type JiraRequestFields struct {
	TimesheetCode string `json:"customfield_10101"`
	Worklog JiraWorklogsField `json:"worklog"`
}

type JiraUser struct {
	Key string `json:"key"`
	Email string `json:"emailAddress"`
}

type JiraIssue struct {
	Key string `json:"key"`
	Fields JiraRequestFields `json:"fields"`

}

type JiraRequest struct {
	User JiraUser `json:"user"`
	Issue JiraIssue `json:"issue"`

}

func NewWorklogRequest(body string) (JiraRequest, error) {
	var jiraRequest JiraRequest
	err := json.Unmarshal([]byte(body),&jiraRequest)
	return jiraRequest, err
}

func (r *JiraRequest) GetLastWorklogTime() int {
	var worklogs = r.Issue.Fields.Worklog.Worklogs
	var lastWorklog = worklogs[len(worklogs)-1]
	return lastWorklog.TimeSpentSeconds
}