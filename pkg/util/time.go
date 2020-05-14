package util

import "time"

type (
	Time time.Time
)

const DefaultTimeFormat = "1970-01-01 00:00:00"

func (t *Time) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(`"`+DefaultTimeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return err
}

func (t *Time) MarshaJSON() ([]byte, error) {
	if time.Time(*t).IsZero() {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(DefaultTimeFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, DefaultTimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) String() string {
	return time.Time(*t).Format(DefaultTimeFormat)
}

func (t *Time) isZero() bool {
	return time.Time(*t).Second() == 0 && time.Time(*t).Nanosecond() == 0
}
