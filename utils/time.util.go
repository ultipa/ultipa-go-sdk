package utils

import (
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
	Time *time.Time
}

func (t *UltipaTime) New(datetime uint64) *UltipaTime{
	n := UltipaTime{
		Datetime: datetime,
	}
	n.Uint64ToTime(datetime)
	return &n
}

// StringToTime , layoutISO := "2006-01-02 15:04:05.000"
// StringToTime , layoutISO := "2006-01-02 15:04:05"
// StringToTime , layoutISO := "2006-01-02"
func (u *UltipaTime) NewFromString(dateString string) (*UltipaTime, error){
	var err error
	n := UltipaTime{}
	layouts := []string {
		"2006-01-02 15:04:05.000",
		"2006-01-02 15:04:05",
		"2006-01-02",
	}

	for _, l := range layouts {
		t, err := time.Parse(l, dateString)
		n.Time = &t
		n.TimeToUint64(&t)
		if err == nil {
			return &n, err
		}
	}

	return nil, err
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
func  (t *UltipaTime) Uint64ToTime(datetime uint64)  (_t *time.Time) {
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

	date := time.Date(int(t.Year),time.Month(t.Month),int(t.Day),int(t.Hour),int(t.Minute),int(t.Second),int(t.Macrosec * 1000), time.UTC)
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

	datetime :=  uint64(0)

	u.Year = uint64(time.Year())
	u.Month = uint64(time.Month())
	u.Day = uint64(time.Day())
	u.Hour = uint64(time.Hour())
	u.Minute = uint64(time.Minute())
	u.Second = uint64(time.Second())
	u.Macrosec = uint64(time.Nanosecond() / 1000)

	yearMonth := u.Year * 13 + u.Month
	datetime = yearMonth << 46
	datetime = datetime | (u.Day << 41)
	datetime = datetime | (u.Hour << 36)
	datetime = datetime | (u.Minute << 30)
	datetime = datetime | (u.Second << 24)
	datetime = datetime | u.Macrosec

	u.Datetime = datetime

	return datetime
}



func (u *UltipaTime) ToString() string {
	return u.Time.Format("2006-01-02 15:04:05.000Z")
}