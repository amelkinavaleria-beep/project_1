package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type RoomDto struct {
	Id             uint64  `json:"id"`
	OrganizationId uint64  `json:"organizationId"`
	Name           string  `json:"name"`
	Description    *string `json:"description,omitempty"`
}

func (d RoomDto) DomainToDto(rm domain.Room) RoomDto {
	return RoomDto{
		Id:             rm.Id,
		OrganizationId: rm.OrganizationId,
		Name:           rm.Name,
		Description:    rm.Description,
	}
}

func (d RoomDto) DomainToDtoCollection(rms []domain.Room) []RoomDto {
	rmsDto := make([]RoomDto, len(rms))
	for i := range rms {
		rmsDto[i] = d.DomainToDto(rms[i])
	}
	return rmsDto
}