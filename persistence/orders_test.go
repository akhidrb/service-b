package persistence

import (
	"testing"
	"time"
)

func TestTimeBeginningOfDay(t *testing.T) {
	t.Log(time.Now().Truncate(24 * time.Hour))
	t.Log(time.Now().Add(24 * time.Hour).Truncate(24 * time.Hour))
}
