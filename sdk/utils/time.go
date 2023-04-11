package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	return TimeToUltipaTime(&unixTime)
}

func NewTime(datetime uint64) *UltipaTime {
	n := UltipaTime{
		Datetime: datetime,
	}
	if datetime == 0 {
		//use same location UTC as n.Uint64ToTime
		unix := time.Unix(0, 0).UTC()
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
	newDateString, err := compensateYear(strings.Trim(dateString, " "))
	if err != nil {
		return nil, err
	}
	layouts := []string{
		"2006-1-2",
		"2006-1-2T15:04:05.000Z0700",
		"2006-1-2T15:04:05.000Z07:00",
		"2006-1-2T15:04:05Z0700",
		"2006-1-2T15:04:05Z07:00",
		"2006-1-2 15:04:05.000",
		"2006-1-2 15:04:05",
		"2006-1-2 15:04",
		"2006-1-2 15",
		"2006/1/2",
		"2006/1/2T15:04:05.000Z07:00",
		"2006/1/2T15:04:05.000Z0700",
		"2006/1/2T15:04:05Z07:00",
		"2006/1/2T15:04:05Z0700",
		"2006/1/2 15:04:05.000",
		"2006/1/2 15:04:05",
		"2006/1/2 15:04",
		"2006/1/2 15",
		"2006-01-02",
		"2006-01-02T15:04:05.000Z07:00",
		"2006-01-02T15:04:05.000Z0700",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z0700",
		"2006-01-02 15:04:05.000",
		"2006010215:04:05.000Z0700",
		"2006010215:04:05.000Z07:00",
		"2006010215:04:05Z0700",
		"2006010215:04:05Z07:00",
		"2006010215:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02 15",
	}

	for _, l := range layouts {
		t, err := time.Parse(l, newDateString)
		if err != nil {
			continue
		}
		return TimeToUltipaTime(&t), err
	}

	return nil, errors.New("parse datetime string failed : " + newDateString)
}

func compensateYear(dateString string) (string, error) {
	if strings.Index(dateString, "-") > -1 {
		return doCompensateYear(dateString, "-")
	} else if strings.Index(dateString, "/") > -1 {
		return doCompensateYear(dateString, "/")
	} else if strings.Index(dateString, ":") > -1 {
		idx := strings.Index(dateString, ":")
		ymdh := dateString[0:idx]
		ymdhLength := len(ymdh)
		if ymdhLength > 10 || ymdhLength < 7 {
			return "", errors.New(fmt.Sprintf("Unable to parse %s as UltipaTime", dateString))
		} else if ymdhLength == 10 {
			return dateString, nil
		} else {
			year := ymdh[:ymdhLength-6]
			yearValue, err := strconv.Atoi(year)
			if err != nil {
				return "", err
			}
			if yearValue > 70 && yearValue < 100 {
				yearValue += 1900
			} else if yearValue < 70 {
				yearValue += 2000
			}
			return fmt.Sprintf("%d%s%s", yearValue, ymdh[ymdhLength-6:], dateString[idx:]), nil
		}
	}
	return dateString, nil
}

func doCompensateYear(dateString string, splitter string) (string, error) {
	if idx := strings.Index(dateString, splitter); idx > 0 {
		year := dateString[0:idx]
		yearValue, err := strconv.Atoi(year)
		if err != nil {
			return "", err
		}
		if yearValue > 70 && yearValue < 100 {
			yearValue += 1900
		} else if yearValue < 70 {
			yearValue += 2000
		}
		return fmt.Sprintf("%d%s", yearValue, dateString[idx:]), nil
	} else {
		return dateString, nil
	}
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

func TimeToUltipaTime(t *time.Time) *UltipaTime {
	toConvertTime := t
	if t == nil {
		defaultTime := time.Unix(0, 0)
		toConvertTime = &defaultTime
	}

	ultipaDateTime := TimeToUint64(*toConvertTime)
	ultipaTime := NewTime(ultipaDateTime)
	ultipaTime.Time = toConvertTime

	datetime := uint64(0)

	year := uint64(toConvertTime.Year())
	month := uint64(toConvertTime.Month())
	day := uint64(toConvertTime.Day())
	hour := uint64(toConvertTime.Hour())
	minute := uint64(toConvertTime.Minute())
	second := uint64(toConvertTime.Second())
	microsec := uint64(toConvertTime.Nanosecond() / 1000)

	yearMonth := year*13 + month
	datetime = yearMonth << 46
	datetime = datetime | (day << 41)
	datetime = datetime | (hour << 36)
	datetime = datetime | (minute << 30)
	datetime = datetime | (second << 24)
	datetime = datetime | microsec

	return &UltipaTime{
		Datetime: datetime,
		Year:     year,
		Month:    month,
		Day:      day,
		Hour:     hour,
		Minute:   minute,
		Second:   second,
		Macrosec: microsec,
		Time:     toConvertTime,
	}
}

func (u *UltipaTime) String() string {
	return u.Time.Format("2006-01-02T15:04:05.000Z07:00")
}

// Get Timestamp , Second
func (u *UltipaTime) GetTimeStamp() uint32 {
	return uint32(u.Time.Unix())
}
