# Лабораторная работа №3 Серверная часть
### Запуск:
```go run main.go```  
### Компиляция proto файла:
```protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld```