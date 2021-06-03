# gqlgenc-omitempty-example

run [dgraph](https://dgraph.io) in [docker](https://docs.docker.com/engine/install/)
```
$ docker-compose up -d
```

set the schema with [dgraph-admin](https://github.com/blukai/dgraph-admin)
```
$ dgraph-admin update-schema schema.graphql
```

generate modelds and client with [gqlgenc](https://github.com/Yamashou/gqlgenc)
```
$ gqlgenc
```

run the example
```
$ go run main.go
```

and fail
```
{"networkErrors":null,"graphqlErrors":[{"message":"Non-nullable field 'name' (type String!) was not present in result from Dgraph.  GraphQL error propagation triggered.","path":["updatePlanet","planet",0,"name"],"locations":[{"line":10,"column":2}]}]}
```
