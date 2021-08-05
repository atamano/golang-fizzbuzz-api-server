# Fizz Buzz project

## Getting Started

### Prerequisites

- Docker: needed for local development
- fswatch (optional): for live reload

### Setup

Set env variables

```
$ cp .env.example .env
```

| Name                | Description                           |
| ------------------- | ------------------------------------- |
| `SERVER_PORT`       | Server port                           |
| `DATABASE_HOST`     | Database host                         |
| `DATABASE_PORT`     | Database port                         |
| `DATABASE_NAME`     | Database name                         |
| `DATABASE_USER`     | Database user                         |
| `DATABASE_PASSWORD` | Database password                     |
| `DATABASE_DEBUG`    | Set to true to debug queries          |
| `GIN_MODE`          | debug (local) or release (production) |

### Run in development

```
make start
```

To init database

```
$ make migration-init
$ make migration-up
```

To add live reload:

```
$ make livereload
```

## Running the tests

```
$ make test
```

## Try it

```
curl -d '{"int1": 4, "int2": 2, "limit": 20, "str1": "toto", "str2": "tata"}' localhost:8080/v1/fizzbuzz
```

```
curl -d '{"int1": 5, "int2": 3, "limit": 10, "str1": "foo", "str2": "bar"}' localhost:8080/v1/fizzbuzz
```

```
curl  localhost:8080/v1/stats
```

## Built With

- [Gin](https://github.com/gin-gonic/gin) - Http framework
- [Go pg](https://github.com/go-pg/pg) - Golang ORM with focus on PostgreSQL features and performance
- [Envconfig](https://github.com/kelseyhightower/envconfig) - Managing configuration data from environment variables
- [Logrus](https://github.com/sirupsen/logrus) - Logger
