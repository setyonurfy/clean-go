## Directory Structure
...

/cmd                    //Entry point

/internal               //internal workflow

    /entity             //Business domain model

    /usecase            //Interface & main business logic

    /repository         //Interface (abstraction) pf data layer

    /infrastructure     //Implementation of data layer 

    /delivery           // Handler of usecase

/pkg                    //helper & utility

/migrations             //Database migrations

...

## Database

DB: PostgreSQL

`golang-migrate` -> create & execute migrations (migration tool)

### Install golang-migrate

Using Homebrew (macOS 11+):

```
brew install golang-migrate
```

Using Go (works on older macOS, e.g. Catalina):

```
GOTOOLCHAIN=local go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
```

Make sure `$(go env GOPATH)/bin` is in your `PATH`.

### Create the database

`migrate` does not create the database itself, so create it first:

```
createdb -U postgres -h localhost clean-go
```

or via `psql`:

```
CREATE DATABASE "clean-go";
```

### Migrations

Create a new migration:

```
./migrations/create-migration.sh create_users_table
```

Run migrations:

```
./migrations/migrate.sh up
```

Roll back the last migration:

```
./migrations/migrate.sh down 1
```

### Database connection
 
The connection is initialized in `internal/infrastructure/pgsql/init.go` using `sqlx` and the environment variables above. It is called from `cmd/main.go` when the app starts.
 
## Run
 
From the project root (the folder that contains `go.mod`):
 
```
go run cmd/main.go
```
 
If the connection works, you will see:
 
```
Successfully connected to the database
PGSQL database initialized successfully
```