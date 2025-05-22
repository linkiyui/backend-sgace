package domain

type ISolicitudeRepository interface {
	CreateActivity(activity *Activity) error
	UpdateActivity(activity *Activity) error
	GetActivityByID(id string) (*Activity, error)
	DeleteActivity(activity_id string) error
	GetAllActivities() ([]Activity, error)
}
