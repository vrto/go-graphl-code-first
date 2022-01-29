### Code-first GraphQL with Go

I wrote this small program to learn how to write a GraphQL server in Go.

It defines a tiny domain (Pubs with Beers) and runs a bunch of hardcoded queries and mutations. The results are in stdout.

**Tech used:**
- Go 1.17
- [`graphql-go`](https://github.com/graphql-go/graphql) library, GraphQL schema defined in code
- PostgreSQL, the Docker image w/ initial SQL structure is in `pgsql` folder

### Running

- create PostgreSQL Docker image that contains `db.sql` script with initial schema
  - `$ docker build pgsql -t pgsql`
- run the Docker container with this modified Postgres (`pgsql` is the image name from above)
  - `$ docker run --name pgsql-pubdb -p5432:5432 -d pgsql`
- build the Go app
  - `$ go build .`
- run the compiled binary
  - `$ ./pub`

Inspect PostgreSQL data by running `psql` like so:
```
$ docker exec -it pgsql-pubdb psql -U gopher pubdb
```