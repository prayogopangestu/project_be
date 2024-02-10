package dto

type UserUpdateDTO struct {
	ID         uint   `json:"id" form:"id" binding:"required"`
	Nickname   string ` json:"nickname" `
	First_name string ` json:"first_name"  `
	Last_name  string ` json:"last_name"  `
	Email      string ` json:"email" binding:"required" `
	Password   string ` json:"password" `
	Phone      string ` json:"phone" `
	Otoritas   uint32 ` json:"otoritas"  `
	Status     string `json:"status"  `
}

type ListUserDTO struct {
	Search              *string `json:"Search"`
	Transaction_tglDari string  `json:"Transaction_tglDari"`
	Transaction_tglTo   string  `json:"Transaction_tglTo"`
	Limit               string  `json:"Limit"`
	Page                string  `json:"Page"`
	Order               string  `json:"Order"`
}
