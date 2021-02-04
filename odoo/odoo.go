package odoo

import (
	"errors"
	odoo "jira-timesheet/odoo_api_wrapper"
	"os"
)
func CreateTimesheetLine(name string, email string, timesheetCode string, hours float64, date string) error{
	c, err := odoo.NewClient(&odoo.ClientConfig{
		Admin:    os.Getenv("ODOO_ADMIN"),
		Password:  os.Getenv("ODOO_PASSWORD"),
		Database:  os.Getenv("ODOO_DB"),
		URL:       os.Getenv("ODOO_URL"),
	})
	if err != nil {
		return errors.New("Can't connect to Odoo, Error: "+err.Error())
	}
	var taskSearch = odoo.Criteria{}
	taskSearch.Add("code","=",timesheetCode)
	task, err := c.FindProjectTask(&taskSearch)
	if err != nil || task == nil{
		return errors.New("Can't find any Odoo task with "+timesheetCode+" code")
	}


	var employeeSearch = odoo.Criteria{}
	employeeSearch.Add("work_email","=",email)
	employee, err := c.FindHrEmployee(&employeeSearch)
	if err != nil || employee == nil{
		return errors.New("Can't find employee with email "+email)
	}



	var accountAnalyticLine = odoo.AccountAnalyticLine{
		EmployeeId: odoo.NewMany2One(employee.Id.Get(), employee.Name.Get()),
		TaskId: odoo.NewMany2One(task.Id.Get(), task.Name.Get()),
		ProjectId: task.ProjectId,
		UnitAmount: odoo.NewFloat(hours),
		UserId: employee.UserId,
		Date: odoo.NewString(date),
		Name: odoo.NewString(name),
		ParentId: task.ParentId,
		}

	_, err = c.CreateAccountAnalyticLine(&accountAnalyticLine)

	if err != nil {
		return errors.New("Can't create Analytic Line. Error: "+err.Error())
	}
	return nil



}