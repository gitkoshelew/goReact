Simple way to init postgres instance:
1. run docker-compose 
`docker-compose up --build`
or run docker-compose for developers
`docker-compose --file docker-compose-dev.yml`
2. Create database `goreact`
3. Init database with `initdb.sql`
4. Change password for app user
`ALTER USER goreact_app WITH PASSWORD 'changeme';`
5. Update password in `config.yaml`
6. Copy logs from container:
`docker cp goreact_goreact_1:api/logs /logs`  
7. Run APIs tests:
`go test -v -race -timeout 30s ./...`