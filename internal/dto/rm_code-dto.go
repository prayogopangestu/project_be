package dto

import "project/internal/models"

type RmCodeCreateDTO models.RmCodeM

type RmCodeUpdateDTO models.RmCodeM

type ListRmCodeDTO struct {
	Search              *string `json:"Search"`
	Transaction_tglDari string  `json:"Transaction_tglDari"`
	Transaction_tglTo   string  `json:"Transaction_tglTo"`
	Limit               string  `json:"Limit"`
	Page                string  `json:"Page"`
	Order               string  `json:"Order"`
}
