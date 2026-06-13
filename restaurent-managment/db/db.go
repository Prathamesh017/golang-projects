package db

import (
	"fmt"
	"context"
	"github.com/jackc/pgx/v5"
)

//We need a better way than this , as we need the db instance in other files
/*
func ConnectDB() {

	fmt.Println("Starting Connection With Postgres DB")
	connectStr:="postgres://postgres:postgres@localhost:5433/restaurant_db?sslmode=disable"
	connection, err := pgx.Connect(context.Background(), connectStr)

	if err !=nil{
		fmt.Println("Error Connecting With DB")
		panic(err)
	}

	//have to close the connection when the application is stopped
	defer connection.Close(context.Background())

	fmt.Println("Connected to Postgres!")

}
*/

func ConnectDB() (*pgx.Conn ,error){
	connectStr:="postgres://postgres:postgres@localhost:5433/restaurant_db?sslmode=disable"
	connection, err := pgx.Connect(context.Background(), connectStr)

	if err !=nil{
		fmt.Println("Error Connecting With DB")
		return nil,err
	}

	return connection,nil
}