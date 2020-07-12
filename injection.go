package gotci

import "time"

//Injection 投与データ
type Injection struct {
	time   time.Time
	amount float64
	unit   *Unit
}

//NewInjection 初期化
func NewInjection(time time.Time, amount float64, unit *Unit) Injection {
	return Injection{
		time:   time,
		amount: amount,
		unit:   unit,
	}
}

//単位から判断して単位時間あたりの投与量を返す
func (in Injection) getAmount(h float64) (bolus float64, continuous float64) {
	if in.unit == nil {
		return
	}
	switch in.unit.unitType {
	case UnitTypeBolus:
		switch in.unit.name {
		case UnitNameMilligram:
			bolus = in.amount * 1000.0 * h
		}
	case UnitTypeContinuous:
		switch in.unit.name {
		case UnitNameMilligramPerLiter:
			bolus = in.amount * 1000.0 * h
		}
	}

	return
}
