package models

type User struct {
	Id             uint   `gorm:"PrimaryKey" json:"id"`
	FullName       string `json:"fullname"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	IsAdmin        bool   `json:"isadmin"`
	IsBanned       bool   `json:"isbanned"`
	ProfilePicture string `json:"profilepicture"`
}
type Offer struct {
	Id          uint        `gorm:"PrimaryKey" json:"id"`
	Picture     string      `json:"picture"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CostRange   []CostRange `gorm:"many2many:offers_costrange;" json:"cost_range,omitempty"`
}

type CostRange struct {
	Id        uint   `gorm:"PrimaryKey" json:"id"`
	FlatRange string `json:"flatrange"`
	CostRange string `json:"costrange"`
	OfferId   uint   `json:"offer_id"`
	Offer     Offer  `gorm:"foreignKey:OfferId" json:"offer,omitempty"`
}

type Application struct {
	Id               uint      `gorm:"PrimaryKey" json:"id"`
	OfferId          uint      `json:"offer_id"`
	Offer            Offer     `gorm:"foreignKey:OfferId" json:"offer"`
	CostRangeId      uint      `json:"cost_range_id"`
	CostRange        CostRange `gorm:"foreignKey:CostRangeId" json:"cost_range"`
	Time             string    `json:"time"`
	Date             string    `json:"date"`
	IsBooked         bool      `json:"is_booked"`
	VerificationCode int       `json:"verification"`
	Verified         bool      `json:"verified"`
	IsApproved       bool      `json:"is_approved"`
	IsDeclined       bool      `json:"is_declined"`
	Done             bool      `json:"done"`
	UserId           uint      `json:"user_id"`
	Phone            string    `json:"phone"`
	User             User      `gorm:"foreignKey:UserId" json:"user"`
}

type Review struct {
	Id            uint        `gorm:"PrimaryKey" json:"id"`
	FirstRate     int32       `json:"first_rate"`
	SecondRate    int32       `json:"second_rate"`
	ApplicationId uint        `json:"application_id"`
	Application   Application `gorm:"foreignKey:ApplicationId" json:"application"`
	UserId        uint        `json:"user_id"`
	User          User        `gorm:"foreignKey:UserId" json:"user"`
	Review        string      `json:"review"`
	IsApproved    bool        `json:"is_approved"`
	IsDeclined    bool        `json:"is_declined"`
}

//type Portfolio struct {
//	Id            uint        `gorm:"PrimaryKey" json:"id"`
//	Pictures      []Picture   `gorm:"many2many:portfolioPictures;" json:"pictures,omitempty"`
//	Description   string      `json:"description"`
//	ReviewId      uint        `json:"review_id"`
//	Review        Review      `gorm:"foreignKey:ReviewId" json:"review"`
//	UserId        uint        `json:"user_id"`
//	User          User        `gorm:"foreignKey:UserId" json:"user"`
//	ApplicationId uint        `json:"application_id"`
//	Application   Application `gorm:"foreignKey:ApplicationId" json:"application"`
//}
//type Picture struct {
//	Id      uint   `gorm:"PrimaryKey" json:"id"`
//	Picture string `json:"picture"`
//}
