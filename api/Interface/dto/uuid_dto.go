package dto

import "github.com/Reeeid/TodoTetris/Domain/model"

type UUIDResponse struct {
	UUID string `json:"uuid"`
}

func (u *UUIDResponse) ToDomain() *model.UUID {
	return &model.UUID{
		UUID: u.UUID,
	}
}

func FromUUIDDomain(m *model.UUID) *UUIDResponse {
	return &UUIDResponse{
		UUID: m.UUID,
	}

}
