package postgres

import (
	"CleanOffice/config"
	"CleanOffice/internal/lib/jwt"
	"CleanOffice/internal/models"
	"CleanOffice/internal/repository"
	"CleanOffice/pkg/client"
	"CleanOffice/pkg/convert"
	"CleanOffice/pkg/generatives"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
)

//USER

func RegisterUser(c *gin.Context) {
	var registerBody *models.User
	err := c.ShouldBindJSON(&registerBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error":  true,
			"result": "не введены данные",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(registerBody.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Ошибка хеширования",
		})
		return
	}

	if registerBody.Username == "admin" && registerBody.Password == "admin" {
		registerBody.IsAdmin = true
	}

	newUser := models.User{
		FullName: registerBody.FullName,
		Password: string(hash),
		Username: registerBody.Username,
		IsAdmin:  registerBody.IsAdmin,
	}

	result := repository.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error":  result.Error,
			"result": "ошибка при создании пользователя",
		})
		return
	}
	c.JSON(200, gin.H{"result": "ok"})
}
func Login(c *gin.Context) {
	var loginData models.User
	err := c.ShouldBindJSON(&loginData)
	if err != nil {
		c.JSON(400, gin.H{
			"error":  true,
			"result": "Не введен логин или пароль",
		})
		return
	}
	var user models.User
	isLoginValid := repository.DB.First(&user, "username = ?", loginData.Username).Error
	if errors.Is(isLoginValid, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":  true,
			"result": "Не правильно введен логин",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  true,
			"result": "Не правильно введен пароль",
		})
		return
	}
	accessToken, refreshToken, err := jwt.GenerateTokenPair(user)
	refreshdata := base64.StdEncoding.EncodeToString([]byte(refreshToken))
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", accessToken, 15*60, "/", "localhost", false, true)
	c.SetCookie("RefreshToken", refreshdata, 3600*60*24, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"accesToken":   accessToken,
		"refreshToken": refreshToken,
	})
}
func Validate(c *gin.Context) {

	user, exists := c.Get("user")
	if exists {

		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})

	}
}
func Logout(c *gin.Context) {
	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.SetCookie("RefreshToken", "", -1, "/", "", false, true)
	c.String(http.StatusOK, "Вы вышли из аккаунта")
}

func UpdateUser(c *gin.Context) {
	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	userId := usr.Id

	file, err := c.FormFile("picture")
	fmt.Println(err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "нет картинки",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	env, _ := config.LoadConfig()
	imgPath := env.ImgPath

	destinationPath := imgPath + newFileName

	fmt.Println("Destination Path:", destinationPath)

	if err := c.SaveUploadedFile(file, destinationPath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Невозможно сохранить картинку",
			"error":   err.Error(),
		})
		return
	}
	var user models.User
	repository.DB.First(&user, userId)
	repository.DB.Model(&user).Updates(models.User{
		ProfilePicture: newFileName,
		Username:       c.PostForm("username"),
		Password:       c.PostForm("password"),
	})

	c.JSON(200, gin.H{
		"status": "updated",
	})
}
func UploadUserPicture(c *gin.Context) {
	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	userId := usr.Id

	file, err := c.FormFile("picture")
	fmt.Println(err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "нет картинки",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	env, _ := config.LoadConfig()
	imgPath := env.ImgPath

	destinationPath := imgPath + newFileName

	fmt.Println("Destination Path:", destinationPath)

	if err := c.SaveUploadedFile(file, destinationPath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Невозможно сохранить картинку",
			"error":   err.Error(),
		})
		return
	}
	var user models.User
	repository.DB.First(&user, userId)
	repository.DB.Model(&user).Updates(models.User{
		ProfilePicture: newFileName,
	})

	c.JSON(200, gin.H{
		"status": "updated",
	})
}

func GetAllUsers(c *gin.Context) {
	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 1 {
		limit = 3
	}
	offset := (page - 1) * limit

	var users []models.User

	repository.DB.Limit(limit).Offset(offset).Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}

func BanUser(c *gin.Context) {
	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	userId := c.Param("user_id")
	var user models.User
	repository.DB.Where("is_banned = ?", false).First(&user, userId)
	repository.DB.Model(&user).Updates(models.User{
		IsBanned: true,
	})
	c.JSON(200, gin.H{
		"result": "banned",
	})
}
func UnbanUser(c *gin.Context) {
	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	var userbody models.User

	err := c.ShouldBindJSON(&userbody)
	if err != nil {
		return
	}

	userId := c.Param("user_id")
	var user models.User
	repository.DB.First(&user, userId)

	if user.IsBanned {
		userbody.IsBanned = false
	}
	fmt.Println(userbody.IsBanned)

	repository.DB.Model(&user).Update("is_banned", userbody.IsBanned)

	c.JSON(200, gin.H{
		"result": "unbanned",
	})
}

func GetBannedUsers(c *gin.Context) {

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	var users []models.User
	repository.DB.Where("is_banned = ?", true).Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})

}

//OFFER

func CreateOffer(c *gin.Context) {

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}

	file, err := c.FormFile("picture")
	fmt.Println(err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "нет картинки",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	env, _ := config.LoadConfig()
	imgPath := env.ImgPath

	destinationPath := imgPath + newFileName

	fmt.Println("Destination Path:", destinationPath)

	if err := c.SaveUploadedFile(file, destinationPath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Невозможно сохранить картинку",
			"error":   err.Error(),
		})
		return
	}

	newOffer := models.Offer{
		Name:        c.PostForm("name"),
		Picture:     newFileName,
		Description: c.PostForm("description"),
	}
	result := repository.DB.Create(&newOffer)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error":  result.Error,
			"result": "ошибка при создании улсуги",
		})
		return
	}
	c.JSON(200, gin.H{"result": "ok"})

}

func EditOffer(c *gin.Context) {

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	file, err := c.FormFile("picture")
	fmt.Println(err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "нет картинки",
		})
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	env, _ := config.LoadConfig()
	imgPath := env.ImgPath

	destinationPath := imgPath + newFileName

	fmt.Println("Destination Path:", destinationPath)

	if err := c.SaveUploadedFile(file, destinationPath); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Невозможно сохранить картинку",
			"error":   err.Error(),
		})
		return
	}
	offerId := c.Param("offer_id")
	var offer models.Offer

	repository.DB.First(&offer, offerId)
	repository.DB.Model(&offer).Updates(models.Offer{
		Name:        c.Param("name"),
		Picture:     newFileName,
		Description: c.Param("description"),
	})
	c.JSON(200, gin.H{
		"result": "updated",
	})

}

func DeleteOffer(c *gin.Context) {

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	offerId := c.Param("offer_id")
	var offer models.Offer
	repository.DB.First(&offer, offerId)
	var costrange models.CostRange
	repository.DB.Where("offer_id = ? ", offerId).Delete(&costrange)
	repository.DB.Model(&offer).Delete(models.Offer{}, offerId)

	c.Status(200)

}

func GetOffers(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 1 {
		limit = 3
	}
	offset := (page - 1) * limit

	var offers []models.Offer
	repository.DB.Preload("CostRange").Limit(limit).Offset(offset).Find(&offers)
	c.JSON(200, gin.H{
		"offers": offers,
	})
}

func GetOfferById(c *gin.Context) {
	offerId := c.Param("offer_id")
	var offer models.Offer
	repository.DB.Preload("CostRange").First(&offer, offerId)
	c.JSON(200, gin.H{
		"offer": offer,
	})
}

func CreateCostRange(c *gin.Context) {

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}

	offerId := c.Param("offer_id")

	var rangeBody models.CostRange
	err := c.ShouldBindJSON(&rangeBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	newRange := models.CostRange{
		FlatRange: rangeBody.FlatRange,
		CostRange: rangeBody.CostRange,
		OfferId:   uint(convert.ConvertStringUint(offerId)),
	}
	result := repository.DB.Create(&newRange)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error":  result.Error,
			"result": "ошибка при создании ценнового диапазона",
		})
		return
	}
	c.JSON(200, gin.H{"result": "ok"})

}

func EditCostRange(c *gin.Context) {

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}

	offerId := c.Param("offer_id")
	costId := c.Param("cost_id")

	var rangeBody models.CostRange
	err := c.ShouldBindJSON(&rangeBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	var costRange models.CostRange
	repository.DB.Where("offer_id = ?", offerId).First(costRange, costId)
	repository.DB.Model(&costRange).Updates(models.CostRange{
		FlatRange: rangeBody.FlatRange,
		CostRange: rangeBody.CostRange,
		OfferId:   uint(convert.ConvertStringUint(offerId)),
	})
	c.JSON(200, gin.H{"result": "updated"})

}
func DeleteCostRange(c *gin.Context) {

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}

	var costRange models.CostRange
	offerId := c.Param("offer_id")
	costId := c.Param("cost_id")

	repository.DB.Where("offer_id = ?", offerId).First(costRange, costId)
	repository.DB.Delete(&costRange)

	c.Status(200)
}

func GetCostRange(c *gin.Context) {
	offerId := c.Param("offer_id")
	var costRange []models.CostRange
	repository.DB.Where("offer_id=?", offerId).Find(&costRange)
	c.JSON(200, gin.H{
		"costRange": costRange,
	})
}

//Aplication

func CreateApplication(c *gin.Context) {

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	userId := usr.Id

	offerId := c.Param("offer_id")
	costId := c.Param("cost_id")

	var applicationBody models.Application
	err := c.ShouldBindJSON(&applicationBody)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Пожалуйста заполните заявку",
		})
		return
	}
	newApplication := models.Application{
		OfferId:     uint(convert.ConvertStringUint(offerId)),
		CostRangeId: uint(convert.ConvertStringUint(costId)),
		Time:        applicationBody.Time,
		Date:        applicationBody.Date,
		IsBooked:    true,
		UserId:      userId,
	}
	result := repository.DB.Create(&newApplication)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error":  result.Error,
			"result": "ошибка при оставлении заявки",
		})
		return
	}

	c.JSON(200, gin.H{"result": "ok", "applicationId": newApplication.Id})

}

func SendVerificationCode(c *gin.Context) {

	applicationId := c.Param("application_Id")

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	userId := usr.Id

	verificationCode := generatives.RandomNumbers()

	var application models.Application
	var body models.Application

	err := c.ShouldBindJSON(&body)

	if err != nil {
		return
	}
	if len(body.Phone) < 12 || len(body.Phone) > 13 {
		log.Println("телефон  содержит меньше 12 или большк 13 cимволов:", body.Phone)
		c.JSON(http.StatusBadRequest, &gin.H{
			"Message": "телефон  содержит меньше 12 или большк 13 cимволов",
			"Status":  false,
		})
		return
	}

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if !re.MatchString(body.Phone) {
		log.Println("Номер телефона введен не правильно:", body.Phone)
		c.JSON(http.StatusBadRequest, &gin.H{
			"Message": "Номер телефона введен не правильно",
			"Status":  false,
		})
		return
	}

	repository.DB.Where("user_id = ?", userId).First(&application, applicationId)
	repository.DB.Model(&application).Updates(models.Application{
		VerificationCode: verificationCode,
		Phone:            body.Phone,
	})

	client.SendSmsToClient(body.Phone, strconv.Itoa(verificationCode))

	fmt.Println()

	c.Status(200)
}
func VerifyCode(c *gin.Context) {

	applicationId := c.Param("application_Id")

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	userId := usr.Id

	var application models.Application
	var body models.Application

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repository.DB.Where("user_id", userId).First(&application, applicationId)

	if application.VerificationCode == body.VerificationCode {
		body.Verified = true
	}

	fmt.Println(body.Verified)

	repository.DB.Model(&application).Updates(models.Application{
		Verified: body.Verified,
	})
	c.JSON(200, gin.H{
		"status": "verified",
	})

}

func GetAllBookedApplication(c *gin.Context) {

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 1 {
		limit = 3
	}
	offset := (page - 1) * limit

	var applications []models.Application
	repository.DB.Preload("Offer").Preload("CostRange").Preload("User").Where("is_booked = ?", true).Limit(limit).Offset(offset).Find(&applications)

	c.JSON(200, gin.H{
		"applications": applications,
	})
}

func GetApprovedApplications(c *gin.Context) {
	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 1 {
		limit = 3
	}
	offset := (page - 1) * limit

	var applications []models.Application
	repository.DB.Preload("Offer").Preload("CostRange").Preload("User").Where("is_approved = ?", true).Limit(limit).Offset(offset).Find(&applications)

	c.JSON(200, gin.H{
		"applications": applications,
	})
}
func GetDeclinedApplications(c *gin.Context) {
	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 1 {
		limit = 3
	}
	offset := (page - 1) * limit

	var applications []models.Application
	repository.DB.Preload("Offer").Preload("CostRange").Preload("User").Where("is_declined = ?", true).Limit(limit).Offset(offset).Find(&applications)

	c.JSON(200, gin.H{
		"applications": applications,
	})
}

func ApproveApplication(c *gin.Context) {
	applicationId := c.Param("application_id")

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	var application models.Application
	repository.DB.First(&application, &applicationId)
	repository.DB.Model(&application).Update("is_booked", false)
	repository.DB.Model(&application).Updates(models.Application{
		IsApproved: true,
	})

}

func DeclineApplication(c *gin.Context) {
	applicationId := c.Param("application_id")

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	var application models.Application
	repository.DB.First(&application, &applicationId)
	repository.DB.Model(&application).Update("is_booked", false)
	repository.DB.Model(&application).Update("is_approved", false)
	repository.DB.Model(&application).Updates(models.Application{
		IsDeclined: true,
	})
}
func DoneApplication(c *gin.Context) {
	applicationId := c.Param("application_id")

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	var application models.Application
	repository.DB.First(&application, &applicationId)
	repository.DB.Model(&application).Update("is_booked", false)
	repository.DB.Model(&application).Updates(models.Application{
		Done: true,
	})
}

// REVIEW

func CreateReview(c *gin.Context) {

	applicationId := c.Param("application_id")

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	userId := usr.Id

	var reviewBody models.Review
	err := c.ShouldBindJSON(&reviewBody)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	newReview := models.Review{
		ApplicationId: uint(convert.ConvertStringUint(applicationId)),
		UserId:        userId,
		FirstRate:     reviewBody.FirstRate,
		SecondRate:    reviewBody.SecondRate,
		Review:        reviewBody.Review,
	}
	result := repository.DB.Create(&newReview)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error":  result.Error,
			"result": "ошибка при оставлении заявки",
		})
		return
	}

	c.JSON(200, gin.H{"result": "ok"})

}

func ApproveReview(c *gin.Context) {

	reviewId := c.Param("review_id")

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	var review *models.Review
	repository.DB.Where("is_approved = ? AND is_declined = ?", false, false).First(&review, reviewId)
	repository.DB.Model(&review).Updates(models.Review{
		IsApproved: true,
	})
	c.Status(200)

}
func DeclineReview(c *gin.Context) {

	reviewId := c.Param("review_id")

	exists, _ := c.Get("user")
	if exists == nil {
		return
	}
	usr, _ := c.MustGet("user").(models.User)
	IsAdmin := usr.IsAdmin
	if !IsAdmin {
		c.JSON(400, gin.H{
			"message": "вы не являетесь админом",
		})
		return
	}
	var review *models.Review
	repository.DB.Where("is_approved = ? AND is_declined = ?", false, false).First(&review, reviewId)
	repository.DB.Model(&review).Updates(models.Review{
		IsDeclined: true,
	})
	c.Status(200)

}

func GetReviews(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit < 1 {
		limit = 3
	}
	offset := (page - 1) * limit

	var reviews []models.Review
	repository.DB.Preload("User").Preload("Application").Where("is_declined = ? AND is_approved = ?", false, true).Limit(limit).Offset(offset).Find(&reviews)
	c.JSON(200, gin.H{
		"reviews": reviews,
	})
}
