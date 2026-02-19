package main

// WaxRecommendation представляет одну рекомендацию мази
type WaxRecommendation struct {
	Name        string   // Название мази
	Style       string   // "classic", "skate", "both"
	TempMin     int      // Минимальная температура (например, -5)
	TempMax     int      // Максимальная температура (+5)
	HumidityMin int      // Минимальная влажность (0-100)
	HumidityMax int      // Максимальная влажность
	SnowTypes   []string // Типы снега: "fresh", "old", "wet", "dry"
	TrackTypes  []string // Типы трассы: "hard", "soft", "icy"
}

// База данных в памяти (пока заменим потом на PostgreSQL)
var waxDB = []WaxRecommendation{
	{
		Name:        "SWIX V30 (синий)",
		Style:       "both",
		TempMin:     -3,
		TempMax:     0,
		HumidityMin: 50,
		HumidityMax: 80,
		SnowTypes:   []string{"old", "wet"},
		TrackTypes:  []string{"hard", "icy"},
	},
	{
		Name:        "SWIX V40 (фиолетовый)",
		Style:       "classic",
		TempMin:     -5,
		TempMax:     -2,
		HumidityMin: 60,
		HumidityMax: 90,
		SnowTypes:   []string{"fresh", "old"},
		TrackTypes:  []string{"soft"},
	},
	{
		Name:        "START OSLO (жёлтый)",
		Style:       "skate",
		TempMin:     -1,
		TempMax:     3,
		HumidityMin: 40,
		HumidityMax: 70,
		SnowTypes:   []string{"wet", "dry"},
		TrackTypes:  []string{"hard"},
	},
	// Добавь свои варианты
}

// FilterWaxes подбирает мази по заданным критериям
func FilterWaxes(temp, humidity int, snow, track string, style string) []WaxRecommendation {
	var result []WaxRecommendation
	for _, wax := range waxDB {
		// Стиль
		if style != "any" && wax.Style != style && wax.Style != "both" {
			continue
		}
		// Температура
		if temp < wax.TempMin || temp > wax.TempMax {
			continue
		}
		// Влажность
		if humidity < wax.HumidityMin || humidity > wax.HumidityMax {
			continue
		}
		// Тип снега (проверяем, что снег есть в списке разрешённых)
		if !contains(wax.SnowTypes, snow) {
			continue
		}
		// Тип трассы
		if !contains(wax.TrackTypes, track) {
			continue
		}
		result = append(result, wax)
	}
	return result
}

// contains вспомогательная функция для поиска в срезе
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
