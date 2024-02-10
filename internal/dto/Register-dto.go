package dto

type RegisterDTO struct {
	Nickname   string ` json:"nickname" binding:"required"`
	First_name string ` json:"first_name" binding:"required"`
	Last_name  string ` json:"last_name" binding:"required"`
	Email      string ` json:"email" binding:"required"`
	Password   string ` json:"password" binding:"required"`
	Phone      string ` json:"phone" binding:"required"`
	Otoritas   uint32 ` json:"otoritas" binding:"required"`
	Status     string `json:"status" binding:"required"`
}
