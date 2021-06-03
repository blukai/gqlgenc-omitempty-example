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

and the response is fine
```
[]*gen.PlanetFragment{
  &gen.PlanetFragment{
    ID: "0x2e",
    Name: "Jupiter",
    Moons: []*struct { ID string "json:\"id\" graphql:\"id\""; Name string "json:\"name\" graphql:\"name\"" }{},
  },
  &gen.PlanetFragment{
    ID: "0x2f",
    Name: "Earth",
    Moons: []*struct { ID string "json:\"id\" graphql:\"id\""; Name string "json:\"name\" graphql:\"name\"" }{
      &struct { ID string "json:\"id\" graphql:\"id\""; Name string "json:\"name\" graphql:\"name\"" }{
        ID: "0x30",
        Name: "Moon",
      },
    },
  },
}
```
