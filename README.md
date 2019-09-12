# go-multitenancy

Basic multitenancy example using go with Postgres.

## Run app

clone repository and exec command: 

```bash
$ docker-compose -f docker/docker-compose.yml up -d --build
```

## running example

Creating todo for `tenant1`
```bash
$ curl --request POST \
    --url http://localhost:8080/go-multitenancy/tenant1/todos \
    --header 'content-type: application/json' \
    --data '{
  	"description": "Todo 1"
  }'
$ 
```

Retrieving todos for `tenant1`
```bash
$ curl --request GET \
  --url http://localhost:8080/go-multitenancy/tenant1/todos
$ 
```

Retrieving todos for `tenant2`
```bash
$ curl --request GET \
  --url http://localhost:8080/go-multitenancy/tenant2/todos
$ 
```
