go run greeter_server/main.go
node greeter_client/greeter_client.js
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld


go run main.go
node greeter_client.js