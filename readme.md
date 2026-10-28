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

## dependency
gin ->  'go get github.com/gin-gonic/gin'

pq  ->  'go get github.com/lib/pq'

godotenv  ->  'go get github.com/joho/godotenv'

sqlx  ->  'go get github.com/jmoiron/sqlx'

## Database

DB: Postgresql

'golang-migrate' -> Create & Execute migration (migration tool)

to install 'golang-migrte' using homebrew

...

brew install golang-migrate
...