# Go-CQRS-Demo

## Prerequisites

- docker
- make
- go1.16+
- [K6](https://k6.io/)

## Run

1. `make up`
2. `make migrate`
3. `make all`
4. choose the service you want to run under `.out`
5. run `order-server`
6. run `writer`

## K6 Test

```shell
cd test
make up
make stress-test
```