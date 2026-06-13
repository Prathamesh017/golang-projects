package services

import (
	"context"
	"github.com/jackc/pgx/v5"
)


type UserService struct{
    db *pgx.Conn
}


func NewUserService(db *pgx.Conn) *UserService{
	return &UserService{
		db:db,
	}
}




func (service *UserService) CreateUser(name string,email string,passwordHash string) (int64,error){
  var userId int64;
  err:=service.db.QueryRow(context.Background(),
`
		INSERT INTO users (
			name,
			email,
			password_hash
		)
		VALUES (
			$1,
			$2,
			$3
		)
		RETURNING id
		`,name,
		email,
		passwordHash,).Scan(&userId)

	if err!=nil{
		return -1,err
	}
   return userId,nil
}


type UserCredentials struct {
	ID           int64
	PasswordHash string
}
func (service *UserService) GetUserByEmail(email string) (UserCredentials,error){

	   var user UserCredentials;
		err := service.db.QueryRow(
		context.Background(),
		`
		SELECT
		id,
		password_hash
		FROM users
		WHERE email = $1
		`,
		email,
	).Scan(
		&user.ID,

		&user.PasswordHash,
	)

	if err != nil {
		return UserCredentials{}, err
	}

	return user, nil
}