package color

func GetRgbas() [5]string {
	return [5]string{
		"rgba(255,165,0,1)", // orange
		"rgba(255,0,0,1)",   // red
		"rgba(124,252,0,1)", // lawngreen
		"rgba(0,255,255,1)", // cyan
		"rgba(0,0,255,1)",   // blue
	}
}

func GetRgbaByYear(year int64) string {
	switch year {
	case 2020, 2025, 2030:
		return "rgba(255,165,0,1)" // orange
	case 2021, 2026, 2031:
		return "rgba(124,252,0,1)" // lawngreen
	case 2022, 2027, 2032:
		return "rgba(0,0,255,1)" // blue
	case 2023, 2028, 2033:
		return "rgba(255,0,0,1)" // red
	case 2024, 2029, 2034:
		return "rgba(0,255,255,1)" // cyan
	default:
		return "rgba(0,0,0,1)"
	}
}
