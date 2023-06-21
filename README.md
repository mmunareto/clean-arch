# clean-arch
Go Expert challenge to create an application using clean arch.

## How to run

**1.** First run the command below to start the dependencies:

```bash
$ make infra-up
```
_Note: If you want to stop infra, run the command below:_

```bash
$ make infra-down
```


**2.** After start dependencies, run the command below to execute migrations:

```bash
$ make migrate-create
```
_Note: If you want to drop tables, run the command below:_

```bash
$ make migrate-destroy
```

**3.** Finally, run the command below to start the application:
```bash
$ make run
```

## How to test
### Webserver
To test the webserver, you can use the commands below:
#### Create order operation:
```bash
$ curl -X POST --location "http://localhost:8000/order" --http1.1 \
    -H "Host: localhost:8000" \
    -H "Content-Type: application/json" \
    -d "{
          \"id\":\"web-server\",
          \"price\": 100.5,
          \"tax\": 0.5
        }"
```
#### List Orders operation:
```bash
$ curl -X GET --location "http://localhost:8000/order" --http1.1 \
    -H "Host: localhost:8000" \
    -H "Content-Type: application/json"
```
### gRPC server
To test the grpc server, you need install evans and run the commands below:
```bash
$ evans -r repl
```
##### Create order operation:
```bash
pb.OrderService@127.0.0.1:50051> call CreateOrder
// Inputs
id (TYPE_STRING) => grpc-server
price (TYPE_FLOAT) => 430
tax (TYPE_FLOAT) => 2.0
```
#### List orders operation:
```bash
call ListOrders
```

### GraphQL server
To test the graphql server, you need access the playground url http://localhost:8080/ and run the commands below:
```bash
mutation createOrder {
  createOrder(
    input: {id: "graphql-server", Price: 500.0, Tax: 2.0}
  ) {
    id
    Price
    Tax
    FinalPrice
  }
}
```
```bash
query queryOrders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```
