package transform

import (
	"fmt"
	db "server/db/sqlc"
	"sort"
	"strconv"
	"time"
)

type FilteredDatasRow map[string][]db.GetDatasRow

type ByDate []db.GetDatasRow

func (d ByDate) Len() int           { return len(d) }
func (d ByDate) Less(i, j int) bool { return d[i].CreatedAt.Time.Before(d[j].CreatedAt.Time) }
func (d ByDate) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func sortByDate(data []db.GetDatasRow) (d []db.GetDatasRow) {
	byDate := ByDate(data)
	sort.Sort(byDate)
	return byDate
}

func getHourlyKey(time time.Time) (key string) {
	return fmt.Sprintf("%v-%v-%v-%v", time.Year(), time.Month(), time.Day(), time.Hour())
}

func getDailyKey(time time.Time) (key string) {
	return fmt.Sprintf("%v-%v-%v", time.Year(), time.Month(), time.Day())
}

func getMonthlyKey(time time.Time) (key string) {
	return fmt.Sprintf("%v-%v", time.Year(), time.Month())
}

func getYearlyKey(time time.Time) (key string) {
	return fmt.Sprintf("%v", time.Year())
}

func getAverage(data []db.GetDatasRow, getKey func(time time.Time) (key string)) (d []db.GetDatasRow) {
	var averageData []db.GetDatasRow
	filteredDatasRow := make(FilteredDatasRow)

	for _, d := range data {
		time := getKey(d.CreatedAt.Time)
		if _, ok := filteredDatasRow[time]; ok {
			filteredDatasRow[time] = append(filteredDatasRow[time], d)
		} else {
			filteredDatasRow[time] = []db.GetDatasRow{
				d,
			}
		}

	}

	for _, f := range filteredDatasRow {
		var average float64 = 0.0
		var dateRow = f[0]
		for _, v := range f {
			average += v.Value
		}

		twoDecFloat, err := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(average)/float64(len(f))), 64)

		if err != nil {
			fmt.Println("error")
			continue
		}

		dateRow.Value = twoDecFloat
		averageData = append(averageData, dateRow)
	}

	return averageData
}

func GetHourlyAverageData(data []db.GetDatasRow) (d []db.GetDatasRow) {
	return sortByDate(getAverage(data, getHourlyKey))
}

func GetDailyAverageData(data []db.GetDatasRow) (d []db.GetDatasRow) {
	return sortByDate(getAverage(data, getDailyKey))
}

func GetMonthlyAverageData(data []db.GetDatasRow) (d []db.GetDatasRow) {
	return sortByDate(getAverage(data, getMonthlyKey))
}

func GetYearlyAverageData(data []db.GetDatasRow) (d []db.GetDatasRow) {
	return sortByDate(getAverage(data, getYearlyKey))
}
