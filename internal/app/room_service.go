package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type roomService struct {
	roomRepo database.RoomRepository
}

type RoomService interface {
	Save(rm domain.Room) (domain.Room, error)
	FindList(orgId uint64, page int) (domain.Rooms, error)
	Update(rm domain.Room) (domain.Room, error)
	Delete(id uint64) error
	Find(id uint64) (interface{}, error)
}

func NewRoomService(rr database.RoomRepository) RoomService {
	return roomService{
		roomRepo: rr,
	}
}

func (s roomService) Save(rm domain.Room) (domain.Room, error) {
	rm, err := s.roomRepo.Save(rm)
	if err != nil {
		log.Printf("roomService.Save(s.roomRepo.Save): %s", err)
		return domain.Room{}, err
	}

	return rm, nil
}

func (s roomService) Find(rid uint64) (interface{}, error) {
	rm, err := s.roomRepo.FindByRoomId(rid)
	if err != nil {
		log.Printf("roomService.Find(s.roomRepo.FindByRoomId): %s", err)
		return nil, err
	}

	return rm, nil

}

func (s roomService) Update(rm domain.Room) (domain.Room, error) {
	rms, err := s.roomRepo.Update(rm)
	if err != nil {
		log.Printf("roomService.Update(s.roomRepo.Update): %s", err)
		return domain.Room{}, err
	}

	return rms, nil
}

func (s roomService) Delete(id uint64) error {
	err := s.roomRepo.Delete(id)
	if err != nil {
		log.Printf("roomService.Delete(s.roomRepo.Delete): %s", err)
		return err
	}

	return nil
}

func (s roomService) FindList(orgId uint64, page int) (domain.Rooms, error) {
	pagination := domain.Pagination{
		Page:         uint64(page),
		CountPerPage: 20,
	}
	return s.roomRepo.FindList(pagination, orgId)
}
