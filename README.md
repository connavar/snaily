# snaily

Note :: The Go program proxies the Angular app.

## Installing

1. Start the database
```bash
    ./start-pg.sh
```
2. Run setup.sql against postgres::snaily@localhost:5000/madast

:D  Will try automate this sometime. 

## Running the application

1. Start the angular application:
```bash
    cd snaily-web
    npm i
    ng serve
```

2. Start the go app
```bash
    go run application.go
```

IMPORTANT: Open the app in your browser at :8080 NOT :4200!

