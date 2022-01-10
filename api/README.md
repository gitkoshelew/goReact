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
6. Cope logs from container:
`docker cp goreact_goreact_1:api/logs /logs`  

Endpoints info:
-----------------------------------------------------------------------------
POST    /registration - register a new User

Request data:
{
    "email": string,            
    "password": string,
    "role": string,             NOT NULL # "client", "employee", "anonymous" #
    "verified": bool,           NOT NULL
    "name": string,             NOT NULL
    "sName": string,            NOT NULL
    "mName": string,            
    "sex": string,              NOT NULL 
    "birthDate": time.Time,     NOT NULL # format: "2111-01-01" #
    "address": string,          
    "phone": string,            NOT NULL
    "photo": string             
}

Response data:

STATUS 201:
{
    id: int
}

STATUS 400:
http.Error - "Email already in use", "Bad request"
-----------------------------------------------------------------------------
POST    /login - login

Request data:
{
		Email    string         NOT NULL
		Password string         NOT NULL
}

Response data:

STATUS 201:
Cookie:
{
			Name:     "JWT",
			Value:    tk.AccessToken,
			HttpOnly: true,
}
Header: access_token
JSON:
{
    "userId": int,
    "email": string,
    "role": string, 
    "verified": bool,       
    "name": string,       
    "sName": string,          
    "mName": string,            
    "sex": string,
    "birthDate": time.Time,     # format: "2111-01-01" #
    "address": string,          
    "phone": string,          
    "photo": string             
}

STATUS 401:
http.error: "Invalid login or password", "Bad request"