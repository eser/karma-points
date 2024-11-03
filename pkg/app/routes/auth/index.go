package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/eser/go-service/pkg/app/config"
	"github.com/eser/go-service/pkg/bliss/httpfx"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func IndexRoutes(routes httpfx.Router, appConfig *config.AppConfig) {
	routes.
		Route("POST /register", Register(appConfig)).
		HasSummary("Register a new user").
		HasDescription("Registers a new user with a username and password").
		HasResponse(http.StatusOK)

	routes.
		Route("POST /login", Login(appConfig)).
		HasSummary("Login a user").
		HasDescription("Logs in a user and returns a JWT token").
		HasResponse(http.StatusOK)
}

func Register(appConfig *config.AppConfig) httpfx.Handler {
	return func(ctx *httpfx.Context) httpfx.Result {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := ctx.Request; err != nil {
			return ctx.Results.BadRequest()
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return ctx.Results.BadRequest()
		}

		log.Printf("Hashed password: %s", string(hashedPassword))

		// TODO(@sameterkanboz, @eser): implement the database operations
		// err = appConfig.DB.CreateUser(req.Username, string(hashedPassword))
		// if err != nil {
		//     return ctx.Results.InternalServerError()
		// }

		return ctx.Results.PlainText("User registered successfully")
	}
}

func Login(appConfig *config.AppConfig) httpfx.Handler {
	return func(ctx *httpfx.Context) httpfx.Result {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// if err := ctx.Request; err != nil {
		//	return ctx.Results.BadRequest()
		//}

		// Retrieve the user from the database (pseudo-code)
		// user, err := appConfig.DB.GetUserByUsername(req.Username)
		// if err != nil {
		//     return ctx.Results.Unauthorized("Invalid username or password")
		// }

		// Compare the hashed password
		// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		// if err != nil {
		//     return ctx.Results.Unauthorized("Invalid username or password")
		// }
		const TokenExpiration = 72 * time.Hour
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": req.Username,
			"exp":      time.Now().Add(TokenExpiration).Unix(),
		})

		tokenString, err := token.SignedString([]byte(appConfig.JWTSecret))
		if err != nil {
			return ctx.Results.BadRequest()
		}

		return ctx.Results.Json("token: " + tokenString)
	}
}
