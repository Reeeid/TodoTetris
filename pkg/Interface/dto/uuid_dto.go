package dto

import "github.com/Reeeid/TodoTetris/pkg/Domain/model"

type UUIDResponse struct {
	UUID string `json:"uuid"`
}

func FromUUIDDomain(m *model.UUID) *UUIDResponse {
	return &UUIDResponse{
		UUID: m.UUID,
	}

}
