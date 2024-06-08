package model

import "time"

type (
	GenericRes struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    any         `json:"data,omitempty"`
		Meta    *Pagination `json:"meta,omitempty"`
	}
	EmployeeFilter struct {
		ID         int       `form:"id" json:"id"`
		Name       string    `form:"name" json:"name"`
		Position   string    `form:"position" json:"position"`
		Salary     float64   `form:"salary" json:"salary"`
		Page       int       `form:"page,default=1"`
		Limit      int       `form:"limit,default=20"`
		ReturnData bool      `form:"return_data" json:"return_data"`
		IsActive   *bool     `form:"is_active" json:"is_active"`
		CreatedAt  time.Time `form:"created_at" time_format:"2006-01-02"`
	}
	Pagination struct {
		CurrentPage    int `json:"current_page,omitempty"`
		TotalPages     int `json:"total_pages,omitempty"`
		TotalDataCount int `json:"total_data_count,omitempty"`
	}
	UpdateEmployeeReq struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Mobile    string    `json:"mobile"`
		Position  string    `json:"position"`
		Salary    float64   `json:"salary"`
		IsActive  bool      `json:"is_active"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
