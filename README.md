# paymentgw
This is a hypothetical gateway which implements hexagonal architecture and processes a payment transaction and responds randomly. Sometime it can fail and respond with an error.

## How to build

* To build project
```sh
./build.sh
```

* To build docker image
```sh
docker build . -t paymentgw
```

## How to run

* from docker image
```sh
docker run -p 8287:8287 --name paymentgwtest --rm paymentgw
```

* from the built project
```sh
./bin/paymentgwd-darwin-amd4
```

## How to consume

* Successful transaction
```sh
curl -XPOST -d '{"transaction_id":"1","company_key":"123", "amount":50000.34}' http://localhost:8287/payments

{"data":{"confirmation_id":"85c10e9f-3fab-4edf-a22d-06e2667afedd","transaction_id":"1","company_key":"123","amount":50000.34,"success":true,"created_at":"2019-12-29T13:47:33.14272-05:00"}}
```

* With a validation error
```sh
curl -XPOST -d '{"transaction_id":"1","company_key":"123", "amount":-50000.34}' http://localhost:8287/payments

{"data":{"transaction_id":"1","company_key":"123","amount":-50000.34,"success":false},"errors":["payment amount cannot be less than zero"]}
```

* With an unexpected error
```sh
curl -XPOST -d '{"transaction_id":"1","company_key":"123", "amount":50000.34}' http://localhost:8287/payments

{"data":{"transaction_id":"1","company_key":"123","amount":50000.34,"success":false},"errors":["something went wrong"]}
```

