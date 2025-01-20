package conv

func FarenheitToKelvin(f float64) float64 {
	return ((f - 32) * 5 / 9) + 273.15
}
