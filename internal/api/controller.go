package api

import (
	"CleanOffice/internal/lib/jwt"
	"CleanOffice/internal/models"
	"CleanOffice/internal/repository/postgres"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User interface {
	registerNewUser(c *gin.Context)
	login(c *gin.Context)
	validate(c *gin.Context)
	logout(c *gin.Context)
	updateToken(c *gin.Context)
	uploadProfilePicture(c *gin.Context)
	getAllUsers(c *gin.Context)
	banUser(c *gin.Context)
	unbanUser(c *gin.Context)
	getBannedUsers(c *gin.Context)
	updateUser(c *gin.Context)
}
type Offer interface {
	createOffer(c *gin.Context)
	updateOffer(c *gin.Context)
	deleteOffer(c *gin.Context)
	getOffers(c *gin.Context)
	getOfferById(c *gin.Context)
	createCostRange(c *gin.Context)
	editCostRange(c *gin.Context)
	deleteCostRange(c *gin.Context)
	getCostRange(c *gin.Context)
}
type Application interface {
	createApplication(c *gin.Context)
	sendVerificationCode(c *gin.Context)
	verifyCode(c *gin.Context)
	getAllBookedApplication(c *gin.Context)
	getApprovedApplications(c *gin.Context)
	getDeclinedApplications(c *gin.Context)
	approveApplication(c *gin.Context)
	declineApplication(c *gin.Context)
	doneApplication(c *gin.Context)
}
type Review interface {
	createReview(c *gin.Context)
	approveReview(c *gin.Context)
	declineReview(c *gin.Context)
	getReviews(c *gin.Context)
}

func (h *Handler) registerNewUser(c *gin.Context) {
	postgres.RegisterUser(c)
}
func (h *Handler) login(c *gin.Context) {
	postgres.Login(c)
}
func (h *Handler) validate(c *gin.Context) {
	postgres.Validate(c)
}
func (h *Handler) logout(c *gin.Context) {
	postgres.Logout(c)
}

func (h *Handler) updateToken(c *gin.Context) {
	user, _ := c.MustGet("user").(models.User)

	accessToken, refreshToken, err := jwt.GenerateTokenPair(user)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Failed to generate token",
		})
	}
	refreshdata := base64.StdEncoding.EncodeToString([]byte(refreshToken))
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", accessToken, 15*60, "/", "localhost", false, true)
	c.SetCookie("RefreshToken", refreshdata, 3600*60*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"accesToken":   accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) uploadProfilePicture(c *gin.Context) {
	postgres.UploadUserPicture(c)
}

func (h *Handler) getAllUsers(c *gin.Context) {
	postgres.GetAllUsers(c)
}

func (h *Handler) banUser(c *gin.Context) {
	postgres.BanUser(c)
}
func (h *Handler) unbanUser(c *gin.Context) {
	postgres.UnbanUser(c)
}
func (h *Handler) getBannedUsers(c *gin.Context) {
	postgres.GetBannedUsers(c)
}
func (h *Handler) updateUser(c *gin.Context) {
	postgres.UpdateUser(c)
}

//OFFER

func (h *Handler) createOffer(c *gin.Context) {
	postgres.CreateOffer(c)
}
func (h *Handler) updateOffer(c *gin.Context) {
	postgres.EditOffer(c)
}
func (h *Handler) deleteOffer(c *gin.Context) {
	postgres.DeleteOffer(c)
}
func (h *Handler) getOffers(c *gin.Context) {
	postgres.GetOffers(c)
}
func (h *Handler) getOfferById(c *gin.Context) {
	postgres.GetOfferById(c)
}
func (h *Handler) createCostRange(c *gin.Context) {
	postgres.CreateCostRange(c)
}
func (h *Handler) editCostRange(c *gin.Context) {
	postgres.EditCostRange(c)
}
func (h *Handler) deleteCostRange(c *gin.Context) {
	postgres.DeleteCostRange(c)
}
func (h *Handler) getCostRange(c *gin.Context) {
	postgres.GetCostRange(c)
}

//APPLICATION

func (h *Handler) createApplication(c *gin.Context) {
	postgres.CreateApplication(c)
}
func (h *Handler) sendVerificationCode(c *gin.Context) {
	postgres.SendVerificationCode(c)
}
func (h *Handler) verifyCode(c *gin.Context) {
	postgres.VerifyCode(c)
}
func (h *Handler) getAllBookedApplication(c *gin.Context) {
	postgres.GetAllBookedApplication(c)
}
func (h *Handler) getApprovedApplications(c *gin.Context) {
	postgres.GetApprovedApplications(c)
}
func (h *Handler) getDeclinedApplications(c *gin.Context) {
	postgres.GetDeclinedApplications(c)
}
func (h *Handler) approveApplication(c *gin.Context) {
	postgres.ApproveApplication(c)
}
func (h *Handler) declineApplication(c *gin.Context) {
	postgres.DeclineApplication(c)
}
func (h *Handler) doneApplication(c *gin.Context) {
	postgres.DoneApplication(c)
}

//REVIEW

func (h *Handler) createReview(c *gin.Context) {
	postgres.CreateReview(c)
}
func (h *Handler) approveReview(c *gin.Context) {
	postgres.ApproveReview(c)
}
func (h *Handler) declineReview(c *gin.Context) {
	postgres.DeclineReview(c)
}
func (h *Handler) getReviews(c *gin.Context) {
	postgres.GetReviews(c)
}
