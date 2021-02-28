package jira

import (
	"encoding/json"
)


//Types for JSON Deserialization
type Worklog struct {
	TimeSpentSeconds int `json:"timeSpentSeconds"`
	Date string `json:"started"`
	TimeSpent string `json:"timeSpent"`
}

func (w *Worklog) GetDate() string {
	return w.Date[0:10]
}

func (w *Worklog) GetHours() float64 {
	return float64(w.TimeSpentSeconds) / 3600
}

type WorklogsField struct {
	Worklogs []Worklog `json:"worklogs"`

}
type RequestFields struct {
	TimesheetCode string        `json:"odoo_timesheet_code"` //CUSTOM FIELD WITH ODOO TIMESHEET CODE
	Worklog       WorklogsField `json:"worklog"`
}

type User struct {
	Key string `json:"key"`
	Email string `json:"emailAddress"`
}

type Issue struct {
	Key    string        `json:"key"`
	Fields RequestFields `json:"fields"`

}

type Request struct {
	User  User  `json:"user"`
	Issue Issue `json:"issue"`

}

func NewWorklogRequest(body string) (Request, error) {
	var jiraRequest Request
	err := json.Unmarshal([]byte(body),&jiraRequest)
	return jiraRequest, err
}

func (r *Request) GetLastWorklog() Worklog {
	var worklogs = r.Issue.Fields.Worklog.Worklogs
	var lastWorklog = worklogs[len(worklogs) -1]
	return lastWorklog
}