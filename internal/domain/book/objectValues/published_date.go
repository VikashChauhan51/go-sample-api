package objectValues

import (
	"time"
)

type PublishedDate struct {
	value time.Time
}

func NewPublishedDate(value time.Time) PublishedDate {
	return PublishedDate{value: value}
}

func (pd PublishedDate) Value() time.Time {
	return pd.value
}
