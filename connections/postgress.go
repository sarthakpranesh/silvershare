package connections

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostgresConnector() {
	dsn := "host=ec2-35-170-123-64.compute-1.amazonaws.com user=xntzdkoapaohjo password=cbf8e95011a9ee4f80d754af1493680e68bae1cf46788cbcdff9fccd22bf539a dbname=d5o4imonk0js4f port=5432"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Database not connected")
		os.Exit(1)
	}
	DB.AutoMigrate()
	fmt.Println("Database connected")
}
