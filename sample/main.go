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

	//体重
	const weight float64 = 50.0

	//分布容積
	const v1 float64 = 0.105

	// シミュレータ作成(10秒刻みで計算)
	sim := gotci.NewSimulation(10)
	// モデル作成
	model := gotci.NewModel("プロポフォール", k10, k12, k13, k21, k31, ke0, v1, weight)
	sim.SetModel(&model)
	// 現在から30分を対象とする
	start := time.Now()
	sim.InitRange(start, 30)
	// ボーラス20mg投与
	sim.Bolus(&start, 20.0, gotci.NewUnit(gotci.UnitNameMilligram, gotci.UnitTypeBolus))
	// 持続10mg投与(conc. 50 μg/mL)
	sim.Continuous(&start, nil, 10.0, gotci.NewUnit(gotci.UnitNameMicrogramPerMilliliter, gotci.UnitTypeContinuous), 50.0)
	// 計算開始
	sim.Predict()
	// 結果を出力
	sim.ShowResult()
}
