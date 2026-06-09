package database

import (
	"math"
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const DevicesTableName = "devices"

type device struct {
	Id               uint64                `db:"id,omitempty"`
	OrganizationId   uint64                `db:"organization_id"`
	RoomId           *uint64               `db:"room_id"`
	GUID             string                `db:"guid"`
	InventoryNumber  string                `db:"inventory_number"`
	SerialNumber     string                `db:"serial_number"`
	Characteristics  string                `db:"characteristics"`
	Category         domain.DeviceCategory `db:"category"`
	Units            string                `db:"units"`
	PowerConsumption float64               `db:"power_consumption"`
	CreatedDate      time.Time             `db:"created_date"`
	UpdatedDate      time.Time             `db:"updated_date"`
	DeletedDate      *time.Time            `db:"deleted_date"`
}

type DeviceRepository interface {
	Save(d domain.Device) (domain.Device, error)
	Find(id uint64) (domain.Device, error)
	Update(d domain.Device) (domain.Device, error)
	Delete(id uint64) error
	FindList(p domain.Pagination, oId uint64) (domain.Devices, error)
}

type deviceRepository struct {
	coll db.Collection
	sess db.Session
}

func NewDeviceRepository(session db.Session) DeviceRepository {
	return deviceRepository{
		coll: session.Collection(DevicesTableName),
		sess: session,
	}
}

func (r deviceRepository) Save(d domain.Device) (domain.Device, error) {
	dev := r.mapDomainToModel(d)
	now := time.Now()
	dev.CreatedDate = now
	dev.UpdatedDate = now

	err := r.coll.InsertReturning(&dev)
	if err != nil {
		return domain.Device{}, err
	}
	return r.mapModelToDomain(dev), nil
}

func (r deviceRepository) Find(id uint64) (domain.Device, error) {
	var dev device
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&dev)
	if err != nil {
		return domain.Device{}, err
	}
	return r.mapModelToDomain(dev), nil
}

func (r deviceRepository) Update(d domain.Device) (domain.Device, error) {
	dev := r.mapDomainToModel(d)
	dev.UpdatedDate = time.Now()

	err := r.coll.Find(db.Cond{"id": d.Id, "deleted_date": nil}).Update(&dev)
	if err != nil {
		return domain.Device{}, err
	}
	return r.mapModelToDomain(dev), nil
}

func (r deviceRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{
		"deleted_date": time.Now(),
	})
}

func (r deviceRepository) FindList(p domain.Pagination, oId uint64) (domain.Devices, error) {
	var devs []device
	if p.Page == 0 {
		p.Page = 1
	}
	if p.CountPerPage == 0 {
		p.CountPerPage = 20
	}

	query := r.coll.Find(db.Cond{"organization_id": oId, "deleted_date": nil})

	res := query.Paginate(uint(p.CountPerPage))
	err := res.Page(uint(p.Page)).All(&devs)
	if err != nil {
		return domain.Devices{}, err
	}

	totalCount, err := res.TotalEntries()
	if err != nil {
		return domain.Devices{}, err
	}

	return domain.Devices{
		Items: r.mapModelToDomainCollection(devs),
		Total: totalCount,
		Pages: uint(math.Ceil(float64(totalCount) / float64(p.CountPerPage))),
	}, nil
}
func (r deviceRepository) mapDomainToModel(d domain.Device) device {
	return device{
		Id:               d.Id,
		OrganizationId:   d.OrganizationId,
		RoomId:           d.RoomId,
		GUID:             d.GUID,
		InventoryNumber:  d.InventoryNumber,
		SerialNumber:     d.SerialNumber,
		Characteristics:  d.Characteristics,
		Category:         d.Category,
		Units:            d.Units,
		PowerConsumption: d.PowerConsumption,
		CreatedDate:      d.CreatedDate,
		UpdatedDate:      d.UpdatedDate,
		DeletedDate:      d.DeletedDate,
	}
}

func (r deviceRepository) mapModelToDomain(d device) domain.Device {
	return domain.Device{
		Id:               d.Id,
		OrganizationId:   d.OrganizationId,
		RoomId:           d.RoomId,
		GUID:             d.GUID,
		InventoryNumber:  d.InventoryNumber,
		SerialNumber:     d.SerialNumber,
		Characteristics:  d.Characteristics,
		Category:         d.Category,
		Units:            d.Units,
		PowerConsumption: d.PowerConsumption,
		CreatedDate:      d.CreatedDate,
		UpdatedDate:      d.UpdatedDate,
		DeletedDate:      d.DeletedDate,
	}
}

func (r deviceRepository) mapModelToDomainCollection(devs []device) []domain.Device {
	devices := make([]domain.Device, len(devs))
	for i := range devs {
		devices[i] = r.mapModelToDomain(devs[i])
	}
	return devices
}
