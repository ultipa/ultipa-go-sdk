package utils

import (
	"errors"
	"time"
)

type UltipaTime struct {
	Datetime uint64
	Year     uint64
	Month    uint64
	Day      uint64
	Hour     uint64
	Minute   uint64
	Second   uint64
	Macrosec uint64
	Time     *time.Time
}

func NewTimeStamp(datetime int64) *UltipaTime {
	unixTime := time.Unix(datetime, 0)
	ultipaDateTime := TimeToUint64(unixTime)
	return NewTime(ultipaDateTime)
}

func NewTime(datetime uint64) *UltipaTime {
	n := UltipaTime{
		Datetime: datetime,
	}
	if datetime == 0 {
		unix := time.Unix(0, 0)
		n.Time = &unix
	} else {
		n.Uint64ToTime(datetime)
	}
	return &n
}

// StringToTime , layoutISO := "2006-01-02 15:04:05.000"
// StringToTime , layoutISO := "2006-01-02 15:04:05"
// StringToTime , layoutISO := "2006-01-02"
func NewTimeFromStringFormat(dateString string, format string) (*UltipaTime, error) {

	t, err := time.Parse(format, dateString)

	if err != nil {
		return nil, err
	}

	n := UltipaTime{}
	n.Time = &t
	n.Datetime = n.TimeToUint64(&t)

	return &n, err
}
func NewTimeFromString(dateString string) (*UltipaTime, error) {
	n := UltipaTime{}
	layouts := []string{
		"2006-1-2",
		"2006-1-2T15:04:05Z07:00",
		"2006-1-2 15:04:05.000",
		"2006-1-2 15:04:05",
		"2006-1-2 15:04",
		"2006-1-2 15",
		"2006/1/2",
		"2006/1/2T15:04:05Z07:00",
		"2006/1/2 15:04:05.000",
		"2006/1/2 15:04:05",
		"2006/1/2 15:04",
		"2006/1/2 15",
		"2006-01-02",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02 15:04:05.000",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02 15",
	}

	for _, l := range layouts {
		t, err := time.Parse(l, dateString)
		n.Time = &t
		v := n.TimeToUint64(&t)

		if err == nil {
			n.Datetime = v
			return &n, err
		}
	}

	return nil, errors.New("parse datetime string failed : " + dateString)
}

// parse Bytes to year, month....
//int year_month = ((datetime >> 46) & 0x1FFFF);
//int year = year_month / 13;
//int month = year_month % 13;
//int day = ((datetime >> 41) & 0x1F);
//int hour = ((datetime >> 36) & 0x1F);
//int minute = ((datetime >> 30) & 0x3F);
//int second = ((datetime >> 24) & 0x3F);
//int microsec = (datetime & 0xFFFFFF);

//uint64_t year = 0;
//uint64_t month = 0;
//uint64_t day = 0;
//uint64_t hour = 0;
//uint64_t minute = 0;
//uint64_t second = 0;
//uint64_t macrosec = 0;
//
//if (year > 70 && year < 100) {
//year += 1900;
//} else if (year < 70) {
//year += 2000;
//}
//

func (t *UltipaTime) Uint64ToTime(datetime uint64) (_t *time.Time) {
	year_month := (datetime >> 46) & 0x1FFFF

	t.Year = year_month / 13

	if t.Year > 70 && t.Year < 100 {
		t.Year += 1900
	} else if t.Year < 70 {
		t.Year += 2000
	}

	t.Datetime = datetime
	t.Month = year_month % 13
	t.Day = (datetime >> 41) & 0x1F
	t.Hour = (datetime >> 36) & 0x1F
	t.Minute = (datetime >> 30) & 0x3F
	t.Second = (datetime >> 24) & 0x3F
	t.Macrosec = datetime & 0xFFFFFF

	date := time.Date(int(t.Year), time.Month(t.Month), int(t.Day), int(t.Hour), int(t.Minute), int(t.Second), int(t.Macrosec*1000), time.UTC)
	t.Time = &date

	return t.Time
}

//uint64_t datetime = 0;
//uint64_t year_month = year * 13 + month;
//datetime |= (year_month << 46);
//datetime |= (day << 41);
//datetime |= (hour << 36);
//datetime |= (minute << 30);
//datetime |= (second << 24);
//datetime |= macrosec;
func (u *UltipaTime) TimeToUint64(time *time.Time) uint64 {

	u.Datetime = TimeToUint64(*time)

	return u.Datetime
}

func TimeToUint64(time time.Time) uint64 {

	datetime := uint64(0)

	Year := uint64(time.Year())
	Month := uint64(time.Month())
	Day := uint64(time.Day())
	Hour := uint64(time.Hour())
	Minute := uint64(time.Minute())
	Second := uint64(time.Second())
	Macrosec := uint64(time.Nanosecond() / 1000)

	yearMonth := Year*13 + Month
	datetime = yearMonth << 46
	datetime = datetime | (Day << 41)
	datetime = datetime | (Hour << 36)
	datetime = datetime | (Minute << 30)
	datetime = datetime | (Second << 24)
	datetime = datetime | Macrosec

	return datetime
}

func (u *UltipaTime) String() string {
	return u.Time.Format("2006-01-02 15:04:05.000")
}

// Get Timestamp , Second
func (u *UltipaTime) GetTimeStamp() uint32 {
	return uint32(u.Time.Unix())
}
