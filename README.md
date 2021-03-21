# Noto to self: graphql-go Tutorial

https://www.howtographql.com/graphql-go/0-introduction/

## Note

### Getting started

https://www.howtographql.com/graphql-go/1-getting-started/

`validation failed: packages.Load` error occured.
To fix it, setup go module.

```sh
% cd <project dir>
% go mod init
% go get github.com/99designs/gqlgen

# This module is needed.
% go get github.com/vektah/gqlparser/v2

%gqlgen init
```

### Database

https://www.howtographql.com/graphql-go/4-database/

`server.go` needs `go-chi`.

```sh
% go get github.com/go-chi/chi
% go mod vendor
```

### Testing query and mutation.

We can test on browser UI which URL is `http://localhost:8080`.

## Additional sources


```sh
.
|-- dbmigrate   <-- DB migration
|-- dbstart     <-- Start MySQL on docker
`-- dbstop      <-- Stop MySQL on docker
```
