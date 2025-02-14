package query

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/truesch/wild-workouts-go-ddd-example/internal/common/errors"
)

type AvailableHoursHandler struct {
	readModel AvailableHoursReadModel
}

type AvailableHoursReadModel interface {
	AvailableHours(ctx context.Context, from time.Time, to time.Time) ([]Date, error)
}

func NewAvailableHoursHandler(readModel AvailableHoursReadModel) AvailableHoursHandler {
	return AvailableHoursHandler{readModel: readModel}
}

type AvailableHours struct {
	From time.Time
	To   time.Time
}

func (h AvailableHoursHandler) Handle(ctx context.Context, query AvailableHours) (d []Date, err error) {
	start := time.Now()
	defer func() {
		logrus.
			WithError(err).
			WithField("duration", time.Since(start)).
			Debug("AvailableHoursHandler executed")
	}()

	if query.From.After(query.To) {
		return nil, errors.NewIncorrectInputError("date-from-after-date-to", "Date from after date to")
	}

	return h.readModel.AvailableHours(ctx, query.From, query.To)
}
