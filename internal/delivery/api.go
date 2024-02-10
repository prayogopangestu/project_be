package delivery

import (
	"fmt"
	"html/template"
	"net/http"
	"project/config/database"
	"project/internal/controller"
	"project/internal/delivery/api"
	"project/internal/repository"
	"project/internal/service"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var server = database.Server{
	DB: &gorm.DB{},
}

var (
	db *gorm.DB = server.SetupDatabaseConnection()

	jwtService service.JWTService = service.NewJWTService()

	userRepository repository.UserRepository = repository.NewUserRepository(db)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	userService    service.UserService       = service.NewUserService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// subjectFromJWT parses a JWT and extract subject from sub claim.
func subjectFromJWT(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	aToken, err := jwtService.ValidateToken(authHeader)
	claims := aToken.Claims.(jwt.MapClaims)
	subject := fmt.Sprintf("%v", claims["email"])
	if err != nil {
		return ""
	}
	return subject
}

func InitializeRoutes() {

	defer database.CloseDatabaseConnection(db)

	// middleware.NewCasbinInit(db).InitEnforcer() //import config rules
	// auth, err := middleware.NewCasbinMiddleware("config/etc/casbin/model.conf", db, subjectFromJWT)
	// fmt.Printf("%v", auth)

	// if err != nil {
	// 	fmt.Println("error casbin")
	// }
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", h1)

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.Use(CORSMiddleware())
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		// authRoutes.GET("/getToken", authController.LoginBpjs)
	}
	api.RmCodeRoutes(r, db, jwtService)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "5000" //default env
	}
	r.Run(":" + port)

}
