package utils

import "time"

type Time struct {
	time.Time
}

func (m *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = Time{tt}
	return err
}
