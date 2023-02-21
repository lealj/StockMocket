guide used: (https://www.youtube.com/watch?v=KPftgI40WHI&t=776s)

IDE Should prompt you install packages in imports

Install mysql workbench default settings (https://www.youtube.com/watch?v=u96rVINbAUI). 
Having only server/workbench should work, but I just left everything on default.

Create local database on workbench, port should be 3306 by default. Should have to make password. Create schema named "godb" (cylinder button at top)

In dns.go, fill in password in this line (remove <>)

	const DNS = "root:<password>@tcp(localhost:3306)/godb?charset=utf8&parseTime=True&loc=Local"

in ide/terminal, 
	
	go build
	./restapi

It should output "connection successful" if it connected to database successfuly. 

To test sending the api json data, install Postman https://www.postman.com/downloads/ or use web version (i had issues with web)
Put this in "Enter request url"
	
	http://localhost:9000/(directory) (If you wanted to create a user, you would put "users" in (directory) - remove ()

Select "POST" next to "Enter request url" to send data
In body data, fill in json data

	{
		"firstname":"John"
		"lastname":"Smith"
		"email":"..."
		etc.
	}

Hit send. 
After this, refresh schemas in sql workbench, and you should see a table has been made under the schema. Checking the table should have info sent (name, etc.)

Getting information from the api:
Put "GET" next to enter request url. 
In request url, 

	http://localhost:9000/(directory)/(identifier)

Here json body data can be empty. 
If you wanted to get user data for user with id 1, you would put 

	http://localhost:9000/users/1

(id is created automatically)