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

 
 

      
    
 


