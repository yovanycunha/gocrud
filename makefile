init:
	go mod init go-crud
	go get github.com/githubnemo/CompileDaemon
	go install github.com/githubnemo/CompileDaemon
	go get github.com/joho/godotenv
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres

run:
	CompileDaemon -directory="./cmd/go-crud" -command="./cmd/go-crud/go-crud"