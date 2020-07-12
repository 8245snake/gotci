package gotci

//Unit 単位
type Unit struct {
	name     UnitName
	unitType UnitType
}

//UnitType 単位の種別
type UnitType string

//UnitName 単位の名前
type UnitName string

const (
	//UnitTypeBolus ボーラス
	UnitTypeBolus UnitType = "bolus"
	//UnitTypeContinuous 持続
	UnitTypeContinuous UnitType = "continuous"
)

const (
	//UnitNameMilligram ミリグラム
	UnitNameMilligram UnitName = "mg"
	//UnitNameMicrogram マイクログラム
	UnitNameMicrogram UnitName = "μg"
	//UnitNameNanogram ナノグラム
	UnitNameNanogram UnitName = "ng"
	//UnitNameMilligramPerLiter ミリグラム/L
	UnitNameMilligramPerLiter UnitName = "mg/L"
	//UnitNameNanogramPerMilliliter ナノグラム/mL
	UnitNameNanogramPerMilliliter UnitName = "ng/mL"
)

//NewUnit 初期化
func NewUnit(name UnitName, unitType UnitType) Unit {
	return Unit{
		name:     name,
		unitType: unitType,
	}

}
