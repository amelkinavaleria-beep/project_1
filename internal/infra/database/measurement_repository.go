package database

import (
	"fmt"
	"math"
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const MeasurementsTableName = "measurements"

type measurement struct {
	Id          uint64     `db:"id,omitempty"`
	DeviceId    uint64     `db:"device_id"`
	RoomId      uint64     `db:"room_id"`
	Value       uint64     `db:"value"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type measurementRepository struct {
	coll db.Collection
	sess db.Session
}

type MeasurementRepository interface {
	Save(m domain.Measurement) (domain.Measurement, error)
	FindList(p domain.Pagination, mf MeasurementFilters) (domain.Measurements, error)
}

func NewMeasurementRepository(session db.Session) MeasurementRepository {
	return measurementRepository{
		sess: session,
		coll: session.Collection(MeasurementsTableName),
	}
}

type MeasurementFilters struct {
	DeviceId uint64
	RoomId   uint64

	CreatedDateFrom *time.Time
	CreatedDateTo   *time.Time

	Sort string
}

func (r measurementRepository) Save(m domain.Measurement) (domain.Measurement, error) {
	ms := measurement{
		DeviceId:    m.DeviceId,
		RoomId:      m.RoomId,
		Value:       m.Value,
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	err := r.coll.InsertReturning(&ms)
	if err != nil {
		return domain.Measurement{}, err
	}
	return r.mapModelToDomain(ms), nil
}

func (r measurementRepository) FindList(
	p domain.Pagination,
	mf MeasurementFilters,
) (domain.Measurements, error) {
	var ms []measurement

	if mf.DeviceId == 0 {
		return domain.Measurements{}, fmt.Errorf("device_id is required")
	}

	if mf.RoomId == 0 {
		return domain.Measurements{}, fmt.Errorf("room_id is required")
	}

	if p.Page == 0 {
		p.Page = 1
	}

	if p.CountPerPage == 0 {
		p.CountPerPage = 20
	}

	query := r.coll.
		Find(db.Cond{
			"device_id":    mf.DeviceId,
			"room_id":      mf.RoomId,
			"deleted_date": nil,
		})

	if mf.CreatedDateFrom != nil {
		query = query.And("created_date >= ?", *mf.CreatedDateFrom)
	}

	if mf.CreatedDateTo != nil {
		query = query.And("created_date <= ?", *mf.CreatedDateTo)
	}

	switch mf.Sort {
	case "created_date":
		query = query.OrderBy("created_date")
	case "-created_date":
		query = query.OrderBy("-created_date")
	case "value":
		query = query.OrderBy("value")
	case "-value":
		query = query.OrderBy("-value")
	default:
		query = query.OrderBy("-created_date")
	}

	res := query.Paginate(uint(p.CountPerPage))

	err := res.Page(uint(p.Page)).All(&ms)
	if err != nil {
		return domain.Measurements{}, err
	}

	measurements := r.mapModelToDomainPagination(ms)

	totalCount, err := res.TotalEntries()
	if err != nil {
		return domain.Measurements{}, err
	}

	measurements.Total = totalCount
	measurements.Pages = uint(math.Ceil(float64(measurements.Total) / float64(p.CountPerPage)))

	return measurements, nil
}

func (r measurementRepository) mapModelToDomainPagination(ms []measurement) domain.Measurements {
	measurements := make([]domain.Measurement, 0, len(ms))

	for _, m := range ms {
		measurements = append(measurements, domain.Measurement{
			Id:          m.Id,
			DeviceId:    m.DeviceId,
			RoomId:      m.RoomId,
			Value:       m.Value,
			CreatedDate: m.CreatedDate,
			UpdatedDate: m.UpdatedDate,
			DeletedDate: m.DeletedDate,
		})
	}

	return domain.Measurements{
		Items: measurements,
	}
}

func (r measurementRepository) mapModelToDomain(m measurement) domain.Measurement {
	return domain.Measurement{
		Id:          m.Id,
		DeviceId:    m.DeviceId,
		RoomId:      m.RoomId,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}