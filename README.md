## Go PG wrapper

Basic configuration to use [go-pg v10](github.com/go-pg/pg)

Features:
* Supporting debug
* Cloud SQL enabled (using GCP Cloud SQL proxy)

### Example

Connecting to database using hostname:
```go
cfg := pgdb.Config{
    Hostname: "localhost",
    Username: "postgres",
    Password: "",
    Database: "postgres",
    Port:     "", // optional
    Debug:    false,
}

db, err := pgdb.Init(cfg)
if err != nil {
    log.Fatalf("Unable to connect new database, err: %s", err.Error())

    return
}
```

Connecting inside GCP (appengine or VM):
```go

cfg := pgdb.Config{
    Hostname: "project:europe-west2:db-server",
    Username: "postgres", 
    Port:     "", // optional
    Database: "postgres",
    Debug:    false,
}

db, err := pgdb.Init(cfg)
if err != nil {
    log.Fatalf("Unable to connect new database, err: %s", err.Error())

    return
}
```

To use in connection with appengine, Cloud SQL API need to be enabled