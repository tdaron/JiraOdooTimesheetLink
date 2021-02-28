# Fast Setup

First, link aws cli with your AWS account

Then,
- Create a lambda function with API Gateway, and replace Handler , from hello to main ( in Execution Settings ) 
- Update lambda environment variables ( ODOO_ADMIN, ODOO_PASSWORD, ODOO_DB , ODOO_URL , ODOO_WHITELIST ) --> ODOO_WHITELIST IS NOT REQUIRED, BUT USEFUL FOR TESTS
- Replace TestFunction by your lambda name in deploy.sh
- Execute deploy.sh

And now, you just have to go Jira, and add webhook ( issue>weblog_changed) with your API Gateway Public Endpoint
( It will only works on issues with Timesheet Code field ) 

And it should work !

# Environment Variables

## ODOO_ADMIN
This is your odoo username ( user have to be able to edit all timesheets ) 
## ODOO_PASSWORD
Your odoo password
## ODOO_DB
Put the odoo's database name
## ODOO_URL
Here is the url of your odoo
## ODOO_WHITELIST
The whitelist is, when defined, a list of projects lambda will care, your lambda'll ignore all projects outside this list

Format: TIMESHEETCODE,HEREANOTHER,ANOTHERONE

Exemple: DFA-DF2-P1,DFA-DF2-P1

## USERS_WHITELIST
This is the users whitelist, if this env var is set, your lambda'll refuse all requests for users that are not in this whitelist

Format: email,email2,email3

Exemple: alice@github.com,bob@github.com

## TEAMS_WEBHOOK_URL
That is the incoming webhook url for live notifications in Microsoft Teams ( Success and Errors)

