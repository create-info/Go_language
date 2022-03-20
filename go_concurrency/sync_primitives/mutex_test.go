package sync_primitives_test

import (
	"Go_language/go_concurrency/sync_primitives"
	"testing"
)

func TestSumWithOutMutex(t *testing.T) {
	sync_primitives.SumWithOutMutex()
}
