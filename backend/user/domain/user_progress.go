package domain

type UserProgress struct {
	ToatlActivities     int64 `json:"total_activities"`
	AcceptedSolicitudes int64 `json:"accepted_solicitudes"`
	TotalSolicitudes    int64 `json:"total_solicitudes"`
	RefuseSolicitudes   int64 `json:"refuse_solicitudes"`
	Completed           int64 `json:"completed"`
}
