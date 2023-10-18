package main

import (
	"os"

	"github.com/Matheus-Campos/codepix/application/grpc"
	"github.com/Matheus-Campos/codepix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
