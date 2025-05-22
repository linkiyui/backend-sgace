package domain

import "time"

type Solicitude struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	ActivityID string    `json:"activity_id"`
	Group      string    `json:"group"`
	Faculty    faculty   `json:"faculty"`
	Grade      string    `json:"grade"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Status     Status    `json:"status"`
}

type Status int

const (
	Accepted Status = 1
	Rejected Status = 0
)

type faculty string

const (
	FIO            faculty = "FIO"
	FTI            faculty = "FTI"
	FTE            faculty = "FTE"
	CITEC          faculty = "CITEC"
	FTL            faculty = "FTL"
	CIBERSEGURIDAD faculty = "CIBERSEGURIDAD"
)

func Faculty(faculty string) faculty {
	switch faculty {
	case "FIO":
		return FIO
	case "FTI":
		return FTI
	case "FTE":
		return FTE
	case "CITEC":
		return CITEC
	case "FTL":
		return FTL
	case "CIBERSEGURIDAD":
		return CIBERSEGURIDAD
	default:
		return ""
	}
}
