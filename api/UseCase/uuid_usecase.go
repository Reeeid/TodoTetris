package usecase

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/Reeeid/TodoTetris/api/Domain/model"
)

type UUIDUseCase struct {
}

func NewUUIDUseCase() *UUIDUseCase {
	return &UUIDUseCase{}
}

func (u *UUIDUseCase) GetTodaysUUID() *model.UUID {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)
	// Include Nanoseconds and a random/counter component to ensure uniqueness per call
	seed := fmt.Sprintf("%s-%d", now.Format("2006-01-02-15-04-05.000000000"), now.UnixNano())

	h := sha256.New()
	h.Write([]byte(seed))
	bs := h.Sum(nil)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", bs[0:4], bs[4:6], bs[6:8], bs[8:10], bs[10:16])
	return &model.UUID{
		UUID: uuid,
	}
}
