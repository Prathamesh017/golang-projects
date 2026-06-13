package services

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"restaurent-managment/types"
	"time"
)

type AuthService struct{
	db *pgx.Conn
	userService UserService
}

func NewAuthService(db *pgx.Conn,userService UserService) *AuthService{
	return &AuthService{
		db:db,
		userService: userService,
	}
}




func (service *AuthService) SignUp(req types.SignupRequest) (types.SignupResponse, error) {
	checkUser,err:=service.isUserExists(req.Email);
	if err!=nil{
		return types.SignupResponse{},err
	}
	if checkUser{
		return types.SignupResponse{} ,errors.New("User Already Exists")
	}

	passwordHash,err:=hashPassword(req.Password)

	if err!=nil{
		return types.SignupResponse{},errors.New("Error in Hashing Password")
	}

	userId,err:=service.userService.CreateUser(req.Name,req.Email,passwordHash)

	if err !=nil{
		return types.SignupResponse{},err
	}

	token,err:=generateToken(userId,req.Email);

	if err !=nil{
		return types.SignupResponse{},err
	}

	return types.SignupResponse{
		ID:int(userId),
		Token:token,
	},nil
}


//Check if user already exists 


func (service *AuthService) isUserExists(email string) (bool,error) {
	var exists bool;
	err:= service.db.QueryRow(context.Background(),`
			SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE email = $1
			)`,	email).Scan(&exists)
	if err!=nil{
		return false ,err
	}
	return exists,nil
}


func (service *AuthService) Login(email string ,password string) (types.SignupResponse,error){

	user,err:=service.userService.GetUserByEmail(email);

	if err!=nil{
		return types.SignupResponse{},err
	}

	err = VerifyPassword(
	user.PasswordHash,
	password,
)

if err != nil {
	return types.SignupResponse{},
		errors.New("invalid credentials")
}

	token,err:=generateToken(user.ID,email)

	if err!=nil{
		return types.SignupResponse{},err
	}

	return types.SignupResponse{
		ID:int(user.ID),
		Token:token,
	},nil

}


func hashPassword(password string) (string,error){
bytes, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}






func generateToken(
	userID int64,
	email string,
) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"email": email,
		"exp": time.Now().
			Add(24 * time.Hour).
			Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte("super-secret-key"),
	)
}

func VerifyPassword(
	hashedPassword string,
	password string,
) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}