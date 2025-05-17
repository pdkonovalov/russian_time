package russian_time

import (
	"fmt"
	"time"
)

type RCase int

const (
	// январь
	RCase1 RCase = iota
	// января
	RCase2
	// январей
	RCase3
)

// RCase1 -> январь
//
// RCase2 -> января
var Month = func() map[time.Month]map[RCase]string {
	m := make(map[time.Month]map[RCase]string)

	m[time.March] = make(map[RCase]string)
	m[time.March][RCase1] = "март"
	m[time.March][RCase2] = "марта"

	m[time.April] = make(map[RCase]string)
	m[time.April][RCase1] = "апрель"
	m[time.April][RCase2] = "апреля"

	m[time.May] = make(map[RCase]string)
	m[time.May][RCase1] = "май"
	m[time.May][RCase2] = "мая"

	m[time.June] = make(map[RCase]string)
	m[time.June][RCase1] = "июнь"
	m[time.June][RCase2] = "июня"

	m[time.July] = make(map[RCase]string)
	m[time.July][RCase1] = "июль"
	m[time.July][RCase2] = "июля"

	m[time.August] = make(map[RCase]string)
	m[time.August][RCase1] = "август"
	m[time.August][RCase2] = "августа"

	m[time.September] = make(map[RCase]string)
	m[time.September][RCase1] = "сентябрь"
	m[time.September][RCase2] = "сентября"

	m[time.October] = make(map[RCase]string)
	m[time.October][RCase1] = "октябрь"
	m[time.October][RCase2] = "октября"

	m[time.November] = make(map[RCase]string)
	m[time.November][RCase1] = "ноябрь"
	m[time.November][RCase2] = "ноября"

	m[time.December] = make(map[RCase]string)
	m[time.December][RCase1] = "декабрь"
	m[time.December][RCase2] = "декабря"

	m[time.January] = make(map[RCase]string)
	m[time.January][RCase1] = "январь"
	m[time.January][RCase2] = "января"

	m[time.February] = make(map[RCase]string)
	m[time.February][RCase1] = "февраль"
	m[time.February][RCase2] = "февраля"

	return m
}()

// RCase1 -> минута
//
// RCase2 -> минуты
//
// RCase3 -> минут
var MinuteCount = func() map[RCase]string {
	m := make(map[RCase]string)
	m[RCase1] = "минута"
	m[RCase2] = "минуты"
	m[RCase3] = "минут"
	return m
}()

// RCase1 -> час
//
// RCase2 -> часа
//
// RCase3 -> часов
var HourCount = func() map[RCase]string {
	m := make(map[RCase]string)
	m[RCase1] = "час"
	m[RCase2] = "часа"
	m[RCase3] = "часов"
	return m
}()

// RCase1 -> день
//
// RCase2 -> дня
//
// RCase3 -> дней
var DayCount = func() map[RCase]string {
	m := make(map[RCase]string)
	m[RCase1] = "день"
	m[RCase2] = "дня"
	m[RCase3] = "дней"
	return m
}()

// RCase1 -> месяц
//
// RCase2 -> месяца
//
// RCase3 -> месяцев
var MonthCount = func() map[RCase]string {
	m := make(map[RCase]string)
	m[RCase1] = "месяц"
	m[RCase2] = "месяца"
	m[RCase3] = "месяцев"
	return m
}()

// RCase1 -> год
//
// RCase2 -> года
//
// RCase3 -> лет
var YearCount = func() map[RCase]string {
	m := make(map[RCase]string)
	m[RCase1] = "год"
	m[RCase2] = "года"
	m[RCase3] = "лет"
	return m
}()

func CountCase(n int) RCase {
	switch {
	case n%10 == 1 && n/10 != 1:
		return RCase1
	case (n%10 == 2 || n%10 == 3 || n%10 == 4) && n/10 != 1:
		return RCase2
	default:
		return RCase3
	}
}

// 15:04
func Time(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("15:04")
}

// RCase1 -> 2 январь
//
// RCase2 -> 2 января
func DayMonth(t *time.Time, c RCase) string {
	if t == nil || c == RCase3 {
		return ""
	}
	return fmt.Sprintf("%v %s", t.Day(), Month[t.Month()][c])
}

// RCase1 -> 2 январь 2006
//
// RCase2 -> 2 января 2006
func DayMonthYear(t *time.Time, c RCase) string {
	if t == nil || c == RCase3 {
		return ""
	}
	return fmt.Sprintf("%v %s %v", t.Day(), Month[t.Month()][c], t.Year())
}

// time.Now() = 02.01.2006 15:04
//
// TimeContext(02.01.2006 13:00, RCase1) = 13:00
//
// TimeContext(02.01.2006 13:00, RCase2) = 13:00
//
// TimeContext(01.01.2006 13:00, RCase1) = 1 январь
//
// TimeContext(01.01.2006 13:00, RCase2) = 1 января
//
// TimeContext(01.01.2005 13:00, RCase1) = 1 январь 2005
//
// TimeContext(01.01.2005 13:00, RCase2) = 1 января 2005
func TimeContext(t *time.Time, c RCase) string {
	if t == nil || c == RCase3 {
		return ""
	}
	y, m, d := t.Date()
	ny, nm, nd := time.Now().Date()
	if y == ny && m == nm && d == nd {
		return Time(t)
	} else if y == ny {
		return DayMonth(t, c)
	} else {
		return DayMonthYear(t, c)
	}
}

// 1 день, 2 дня, ... , 1 месяц, 2 месяца, ... , 1 год, 2 года, ...
func RoundDurationByDayToMonthToYear(start, end *time.Time) string {
	if start == nil || end == nil {
		return ""
	}
	if end.Before(*start) {
		return ""
	}

	y, m, d, _, _, _ := diff(*start, *end)
	if y == 0 && m == 0 {
		d++
		return fmt.Sprintf("%v %s", d, DayCount[CountCase(d)])
	}
	if y == 0 {
		return fmt.Sprintf("%v %s", m, MonthCount[CountCase(m)])
	}
	return fmt.Sprintf("%v %s", y, YearCount[CountCase(y)])
}

// 4 минуты
//
// 15 часов
//
// 15 часов 4 минуты
func RoundDurationByMinuteAndHour(start, end *time.Time) string {
	if start == nil || end == nil {
		return ""
	}
	if end.Before(*start) {
		return ""
	}
	minutes := int(end.Sub(*start).Milliseconds() / (60000))
	hours := minutes / 60
	minutes = minutes % 60
	if hours == 0 {
		return fmt.Sprintf("%v %s", minutes, MinuteCount[CountCase(minutes)])
	}
	return fmt.Sprintf(
		"%v %s %v %s",
		hours, HourCount[CountCase(hours)],
		minutes, MinuteCount[CountCase(minutes)])
}

func WeekDay(t *time.Time) string {
	switch t.Weekday() {
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	case time.Sunday:
		return "Воскресение"
	}
	return ""
}

func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
