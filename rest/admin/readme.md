1. run command
 `migrate -path ./migrations -database 'postgres://postgres:changeme@localhost:8081/postgres?sslmode=disable' up` 

2. How to come migrations back 
 `migrate -path ./migrations -database 'postgres://postgres:changeme@localhost:8081/postgres?sslmode=disable' down` 

3. How to fix database 
`migrate -path ./migrations -database 'postgres://user:userpass@0.0.0.0:8087/admindb?sslmode=disable' force 1 `

4. add migr files
`migrate create -ext sql -dir migrations -seq add_photo_for_pet_room `

5. Copy file from container:
`docker cp microserviceadmin_microseviceadmin_1:api/pkg/csv/allusers.csv files`


