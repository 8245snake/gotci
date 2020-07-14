package gotci

import "time"

//Injection 投与データ
type Injection struct {
	time       time.Time
	bolus      float64
	continuous float64
}

//NewInjection 初期化
func NewInjection(time time.Time, bolus float64, continuous float64) Injection {
	return Injection{
		time:       time,
		continuous: continuous,
		bolus:      bolus,
	}
}
