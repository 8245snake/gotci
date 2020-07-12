package main

import (
	"time"

	"github.com/8245snake/gotci"
)

func main() {
	//速度定数(h^-1)
	const k10 float64 = 0.083
	const k12 float64 = 0.471
	const k13 float64 = 0.225
	const k21 float64 = 0.102
	const k31 float64 = 0.006
	const ke0 float64 = 0.114

	// シミュレータ作成(10秒刻みで計算)
	sec := 10
	sim := gotci.NewSimulation(sec)
	// モデル作成
	model := gotci.NewModel("test", k10, k12, k13, k21, k31, ke0, 0.105, 50.0)
	sim.SetModel(&model)
	// 現在から30分を対象とする
	start := time.Now()
	timespan := 30
	sim.SetRange(start, timespan)
	// 10mg投与
	dose := 10.0
	sim.AddInjection(time.Now(), dose, gotci.NewUnit(gotci.UnitNameMilligram, gotci.UnitTypeBolus))
	// 計算開始
	sim.Predict()
	// 結果を出力
	sim.ShowResult()
}
