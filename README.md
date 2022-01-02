# gothirfttutorial
Simple example for thrift with go

### Run server
go run main.go server.go client.go handler.go -server true

### gen
thrift -r --gen go  tutorial.thrift
