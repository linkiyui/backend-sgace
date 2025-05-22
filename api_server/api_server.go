package apiserver

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	activity_controllers "github.com/sgace/backend/activity/infra/controller"
	auth_controllers "github.com/sgace/backend/auth/infra/controller"
	auth_middleware "github.com/sgace/backend/auth/infra/middleware"
	solicitude_controllers "github.com/sgace/backend/solicitude/infra/controllers"
	user_controllers "github.com/sgace/backend/user/infra/controller"
	user_middleware "github.com/sgace/backend/user/infra/middlewares"
)

var apiServer *APIServer

type APIServer struct {
	ginInstance *gin.Engine
}

func NewApiServer() *APIServer {

	ginInstance := gin.New()

	ginInstance.Use(gin.Recovery())

	ginInstance.Use(cors.Default())

	registerRoutes(ginInstance)

	apiServer = &APIServer{
		ginInstance: ginInstance,
	}
	return apiServer
}

func (s *APIServer) Start() {
	url := fmt.Sprintf(":%d", 3000)
	if err := s.ginInstance.Run(url); err != nil {
		log.Fatal("can't start api server on", url)
	}
}

func registerRoutes(r *gin.Engine) {

	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(200, "text/html; charset=utf-8", []byte("Welcome to ATcnea's API"))
	})

	v1 := r.Group("/v1")

	user := v1.Group("/user")

	user.POST("/login", auth_controllers.Login)
	user.POST("/signup", auth_controllers.Signup)
	user.POST("/refresh_token", auth_middleware.VerifyRecoverToken, auth_controllers.RefreshToken)

	userWithAuth := v1.Group("/user").Use(auth_middleware.VerifyLoginToken)
	userWithAuth.GET("/me", user_controllers.GetMyInfo)
	userWithAuth.DELETE("/", user_controllers.DeleteMyUser)
	userWithAuth.PUT("/", user_controllers.UpdateMyUser)
	userWithAuth.GET("/progress/:id", user_controllers.GetUserProgress)

	solicitude := v1.Group("/solicitude")
	solicitude.POST("/", solicitude_controllers.CreateSolicitude)
	solicitude.GET("/", solicitude_controllers.GetAllSolicitudesByUser)
	solicitude.PUT("/:id", solicitude_controllers.UpdateSolicitudeById)
	solicitude.DELETE("/:id", solicitude_controllers.DeleteSolicitudeById)

	activity := v1.Group("/activity").Use(auth_middleware.VerifyLoginToken)
	activity.GET("/", activity_controllers.GetAllActivities)
	activity.GET("/:id", activity_controllers.GetActivityByID)

	admin := v1.Group("/admin").Use(auth_middleware.VerifyLoginToken)
	admin.Use(user_middleware.VerifyAdmin)

	admin.POST("/activity", activity_controllers.CreateActivity)
	admin.PUT("/activity/:id", activity_controllers.UpdateActivity)
	admin.DELETE("/activity/:id", activity_controllers.DeleteActivity)
	admin.POST("solicitude/accept/:id", solicitude_controllers.AcceptSolicitude)
	admin.POST("solicitude/refuse/:id", solicitude_controllers.RefuseSolicitude)
	admin.GET("/solicitude", solicitude_controllers.GetAllSolicitudes)
	admin.GET("/solicitude/:id", solicitude_controllers.GetSolicitudeByID)

}
