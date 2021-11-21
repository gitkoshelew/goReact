Simple way to init postgres instance:
1. Create docker container  
`docker run --name goreact-pg -e POSTGRES_PASSWORD=1265 -e POSTGRES_DB=goreact -p 8081:5432 -d postgres`
2. Create database `goreact`
3. Init database with `initdb.sql`
4. Change password for app user
`ALTER USER goreact_app WITH PASSWORD 'changeme';`
5. Update password in `config.yaml`