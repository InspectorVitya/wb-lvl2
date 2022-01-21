package storage

import (
	"context"
	"github.com/inspectorvitya/wb-lvl2/develop/dev11/internal/model"
)

type CalendarStorage interface {
	Save(ctx context.Context, event model.Event) error
	Update(ctx context.Context, event model.Event) error
	Delete(ctx context.Context, id int) error
	Load(ctx context.Context, interval string) ([]model.Event, error)
}
