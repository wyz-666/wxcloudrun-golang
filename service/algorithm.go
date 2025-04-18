package service

import (
	"math"
	"regexp"
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

func AddFitPriceToStats(data []response.MonthlyPriceStats) []response.MonthlyPriceStats {
	if len(data) == 0 {
		return nil
	}

	baseMonth := parseMonthNumber(data[0].Month)

	// Step 1: 拿出 X 和 Y 值
	var X, Y []float64
	for _, d := range data {
		relativeMonth := parseMonthNumber(d.Month) - baseMonth + 1
		X = append(X, float64(relativeMonth))
		Y = append(Y, d.AvgPrice)
	}

	// Step 2: 拟合线性函数
	n := float64(len(X))
	var sumX, sumY, sumXY, sumXX float64
	for i := range X {
		sumX += X[i]
		sumY += Y[i]
		sumXY += X[i] * Y[i]
		sumXX += X[i] * X[i]
	}
	a := (n*sumXY - sumX*sumY) / (n*sumXX - sumX*sumX)
	b := (sumY - a*sumX) / n

	// Step 3: 写入每项的拟合值
	for i := range data {
		x := X[i]
		fit := a*x + b
		data[i].FitPrice = math.Round(fit*100) / 100 // 保留两位小数
	}

	return data
}

// 辅助函数：解析 "2025年3月" → 202503
func parseMonthNumber(monthStr string) int {
	re := regexp.MustCompile(`(\\d{4})年(\\d{1,2})月`)
	matches := re.FindStringSubmatch(monthStr)
	if len(matches) == 3 {
		year, _ := strconv.Atoi(matches[1])
		month, _ := strconv.Atoi(matches[2])
		return year*100 + month
	}
	return 0
}
