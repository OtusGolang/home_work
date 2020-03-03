// +build bench

package hw10_program_optimization //nolint:golint,stylecheck

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	mb          int64 = 1 << 20
	memoryLimit       = 30 * mb

	timeLimit = 140 * time.Millisecond
)

// go test -v -count=1 -timeout=30s -tags bench .
// Locally ~70 ms, ~28 Mb
func TestGetDomainStat_Time_And_Memory(t *testing.T) {
	bench := func(b *testing.B) {
		data, err := os.Open("testdata/users.dat")
		require.NoError(t, err)

		b.StartTimer()
		_, err = GetDomainStat(data, "biz")
		b.StopTimer()
		require.NoError(t, err)
	}

	result := testing.Benchmark(bench)
	mem := int64(result.MemBytes)/mb
	t.Logf("time used: %s", result.T)
	t.Logf("memory used: %dMb", mem)

	require.Less(t, int64(result.T), int64(timeLimit), "the program is too slow")
	require.Less(t, mem, memoryLimit, "the program is too greedyot ")
}
