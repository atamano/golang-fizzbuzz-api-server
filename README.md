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

## Built With

- [Gin](https://github.com/gin-gonic/gin) - Http framework
- [Ozzo validation](https://github.com/go-ozzo/ozzo-validation) - Requests validator
- [Go pg](https://github.com/go-pg/pg) - Golang ORM with focus on PostgreSQL features and performance
- [Envconfig](https://github.com/kelseyhightower/envconfig) - Managing configuration data from environment variables
- [Logrus](https://github.com/sirupsen/logrus) - Logger
