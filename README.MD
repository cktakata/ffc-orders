# Fast Food Chat
## _Orders Module_

Fast Food Chat is the trademark of the portal to be used as a main access to use our chat robot for attendance client's orders.

This _Orders Module_ will be used to manage orders from each client at each store, using Go as the programming language, MongoDB for data storage, and blockchain technology to ensure transaction trustworthiness.

## API

- List all Orders:
```
curl --location 'http://localhost:8000/api/getAllOrders'
```
- Get Specific Order:
```
curl --location 'http://localhost:8000/api/getOrder/66c4db8dd78fb25ac780a0a8'
```
- Add new order:
```
curl --location 'http://localhost:8000/api/addOrder' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Mary",
    "value": 25.0
}'
```
- Charge Back an order*:
```
curl --location 'http://localhost:8000/api/chargeBackOrder' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Mary",
    "value": -25.0
}'
```
*if Value is "0" (zero), then the value to charge back will be the same as the last one;
- Validate all transactions:
```
curl --location 'http://localhost:8000/api/validateOrders'
```

## Tech

Fast Food Chat - Orders Module, uses a number of open source projects to work properly:

- [GoLang] - API generation
- [MongoDb] - Database storage
- [Docker] - To quickly start and develop the application

## Installation

Simply run:

```sh
docker-compose up --build 
```


## License

MIT

