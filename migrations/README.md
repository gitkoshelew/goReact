How to do migrations 
1. run command
 `migrate -path ./migrations -database 'postgres://postgres:changeme@localhost:8081/postgres?sslmode=disable' up` 
2. How to come migrations back 
 `migrate -path ./migrations -database 'postgres://postgres:changeme@localhost:8081/postgres?sslmode=disable' down` 
3. How to fix database
`migrate -path ./migrations -database 'postgres://postgres:changeme@localhost:8081/postgres?sslmode=disable' force 1 `
4. How to connect to docker db conteiner 

`docker ps`
`docker exec -it ID /bin/bash`
`psql -U postgres`
`\d`


