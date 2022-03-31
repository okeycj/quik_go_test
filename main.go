package main

import (
	"github.com/gin-gonic/gin"
	"github.com/okeycj/quik_go_test/config"
	"github.com/okeycj/quik_go_test/controller"
	"github.com/okeycj/quik_go_test/middleware"
	"github.com/okeycj/quik_go_test/repository"
	"github.com/okeycj/quik_go_test/services"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                    = config.SetupDatabaseConnection()
	userRepository   repository.UserRepository   = repository.NewUserRepository(db)
	walletRepository repository.WalletRepository = repository.NewWalletRepository(db)
	jwtService       services.JWTService         = services.NewJWTService()
	authService      services.AuthService        = services.NewAuthService(userRepository)
	walletService    services.WalletService      = services.NewWalletService(walletRepository)
	authController   controller.AuthController   = controller.NewAuthController(authService, jwtService, walletService)
	walletController controller.WalletController = controller.NewWalletController(walletService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/v1/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	walletRoutes := r.Group("api/v1/wallet", middleware.AuthorizeJWT(jwtService))
	{
		walletRoutes.GET("/:id/balance", walletController.GetWalletBalance)
		walletRoutes.POST("/:id/credit", walletController.CreditWallet)
		walletRoutes.POST("/:id/debit", walletController.DebitWallet)
	}

	r.Run()
}
