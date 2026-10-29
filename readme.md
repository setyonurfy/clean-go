## Directory Structure

```
/cmd                    //Entry point
/internal               //internal workflow
    /entity             //Business domain model
    /usecase            //Interface & main business logic
    /repository         //Interface (abstraction) of data layer
    /infrastructure     //Implementation of data layer
    /delivery           //Handler of usecase
/pkg                    //helper & utility
/migrations             //Database migrations
```

## dependency

gin ->  `go get github.com/gin-gonic/gin`

pq  ->  `go get github.com/lib/pq`

godotenv  ->  `go get github.com/joho/godotenv`

sqlx  ->  `go get github.com/jmoiron/sqlx`

jwt  ->  `go get github.com/golang-jwt/jwt/v5`

bcrypt  ->  `go get golang.org/x/crypto`

## Database

DB: Postgresql

`golang-migrate` -> Create & Execute migration (migration tool)

### Install golang-migrate

Using homebrew (macOS 11+):

```
brew install golang-migrate
```

Using Go (also works on older macOS, e.g. Catalina):

```
GOTOOLCHAIN=local go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
```

Make sure the Go bin folder is in your `PATH`:

```
export PATH=$PATH:$(go env GOPATH)/bin
```

### Create the database

`migrate` does not create the database itself, so create it first:

```
createdb -U postgres -h localhost clean-go
```

or via `psql`:

```
CREATE DATABASE "clean-go";
```

### Environment variables

Create a `.env` file in the project root (already ignored by `.gitignore`):

```
DATABASE_USER=postgres
DATABASE_PASSWORD=postgres
DATABASE_NAME=clean-go
DATABASE_HOST=localhost
DATABASE_PORT=5432
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

## User & Auth

Implemented following the layers in the directory structure above.

| Layer | File | Responsibility |
| --- | --- | --- |
| Entity | `internal/entity/user.go` | `User` struct (ID, Username, Email, Password, CreatedAt, UpdatedAt) |
| Repository | `internal/repository/user_repository.go` | `UserRepository` interface (abstraction of the data layer) |
| Usecase | `internal/usecase/auth_usecase.go` | `AuthUsecase` interface and business logic for login & register |

### UserRepository

Defines the contract only. The implementation that talks to PostgreSQL goes in `internal/infrastructure/pgsql`.

`repository.ErrNotFound` is returned when a record does not exist. The implementation must return it when the query finds no rows (`sql.ErrNoRows`), so the usecase can tell "not found" apart from a real database error.

```go
var ErrNotFound = errors.New("not found")

type UserRepository interface {
	All() ([]*entity.User, error)
	GetByEmail(string) (*entity.User, error)
	GetById(string) (*entity.User, error)
	Create(*entity.User) error
}
```

### AuthUsecase

```go
type AuthUsecase interface {
	Login(email, password string) (string, error)
	Register(username, email, password string) error
}
```

- `Register` checks that the email is not used yet (`GetByEmail` returning `ErrNotFound` means it is free), hashes the password with bcrypt, then saves the user through `UserRepository.Create`.
- `Login` finds the user by email, compares the password with the bcrypt hash, and returns a signed JWT (HS256, valid for 24 hours) containing `user_id`. A wrong email and a wrong password return the same error message (`invalid email or password`).
- `NewAuthUsecase(repo, jwtSecret)` receives the repository and the JWT secret, so the usecase does not depend on the database directly and can be tested without one.

### Status

- [x] Entity
- [x] Repository interface
- [x] Auth usecase (login & register)
- [ ] Repository implementation (pgsql)
- [ ] Delivery (HTTP handler)