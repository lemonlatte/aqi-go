package aqi

type AQICalculator struct {
	aqiTable          AQITable
	concentrationDiff float64
}

func NewAQICalculator(ac AQICategory) (a *AQICalculator) {
	var concentrationDiff float64
	switch ac {
	case PM2_5, CO:
		concentrationDiff = 0.1
	default:
		concentrationDiff = 1
	}
	a = &AQICalculator{
		aqiTable:          aqiTables[ac],
		concentrationDiff: concentrationDiff,
	}
	return
}

func (ac *AQICalculator) CalcAQI(current float64) float64 {
	indexList := []float64{0, 50, 100, 150, 200, 300, 400, 500}
	var lowIndex, highIndex float64
	var lowConcentration, highConcentration float64
	for _, index := range indexList {
		concentration := ac.aqiTable[index]
		if current <= concentration {
			highIndex = index
			highConcentration = concentration
			break
		}
		if index != 0 {
			lowIndex = index + 1
			lowConcentration = concentration + ac.concentrationDiff
		}
	}
	return lowIndex + (highIndex-lowIndex)*(current-lowConcentration)/(highConcentration-lowConcentration)
}

func GetAQI(ac AQICategory, current float64) float64 {
	a := NewAQICalculator(ac)
	return a.CalcAQI(current)
}
