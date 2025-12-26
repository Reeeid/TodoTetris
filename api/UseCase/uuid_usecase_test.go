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

	// 1. 生成テスト
	result := uc.GetTodaysUUID()
	assert.NotNil(t, result, "Result should not be nil")
	assert.NotEmpty(t, result.UUID, "UUID should not be empty")

	// 2. フォーマット確認 (長さ36文字: 8-4-4-4-12)
	assert.Len(t, result.UUID, 36, "Expected UUID length 36")

	// 3. 決定論的であるかの確認
	result2 := uc.GetTodaysUUID()
	assert.Equal(t, result.UUID, result2.UUID, "UUID should be deterministic")

	// 4. (参考) 手動計算と一致するか
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)
	dateStr := now.Format("2006-01-02")

	h := sha256.New()
	h.Write([]byte(dateStr))
	bs := h.Sum(nil)
	expectedUUID := fmt.Sprintf("%x-%x-%x-%x-%x", bs[0:4], bs[4:6], bs[6:8], bs[8:10], bs[10:16])

	assert.Equal(t, expectedUUID, result.UUID, "UUID calculation mismatch")
}
