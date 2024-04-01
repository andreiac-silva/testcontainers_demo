# Testcontainers: Taking your integration tests to the next level
This repository contains a simple example of how to use Testcontainers in your integration tests.

### Most important files
- Take a look at the [postgres.go](https://github.com/andreiac-silva/testcontainers_demo/blob/main/test/integration/postgres.go) file to see how to use Testcontainers with a PostgreSQL container.
- The [suite_test.go](https://github.com/andreiac-silva/testcontainers_demo/blob/main/domain/user/suite_test.go) gets the existing PostgreSQL container and loads the migrations to perform the tests.
- Finally, the [service_test.go](https://github.com/andreiac-silva/testcontainers_demo/blob/main/domain/user/service_test.go) is the Integration Test.

### How to run the tests
- You can start the tests through your favorite IDE or
- Run the following command in the terminal:
```go test ./...```    