 - ✅ [Golang-Core](https://github.com/Denes-cilwal/Base-Go)


### Go Clean-Architecture Implementation
Clean Architecture with Golang with Dependency Injection

### Run Migration Commands
> ⚓️ &nbsp; Add argument `p=host` after `make` if you want to run the migration runner from the host environment instead of docker environment. example; `make p=host migrate-up`

If you are not using docker; ensure that sql-migrate is installed to use migration from the host environment.
To install sql-migrate:
> go get -v github.com/rubenv/sql-migrate/...

<details>
    <summary>Migration commands available</summary>

| Command              | Desc                                                       |
| -------------------- | ---------------------------------------------------------- |
| `make migrate-status`| Show migration status                                      |
| `make migrate-up`    | Migrates the database to the most recent version available |
| `make migrate-down`  | Undo a database migration                                  |
| `make redo`          | Reapply the last migration                                 |
| `make create`        | Create new migration file                                  |

</details>


### Run app with docker
- Update database env variables with credentials defined in `docker-compose.yml`
- Start server using command `docker-compose up -d` or `sudo docker-compose up -d` if permission issues
    > Assumes: Docker is already installed in the machine. 


### Twilio-Go 
### CICD  (With GithubActions)
### Sending Email with Go via Gmail API and OAuth2

###  TODO 
      - Upload Middleware
      - sentry integration
      - Cobra CLI
      - Firebase authentication

 
      
    
      
| Jobs                  | Endpoints                                                  |
| --------------------  | ---------------------------------------------------------- |
| `Create Job`          |    {{BASE_URL}}/job                                        |
| `Get all Jobs`        |    {{BASE_URL}}/job/all                                    |
| `Get Single Job`      |    {{BASE_URL}}/job/single/:id                             |
| `Delete Job`          |    {{BASE_URL}}/job/delete-job/:id                         |
| `Approve&Decline-Job` |    {{BASE_URL}}/job/approve-decline/:jobId                 |
| `Update Public Status`|    {{BASE_URL}}/job//toggle-public-status/:jobId           |

| Job-Request           | Endpoints                                                  |
| --------------------  | ---------------------------------------------------------- |
| `Create Job Request`  |    {{BASE_URL}}/job-request                                |
| `Get Single `         |    {{BASE_URL}}/job-request/single                         |
| `Update`              |    {{BASE_URL}}/job-request/:id                            |

| Job-For               | Endpoints                                                  |
| ----------------------| ---------------------------------------------------------- |
| `Create Job for`      |    {{BASE_URL}}/job-for                                    |
| `Get Single `         |    {{BASE_URL}}/job-for/single                             |
| `Update`              |    {{BASE_URL}}/job-for/:id                                |

| Jobs-Interest              | Endpoints                                                  |
| -------------------------- | ---------------------------------------------------------- |
| `Add Job interest-user`    |    {{BASE_URL}}/job-interest                               |
| `Get Job interest-user`    |    {{BASE_URL}}/job-interest/all                           |
| `Update Job interest-user `|    {{BASE_URL}}/job-interest                               |
| `Delete Job interest users`|    {{BASE_URL}}/job-interest                               |



