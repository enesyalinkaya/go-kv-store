# In-memory Key-Value Store REST API 
A REST API for In-memory Key-Value Store application with Go

It is an In-memory Key-Value Store REST API with Go using standart Go packages. Data Loads from the file when initialization step is executed.

Default `DIR_NAME` is `tmp`, `FILE_NAME` is `data`. You can run the application with custom `DIR_NAME` and `FILE_NAME` variables.

## Installation & Run
```bash
# Download this project
go get github.com/enesyalinkaya/go-kv-store
```
Run application 
```bash
# Build
cd go-kv-store/cmd
go build .
./cmd
# API Endpoint: localhost:8080
```

Run with Docker
```bash
# Build
docker build -t kv-store:latest .
docker run -p 8080:8080 kv-store:latest
# API Endpoint: localhost:8080
```

## Run the tests
```bash
# Run the test
go test -v ./...
```

## Settings
ENV Name | Default Value | Description
--- | --- | ---
PORT | 8080 | Server Port
ADDR |  | Network Adress
READ_TIMEOUT | 60 | Server Read Timeout (seconds)
WRITE_TIMEOUT | 60 | Server Write Timeout (seconds)
IDLE_TIMEOUT | 10 | Server Idle Timeout (seconds)
DIR_NAME | tmp | Dir Name
FILE_NAME | data.txt | Data File Name
AUTO_SAVE_INTERVAL | 5 | Auto Save Interval (Minutes)

# Project Layout
```
go-kv-store
├─ Dockerfile
├─ Readme.md
├─ cmd                                  main applications of the project
│  └─ main.go
├─ controller                           Controller for our application
│  ├─ controller.go
│  └─ controller_test.go
├─ docs                                 API Documentation 
│  └─ kv-store.postman_collection.json 
├─ models                               Models for our application
│  ├─ store.go 
│  └─ store_test.go
├─ pkg
│  ├─ memoryDB                          In-memory Key-Value Store
│  │  └─ memory.go 
│  └─ settings                          Settings helper
│     └─ settings.go 
├─ routers                              Settings the routers
│  └─ routers.go
├─ services                             Services for our application
│  ├─ storeService.go 
│  └─ storeService_test.go
├─ go.mod
└─ todo
```

# REST API
## You can use https://go-kv-store.herokuapp.com as the base URL for tests.
* `PUT /kv/:key`: Set Value
```bash
    curl --request PUT 'http://localhost:8080/kv/testkey' \
    --header 'Content-Type: application/json' \
    --data-raw '{"value": "testvalue"}'
```
```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 21 Oct 2021 14:57:30 GMT
Content-Length: 38

{"key":"testkey","value":"testvalue"}
```

* `GET /kv/:key`: Get Value
```bash
    curl --request GET 'http://localhost:8080/kv/testkey'
```
```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 21 Oct 2021 14:58:55 GMT
Content-Length: 38

{"key":"testkey","value":"testvalue"}
```

* `POST /flush`: Flush Memory
```bash
    curl --request POST 'http://localhost:8080/flush'
```
```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 21 Oct 2021 14:59:05 GMT
Content-Length: 0
```
* `GET /healthcheck`: Health-Check
```bash
    curl --request GET 'http://localhost:8080/healthcheck'
```
```http
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 21 Oct 2021 14:59:10 GMT
Content-Length: 15

{"status":"OK"}
```