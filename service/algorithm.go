package service

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"wxcloudrun-golang/app/handlers/response"
	"wxcloudrun-golang/db/model"
)

func MonthlyAvg1(grouped map[string][]model.MonthQuotation, highIndex, lowIndex, midIndex float64) ([]response.MonthlyPriceStats, error) {
	var result []response.MonthlyPriceStats
	for month, items := range grouped {
		var sumLow, sumHigh, sumPrice float64
		var count int

		for _, q := range items {
			low, err1 := strconv.ParseFloat(q.LowerPrice, 64)
			high, err2 := strconv.ParseFloat(q.HigherPrice, 64)
			if err1 != nil || err2 != nil {
				continue // 忽略格式错误的记录
			}
			sumLow += low
			sumHigh += high
			price := (low + high) / 2
			sumPrice += price
			count++
		}

		if count > 0 {
			rate1 := ((sumLow / float64(count)) / lowIndex) * 100
			rate2 := ((sumHigh / float64(count)) / highIndex) * 100
			rate3 := ((sumPrice / float64(count)) / midIndex) * 100
			result = append(result, response.MonthlyPriceStats{
				Month:     month,
				AvgLow:    sumLow / float64(count),
				AvgHigh:   sumHigh / float64(count),
				AvgPrice:  sumPrice / float64(count),
				LowIndex:  math.Round(rate1*100) / 100,
				HighIndex: math.Round(rate2*100) / 100,
				MidIndex:  math.Round(rate3*100) / 100,
			})
		}
	}

	// 可选：按月份排序（假设格式为 2025年3月）
	sort.Slice(result, func(i, j int) bool {
		return result[i].Month < result[j].Month
	})
	return result, nil
}

func MonthlyAvg2(grouped map[string][]model.MonthQuotation) ([]response.GECMonthlyPriceStats, error) {
	var result []response.GECMonthlyPriceStats

	for key, items := range grouped {
		parts := strings.SplitN(key, "|", 2)
		if len(parts) != 2 {
			continue
		}
		tp, month := parts[0], parts[1]

		var sumPrice float64
		var count int

		for _, q := range items {
			price, err := strconv.ParseFloat(q.Price, 64)
			if err != nil {
				continue
			}
			sumPrice += price
			count++
		}

		if count > 0 {
			result = append(result, response.GECMonthlyPriceStats{
				Type:     tp,
				Month:    month,
				AvgPrice: sumPrice / float64(count),
			})
		}
	}

	// 按类型 + 月份排序（可选）
	sort.Slice(result, func(i, j int) bool {
		if result[i].Type == result[j].Type {
			return result[i].Month < result[j].Month
		}
		return result[i].Type < result[j].Type
	})
	return result, nil
}
