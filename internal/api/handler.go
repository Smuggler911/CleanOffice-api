package api

import (
	"CleanOffice/internal/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User
	Offer
	Application
	Review
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(middleware.Cors()))
	router.Static("/picture/", "/app/pkg/files")
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.registerNewUser)
		auth.POST("/login", h.login)
		auth.GET("/validate", middleware.RequireAuth, h.validate)
		auth.POST("/logout", middleware.RequireAuth, h.logout)
		auth.POST("/update-token", middleware.RequireAuth, middleware.RequireRefresh, h.updateToken)
		auth.PUT("/upload-profile-picture", middleware.RequireAuth, h.uploadProfilePicture)
		auth.PUT("/update-user", middleware.RequireAuth, h.updateUser)

	}

	user := router.Group("/user")
	{
		user.GET("", middleware.RequireAuth, h.getAllUsers)
		user.PUT("/ban-user/:user_id", middleware.RequireAuth, h.banUser)
		user.PUT("/unban-user/:user_id", middleware.RequireAuth, h.unbanUser)
		user.GET("/banned", middleware.RequireAuth, h.getBannedUsers)
	}

	offer := router.Group("/offer")
	{
		offer.GET("", h.getOffers)
		offer.GET("/:offer_id", h.getOfferById)
		offer.POST("/create", middleware.RequireAuth, h.createOffer)
		offer.PUT("/update/:offer_id", middleware.RequireAuth, h.updateOffer)
		offer.DELETE("/delete/:offer_id", middleware.RequireAuth, h.deleteOffer)
		offer.POST("/create-costrange/:offer_id", middleware.RequireAuth, h.createCostRange)
		offer.PUT("/update-costrange/:offer_id/:cost_id", middleware.RequireAuth, h.editCostRange)
		offer.DELETE("/delete-costrange/:offer_id/:cost_id", middleware.RequireAuth, h.deleteCostRange)
		offer.GET("/costrange/:offer_id", h.getCostRange)

	}

	application := router.Group("/application")
	{
		application.POST("/create/:offer_id/:cost_id", middleware.RequireAuth, h.createApplication)
		application.PUT("/send-verification/:application_Id", middleware.RequireAuth, h.sendVerificationCode)
		application.PUT("/verify/:application_Id", middleware.RequireAuth, h.verifyCode)
		application.GET("/booked", middleware.RequireAuth, h.getAllBookedApplication)
		application.GET("/approved", middleware.RequireAuth, h.getApprovedApplications)
		application.GET("/declined", middleware.RequireAuth, h.getDeclinedApplications)
		application.PUT("/approve/:application_id", middleware.RequireAuth, h.approveApplication)
		application.PUT("/decline/:application_id", middleware.RequireAuth, h.declineApplication)
		application.PUT("/done/:application_id", middleware.RequireAuth, h.doneApplication)

	}

	review := router.Group("/review")
	{
		review.POST("/create/:application_id", middleware.RequireAuth, h.createReview)
		review.PUT("/approve/:review_id", middleware.RequireAuth, h.approveReview)
		review.PUT("/decline/:review_id", middleware.RequireAuth, h.declineReview)
		review.GET("", h.getReviews)

	}

	return router

}
