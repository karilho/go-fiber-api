package model

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"github.com/karilho/go-fiber-api/src/configuration/rest_errors"
	"os"
	"time"
)

var (
	KEY = os.Getenv("JWT_SECRET_KEY")
)

func (ud *userDomainStruct) GenerateToken() (string, *rest_errors.RestErr) {
	secret := os.Getenv(KEY)

	// With claims, you choose which information to include in the JWT token from the user domain
	claims := jwt.MapClaims{
		"id":    ud.GetId(),
		"email": ud.GetEmail(),
		"name":  ud.GetName(),
		"age":   ud.GetAge(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_errors.NewInternalServerError(fmt.Sprintf("Error on generate token: %s", err.Error()))

	}

	return tokenString, nil
}

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_errors.RestErr) {
	secret := os.Getenv(KEY)

	//The func is a "callback func" that received the token and return an interface and an error
	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {

		//Check if the token is valid
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return []byte(secret), nil
		}

		return nil, rest_errors.NewBadRequestError("Invalid token")
	})

	if err != nil {
		return nil, rest_errors.NewUnauthorizedError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return nil, rest_errors.NewUnauthorizedError("Invalid token")
	}

	return &userDomainStruct{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   claims["age"].(int8),
	}, nil
}

func VerifyTokenMiddleware(ctx *fiber.Ctx) error {
	secret := os.Getenv(KEY)
	tokenValue := string(ctx.Request().Header.Peek("Authorization"))

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_errors.NewBadRequestError("Invalid token")
	})
	if err != nil {
		errRest := rest_errors.NewUnauthorizedError("No Header Provided")
		ctx.SendString(errRest.Message)
		return errRest
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logger.Error("Invalid token", err)
		ctx.Status(fiber.StatusUnauthorized).JSON(rest_errors.NewUnauthorizedError("Invalid token"))
		return nil
	}

	userDomain := &userDomainStruct{
		id:    claims["id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", userDomain))

	return ctx.Next()
}
