# lecture 8 hw

in order to run 
- create .env file and fill with values like in .env.example
- run tables.sql, in order to create users table 
- run application with ```go run main.go```

### create user 
- ```curl --location 'localhost:8080/user/signup' --header 'Content-Type: application/json' --data-raw '{ "first_name": "dsda", "last_name": "sdd", "email": "aris@gmail.com", "password": "123456789"}'```
### login user  
- ```curl --location 'localhost:8080/user/login' --header 'Content-Type: application/json' --data-raw '{"email": "aris@gmail.com","password": "123456789"}'```
### update user 
- ```curl --location --request PUT 'localhost:8080/user/v2/update/1' --header 'Authorization: tokenmokenkoken' --header 'Content-Type: application/json' --data-raw '{"first_name": "ali","last_name": "aliemes","email": "aris@gmail.com","password": "123456789"}'```
### delete user 
- ```curl --location --request DELETE 'localhost:8080/user/v2/delete/1' --header 'Authorization: tokenmokenkoken'```
