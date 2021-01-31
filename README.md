#Setup

First, link aws cli with your AWS account

Then,
- Create a lambda function with API Gateway, and replace Handler , from hello to main ( in Execution Settings ) 
- Update lambda environment variables ( ODOO_ADMIN, ODOO_PASSWORD, ODOO_DB , ODOO_URL , ODOO_WHITELIST ) --> ODOO_WHITELIST IS NOT REQUIRED, BUT USEFUL FOR TESTS
- Replace TestFunction by your lambda name in deploy.sh
- Execute deploy.sh

And now, you just have to go Jira, and add webhook ( issue>weblog_changed) with your API Gateway Public Endpoint
( It will only works on issues with Timesheet Code field ) 

And it should work !
