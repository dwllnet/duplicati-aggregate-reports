# duplicati-aggregate-reports

A GoLang web server to aggregate Duplicati Backup Reports

Web server built using Blueprint (https://github.com/blue-jay/blueprint) a model-view-controller (MVC) style web skeleton. 


Web Application to aggregate n+1 Duplicati backup client into a single databse.
Duplicati --send-http-url parameter http://this server's ip/clientstatus/<IP address of Duplicati Machine>
Example: http://10.10.10.10/clientstatus/10.10.10.100
