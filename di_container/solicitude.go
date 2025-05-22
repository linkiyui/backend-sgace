package dicontainer

import (
	solicitude_service "github.com/sgace/backend/solicitude/application"
	solicitude_infra "github.com/sgace/backend/solicitude/infra"
	"github.com/sgace/database"
)

func SolicitudeService() *solicitude_service.SolicitudeService {
	repo := solicitude_infra.NewSolicitudePostgresRepository(database.Database)
	return solicitude_service.NewSolicitudeService(repo)
}
