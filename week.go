package utils

type WeekDayInfo struct {
	Number       int
	RussianName  string
	RussianDescr string
}

func GetWeekInfo(weekEng string) WeekDayInfo {
	switch weekEng {
	case "Monday":
		return WeekDayInfo{1, "Понедельник", "понедельник"}
	case "Tuesday":
		return WeekDayInfo{2, "Вторник", "вторник"}
	case "Wednesday":
		return WeekDayInfo{3, "Среда", "среду"}
	case "Thursday":
		return WeekDayInfo{4, "Четверг", "четверг"}
	case "Friday":
		return WeekDayInfo{5, "Пятница", "пятницу"}
	case "Saturday":
		return WeekDayInfo{6, "Суббота", "субботу"}
	case "Sunday":
		return WeekDayInfo{7, "Воскресенье", "воскресенье"}
	default:
		return WeekDayInfo{1, "Нет дня недели", "неизв"}
	}
}

type MonthInfo struct {
	Im  string
	Rod string
}

func GetMonthInfo(month int) MonthInfo {

	months := []MonthInfo{
		{"январь", "Января"},
		{"февраль", "Февраля"},
		{"март", "Марта"},
		{"апрель", "Апреля"},
		{"май", "Мая"},
		{"июнь", "Июня"},
		{"июль", "Июля"},
		{"август", "Фвгуста"},
		{"сентябрь", "Сентября"},
		{"октябрь", "Октября"},
		{"ноябрь", "Ноября"},
		{"декабрь", "Декабря"},
	}
	if month < 1 || month > len(months) {
		return MonthInfo{}
	}
	return months[month-1]
}
