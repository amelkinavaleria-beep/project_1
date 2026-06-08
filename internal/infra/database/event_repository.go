package database

import (
    "time"
    "math"
    "github.com/upper/db/v4"
    "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

const EventsTableName = "events"

type event struct {
    Id          uint64     `db:"id,omitempty"`
    DeviceId    uint64     `db:"device_id"`
    RoomId      uint64     `db:"room_id"`
    Action      string     `db:"action"`
    CreatedDate time.Time  `db:"created_date"`
    UpdatedDate time.Time  `db:"updated_date"`
    DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type EventRepository interface {
    Save(e domain.Event) (domain.Event, error)
    FindList(p domain.Pagination, deviceId, roomId uint64) (domain.Events, error)
}

type eventRepository struct {
    coll db.Collection
}

func NewEventRepository(session db.Session) EventRepository {
    return eventRepository{coll: session.Collection(EventsTableName)}
}

func (r eventRepository) Save(e domain.Event) (domain.Event, error) {
    ev := event{
        DeviceId:    e.DeviceId,
        RoomId:      e.RoomId,
        Action:      e.Action,
        CreatedDate: time.Now(),
        UpdatedDate: time.Now(),
    }
    err := r.coll.InsertReturning(&ev)
    if err != nil {
        return domain.Event{}, err
    }
    return r.mapModelToDomain(ev), nil
}

func (r eventRepository) FindList(p domain.Pagination, deviceId, roomId uint64) (domain.Events, error) {
    var evs []event
    if p.Page == 0 {
        p.Page = 1
    }
    if p.CountPerPage == 0 {
        p.CountPerPage = 20
    }

    query := r.coll.Find(db.Cond{
        "device_id": deviceId,
        "room_id":   roomId,
        "deleted_date": nil,
    }).OrderBy("-created_date")

    res := query.Paginate(uint(p.CountPerPage))
    err := res.Page(uint(p.Page)).All(&evs)
    if err != nil {
        return domain.Events{}, err
    }

    events := r.mapModelToDomainCollection(evs)
    totalCount, err := res.TotalEntries()
    if err != nil {
        return domain.Events{}, err
    }
    events.Total = totalCount
    events.Pages = uint(math.Ceil(float64(events.Total) / float64(p.CountPerPage)))
    return events, nil
}

func (r eventRepository) mapModelToDomain(e event) domain.Event {
    return domain.Event{
        Id:          e.Id,
        DeviceId:    e.DeviceId,
        RoomId:      e.RoomId,
        Action:      e.Action,
        CreatedDate: e.CreatedDate,
        UpdatedDate: e.UpdatedDate,
        DeletedDate: e.DeletedDate,
    }
}

func (r eventRepository) mapModelToDomainCollection(evs []event) domain.Events {
    events := make([]domain.Event, len(evs))
    for i, v := range evs {
        events[i] = r.mapModelToDomain(v)
    }
    return domain.Events{Items: events}
}
