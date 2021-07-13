package controllers

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/sidmohanty11/auth-jwt-golang/models"
)

var MyUser models.User

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("sid")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	// user is not logged in.
	if err != nil {
		return c.JSON(fiber.Map{"error": "unauthorized", "status": "error"})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	// db query where id = iss (id)
	MyUser = models.User{
		Id:       1,
		Name:     "sid",
		Email:    "sid@sid.com",
		Password: "password",
	}

	return c.JSON(fiber.Map{"user": MyUser, "claims": claims})
}

func Login(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)

	if err != nil {
		return err
	}

	// dummy user, replace with db user + more security.
	MyUser = models.User{
		Id:       1,
		Name:     "sid",
		Email:    "sid@sid.com",
		Password: "password",
	}

	if user.Email != MyUser.Email || MyUser.Password != user.Password {
		return c.JSON(fiber.Map{"error": "unauthorized", "status": "error"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(MyUser.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "sid",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "sid",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "logout successful",
	})
}
