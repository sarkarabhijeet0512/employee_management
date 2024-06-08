package employee

import (
	"sync"
	"time"

	"go.uber.org/fx"
)

// Module provides all constructor and invocation methods to facilitate credits module
var Module = fx.Options(
	fx.Provide(
		NewDBRepository,
		NewService,
	),
)

type (
	// User represents the user entity
	Employee struct {
		ID        int       `json:"id" pg:"id,pk"`
		Name      string    `json:"name" pg:"name" binding:"required"`
		Mobile    string    `json:"mobile" pg:"mobile,unique" binding:"required"`
		Position  string    `json:"position" pg:"position"`
		Salary    float64   `json:"salary" pg:"salary"`
		IsActive  bool      `json:"is_active" pg:"is_active"`
		CreatedAt time.Time `json:"created_at" pg:"created_at"`
		UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
		Mutex     *sync.Mutex
	}
)
