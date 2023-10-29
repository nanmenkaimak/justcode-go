# lecture 9 hw

in order to run 
- create .env file and fill with values like in .env.example
- run application with ```go run main.go```

### create gist 
```curl --location 'localhost:8080/' --header 'Content-Type: application/json' --data '{"gist" : {"name": "second gist","description" : "ne che tam"},"commit" : {"comment" : "dep commit"},"files" : [{"name": "main.py","code" : "package main\n import \"fmt\"\nfunc main() {\n fmt.Println(\"ri e\")\n}  "}, {"name": "main2.py","code" : "package main\n import \"fmt\"\nfunc main() {\n fmt.Println(\"che \")\n}  "}]}'```
### get all gists 
```curl --location 'localhost:8080/all'```
### get gist by id 
```curl --location 'localhost:8080/{id}'```
### update gist by id
- ```curl --location --request PUT 'localhost:8080/{id}' --header 'Content-Type: application/json' --data '{"gist": {"name": "third emes gist","description": "che tam"},"commit": {"comment": "second commit"},"files": [{"name": "mainemses.py","code": "package main\n import \"fmt\"\nfunc main() {\n fmt.Printf(\"ri e\")\n}  "}, {"name": "main3.py","code": "package main\n import \"fmt\"\nfunc main() {\n fmt.Println(\"che tam emes \")\n}  "}]}'```
### delete gist by id 
- ```curl --location --request DELETE 'localhost:8080/{id}'```
