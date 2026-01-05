package usecase

import (
	"crypto/sha256"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTodaysUUID(t *testing.T) {
	uc := NewUUIDUseCase()

	// 1. 逕滓・繝・せ繝・
	result := uc.GetTodaysUUID()
	assert.NotNil(t, result, "Result should not be nil")
	assert.NotEmpty(t, result.UUID, "UUID should not be empty")

	// 2. 繝輔か繝ｼ繝槭ャ繝育｢ｺ隱・(髟ｷ縺・6譁・ｭ・ 8-4-4-4-12)
	assert.Len(t, result.UUID, 36, "Expected UUID length 36")

	// 3. 豎ｺ螳夊ｫ也噪縺ｧ縺ゅｋ縺九・遒ｺ隱・
	result2 := uc.GetTodaysUUID()
	assert.Equal(t, result.UUID, result2.UUID, "UUID should be deterministic")

	// 4. (蜿り・ 謇句虚險育ｮ励→荳閾ｴ縺吶ｋ縺・
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)
	dateStr := now.Format("2006-01-02")

	h := sha256.New()
	h.Write([]byte(dateStr))
	bs := h.Sum(nil)
	expectedUUID := fmt.Sprintf("%x-%x-%x-%x-%x", bs[0:4], bs[4:6], bs[6:8], bs[8:10], bs[10:16])

	assert.Equal(t, expectedUUID, result.UUID, "UUID calculation mismatch")
}
