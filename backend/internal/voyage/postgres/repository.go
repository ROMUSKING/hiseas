package postgres

import (
	"context"
	"fmt"

	"github.com/ROMUSKING/hiseas/backend/internal/shared"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VoyageRepository struct {
	pool *pgxpool.Pool
}

func NewVoyageRepository(pool *pgxpool.Pool) *VoyageRepository {
	return &VoyageRepository{pool: pool}
}

// FindNearbyVoyages searches for voyages with start_location within a given radius from a point
func (r *VoyageRepository) FindNearbyVoyages(ctx context.Context, lon, lat, radiusMeters float64) ([]shared.Voyage, error) {
	query := `SELECT id, skipper_id, vessel_id, title, description, start_time, end_time, ST_AsText(start_location), ST_AsText(planned_route)
		FROM voyages
		WHERE ST_DWithin(start_location, ST_MakePoint($1, $2)::geography, $3)`
	rows, err := r.pool.Query(ctx, query, lon, lat, radiusMeters)
	if err != nil {
		return nil, fmt.Errorf("error querying nearby voyages: %w", err)
	}
	defer rows.Close()

	var voyages []shared.Voyage
	for rows.Next() {
		var v shared.Voyage
		var startLoc, plannedRoute string
		if err := rows.Scan(&v.ID, &v.SkipperID, &v.VesselID, &v.Title, &v.Description, &v.StartTime, &v.EndTime, &startLoc, &plannedRoute); err != nil {
			return nil, err
		}
		v.StartLocation = startLoc
		v.PlannedRoute = plannedRoute
		voyages = append(voyages, v)
	}
	return voyages, nil
}
