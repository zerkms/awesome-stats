package lib

import "testing"
import "github.com/stretchr/testify/assert"

type AlwaysUpdater struct {
	cnt *int
}

func (u AlwaysUpdater) Update(page string, stats []Stats) (updateResult, error) {
	*u.cnt++

	return updateResult{
		stats: updateStats{
			found: 1,
		},
	}, nil
}

func TestOnlyFirstFoundIsApplied(t *testing.T) {
	a1, a2 := AlwaysUpdater{
		cnt: new(int),
	}, AlwaysUpdater{
		cnt: new(int),
	}

	u := NewCompositeUpdater([]Updater{a1, a2})
	u.Update("", []Stats{Stats{}})

	assert.Equal(t, 1, *a1.cnt)
	assert.Equal(t, 0, *a2.cnt)
}
