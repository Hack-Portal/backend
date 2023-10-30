package gateways

import (
	"temp/src/entities"
	"temp/src/entities/param"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	h := NewHackathonGateway(dbconn, nil)

	arg := param.CreateHackathon{
		Hackathon: &entities.Hackathon{
			HackathonID: uuid.New().String(),
			Name:        "Hackathon 1",
			Description: "Hackathon 1",
			Icon:        "Sample",
			Link:        "https://www.google.com",
			Expired:     time.Now().AddDate(0, 0, 1),
			StartDate:   time.Now(),
			Term:        1,
			CreatedAt:   time.Now(),
			IsDelete:    false,
		},
		StatusTags: []int{1, 2},
	}

	require.NoError(t, h.Create(arg))
}

func BenchmarkCreate(b *testing.B) {
	h := NewHackathonGateway(dbconn, nil)
	b.ResetTimer()

	b.N = 10000
	for i := 0; i < b.N; i++ {
		arg := param.CreateHackathon{
			Hackathon: &entities.Hackathon{
				HackathonID: uuid.New().String(),
				Name:        "Hackathon 1",
				Description: "Hackathon 1",
				Icon:        "Sample",
				Link:        "https://www.google.com",
				Expired:     time.Now().AddDate(0, 0, 1),
				StartDate:   time.Now(),
				Term:        1,
				CreatedAt:   time.Now(),
				IsDelete:    false,
			},
			StatusTags: []int{1, 2},
		}

		h.Create(arg)
	}
}
