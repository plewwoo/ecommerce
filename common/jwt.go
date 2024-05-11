package common

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// var (
// 	//viper.GetString("jwt_secret_key")
// 	// secretKey     = []byte("your-secret-key")
// 	// refreshSecret = []byte("your-refresh-secret")
// 	secretKey     = []byte(viper.GetString("jwt_secret_key"))
// 	refreshSecret = []byte(viper.GetString("jwt_refresh_secret_key"))
// )

// JWTClaims struct represents the claims we want to include in the token.
type Remark struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Usercode  string `json:"usercode"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Tel       string `json:"Tel"`
}

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Remark Remark `json:"remark"`
	jwt.StandardClaims
}

func GenerateJWTToken(userID uint, remark Remark, expiresIn time.Duration) (string, error) {
	key := []byte(viper.GetString("jwt_secret_key"))

	claims := JWTClaims{
		UserID: userID,
		Remark: remark,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiresIn).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func IsPublicPath(c *fiber.Ctx) bool {
	publicPaths := viper.GetStringSlice("public_path")
	log.Println("full_path:", c.Path())
	path := getLastPathComponent(c.Path())
	log.Println("path:", path)
	log.Println("publicPaths:", publicPaths)
	return StringExistsInList(path, publicPaths)
}

func AuthenticationMiddleware(c *fiber.Ctx) error {
	if IsPublicPath(c) {
		return c.Next()
	}

	// Extract token from the Authorization header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(http.StatusUnauthorized).SendString("No token provided")
	}

	// Check Blacklist
	// if !IsJwtValid(authHeader) {
	// 	return c.Status(http.StatusUnauthorized).SendString("token blacklist")
	// }

	tokenString := authHeader[len("Bearer "):]

	// Parse the token
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("jwt_secret_key")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(http.StatusUnauthorized).SendString("Invalid token signature")
		}
		return c.Status(http.StatusBadRequest).SendString("Bad request")
	}

	if !token.Valid {
		return c.Status(http.StatusUnauthorized).SendString("Invalid token")
	}

	// Token is valid, do something with the claims
	c.Locals("user", claims)
	return c.Next()
}

// func refreshTokenHandler(c *fiber.Ctx) error {
// 	// Extract refresh token from the request
// 	refreshToken := c.FormValue("refresh_token")
// 	if refreshToken == "" {
// 		return c.Status(http.StatusBadRequest).SendString("Refresh token not provided")
// 	}

// 	// Parse the refresh token
// 	claims := &JWTClaims{}
// 	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
// 		return refreshSecret, nil
// 	})

// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			return c.Status(http.StatusUnauthorized).SendString("Invalid refresh token signature")
// 		}
// 		return c.Status(http.StatusBadRequest).SendString("Bad request")
// 	}

// 	if !token.Valid {
// 		return c.Status(http.StatusUnauthorized).SendString("Invalid refresh token")
// 	}

// 	// Refresh token is valid, generate a new access token
// 	newAccessToken, err := GenerateJWTToken(claims.UserID, secretKey, time.Hour*24)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(fiber.Map{"access_token": newAccessToken})
// }

// func FiberLogin(app *fiber.App) {
// 	app.Post("/login", func(c *fiber.Ctx) error {
// 		fmt.Println("mid", viper.GetString("jwt_secret_key"))
// 		fmt.Println("mid", viper.GetString("jwt_refresh_secret_key"))

// 		var user orm.User
// 		if err := c.BodyParser(&user); err != nil {
// 			return err
// 		}

// 		if authenticatedUser, ok := authenticate(user.Username, user.Password); ok {
// 			accessToken, err := GenerateJWTToken(authenticatedUser.ID, []byte(viper.GetString("jwt_secret_key")), time.Hour*24)
// 			if err != nil {
// 				return err
// 			}

// 			refreshToken, err := GenerateJWTToken(authenticatedUser.ID, []byte(viper.GetString("jwt_refresh_secret_key")), time.Hour*24*30) // 30 days
// 			if err != nil {
// 				return err
// 			}

// 			return c.JSON(fiber.Map{"token": accessToken, "refresh_token": refreshToken})
// 		}

// 		return c.Status(http.StatusUnauthorized).SendString("Invalid credentials")
// 	})
// }

// func FiberRefreshToken(c *fiber.Ctx) error {
// 	refreshToken := c.FormValue("refresh_token")
// 	if refreshToken == "" {
// 		return c.Status(http.StatusBadRequest).SendString("Refresh token not provided")
// 	}

// 	claims := &JWTClaims{}
// 	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
// 		return viper.GetString("jwt_refresh_secret_key"), nil
// 	})

// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			return c.Status(http.StatusUnauthorized).SendString("Invalid refresh token signature")
// 		}
// 		return c.Status(http.StatusBadRequest).SendString("Bad request")
// 	}

// 	if !token.Valid {
// 		return c.Status(http.StatusUnauthorized).SendString("Invalid refresh token")
// 	}

// 	// Refresh token is valid, generate a new access token
// 	newAccessToken, err := GenerateJWTToken(claims.UserID, []byte(viper.GetString("jwt_secret_key")), time.Hour*24)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(fiber.Map{"token": newAccessToken})
// }

// func FiberTestProtection(app *fiber.App) {
// 	app.Get("/test-protection", func(c *fiber.Ctx) error {
// 		userClaims := c.Locals("user").(*JWTClaims)
// 		return c.JSON(fiber.Map{"user_id": userClaims.UserID})
// 	})
// }

func GetSessionUserID(c *fiber.Ctx) uint {
	userClaims := c.Locals("user").(*JWTClaims)
	return userClaims.UserID //
}

func GetSession(c *fiber.Ctx) *JWTClaims {
	userClaims := c.Locals("user").(*JWTClaims)
	return userClaims
}

// func IsJwtValid(token string) bool {
// 	var bl orm.JwtBlacklist
// 	result := Database.Model(orm.JwtBlacklist{}).Where("md5 = ?", MD5(token)).First(&bl)
// 	return result.RowsAffected == 0
// }

// func authenticate(username, password string) (orm.User, bool) {
// 	var user orm.User
// 	result := Database.Where("username = ? AND password = ?", username, HashPassword(password)).First(&user)
// 	if result.Error != nil {
// 		return orm.User{}, false
// 	}

// 	return user, true
// }

// func authenticate(username, password string) (orm.User, bool) {
// 	// Replace this with your actual authentication logic
// 	// For simplicity, return a predefined user if credentials are valid
// 	if username == "user" && password == "password" {
// 		user := new(orm.User)
// 		user.ID = "1234"
// 		user.Username = "Admin"
// 		return *user, true
// 	}
// 	return orm.User{}, false
// }

// func main() {
// 	app := fiber.New()

// 	// Use authentication middleware for all routes except excluded ones
// 	app.Use(authenticationMiddleware)

// 	// Example login endpoint
// 	app.Post("/login", func(c *fiber.Ctx) error {
// 		var user User
// 		if err := c.BodyParser(&user); err != nil {
// 			return err
// 		}

// 		if authenticatedUser, ok := authenticate(user.Username, user.Password); ok {
// 			accessToken, err := generateToken(authenticatedUser.ID, authenticatedUser.Username, secretKey, time.Hour*24)
// 			if err != nil {
// 				return err
// 			}

// 			refreshToken, err := generateToken(authenticatedUser.ID, authenticatedUser.Username, refreshSecret, time.Hour*24*30) // 30 days
// 			if err != nil {
// 				return err
// 			}

// 			return c.JSON(fiber.Map{"access_token": accessToken, "refresh_token": refreshToken})
// 		}

// 		return c.Status(http.StatusUnauthorized).SendString("Invalid credentials")
// 	})

// 	// Example protected endpoint
// 	app.Get("/protected", func(c *fiber.Ctx) error {
// 		// Access the user from the authentication middleware
// 		userClaims := c.Locals("user").(*JWTClaims)
// 		return c.JSON(fiber.Map{"user_id": userClaims.UserID, "username": userClaims.Username})
// 	})

// 	// Refresh token endpoint
// 	app.Post("/refresh", refreshTokenHandler)

// 	app.Listen(":3000")
// }
