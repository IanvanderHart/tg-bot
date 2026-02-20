package main

import (
	"context"
	"fmt"
)

// WaxRecommendation структура мази (оставляем как есть)
type WaxRecommendation struct {
	Name        string
	Style       string
	TempMin     int
	TempMax     int
	HumidityMin int
	HumidityMax int
	SnowTypes   []string
	TrackTypes  []string
}

// FilterWaxes выполняет поиск по базе данных
func FilterWaxes(temp, humidity int, snow, track string, style string) ([]WaxRecommendation, error) {
	if dbPool == nil {
		return nil, fmt.Errorf("database not connected")
	}

	// SQL-запрос с параметрами
	query := `
		SELECT name, style, temp_min, temp_max, humidity_min, humidity_max, snow_types, track_types
		FROM waxes
		WHERE $1 BETWEEN temp_min AND temp_max
		  AND $2 BETWEEN humidity_min AND humidity_max
		  AND ($3 = ANY(snow_types) OR $3 IS NULL)
		  AND ($4 = ANY(track_types) OR $4 IS NULL)
		  AND (style = $5 OR style = 'both')
	`

	rows, err := dbPool.Query(context.Background(), query, temp, humidity, snow, track, style)
	if err != nil {
		return nil, fmt.Errorf("database query error: %w", err)
	}
	defer rows.Close()

	var results []WaxRecommendation
	for rows.Next() {
		var w WaxRecommendation
		err := rows.Scan(&w.Name, &w.Style, &w.TempMin, &w.TempMax, &w.HumidityMin, &w.HumidityMax, &w.SnowTypes, &w.TrackTypes)
		if err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		results = append(results, w)
	}
	return results, nil
}
