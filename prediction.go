package gotci

import "time"

//Prediction 予測データ
type Prediction struct {
	time          *time.Time
	concentration float64
}
