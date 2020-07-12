package gotci

import (
	"fmt"
	"time"
)

//Simulation シミュレーション
type Simulation struct {
	//モデル
	model *Model
	//時間間隔(s)
	timestep float64
	//シミュレーション開始時刻
	startTime *time.Time
	//シミュレーション終了時刻
	endTime *time.Time
	//刻み幅
	h float64
	//投与
	injections []Injection
	//血中濃度
	c1 []Prediction
	//効果部位濃度
	ce []Prediction
}

//NewSimulation 初期化
func NewSimulation(timestep int) Simulation {
	h := 60.0 / timestep
	start := time.Now()
	end := start.Add(2 * time.Hour)
	sim := Simulation{
		injections: []Injection{},
		c1:         []Prediction{},
		ce:         []Prediction{},
		timestep:   float64(timestep),
		h:          float64(h),
		startTime:  &start,
		endTime:    &end,
	}
	return sim
}

//SetModel モデルを設定
func (sim *Simulation) SetModel(model *Model) *Simulation {
	sim.model = model
	sim.model.k10 = sim.model.k10 / sim.h
	sim.model.k12 = sim.model.k12 / sim.h
	sim.model.k13 = sim.model.k13 / sim.h
	sim.model.k21 = sim.model.k21 / sim.h
	sim.model.k31 = sim.model.k31 / sim.h
	sim.model.ke0 = sim.model.ke0 / sim.h
	return sim
}

//SetRange 時間設定(startを含む時刻からminute分だけ指定)
func (sim *Simulation) SetRange(start time.Time, minutes int) *Simulation {
	//開始時刻（秒数を0に）
	date := time.Date(start.Year(), start.Month(), start.Day(), start.Hour(), start.Minute(), 0, 0, time.Local)
	sim.startTime = &date
	//終了時刻
	end := date.Add(time.Duration(minutes) * time.Minute)
	sim.endTime = &end
	//空の投与量を作成
	inTime := date
	for {
		injection := NewInjection(inTime, 0.0, nil)
		sim.injections = append(sim.injections, injection)
		inTime = inTime.Add(time.Duration(sim.timestep) * time.Second)
		if inTime.After(end) {
			break
		}
	}
	return sim
}

//Predict シミュレーションする
func (sim *Simulation) Predict() {
	//濃度(mg/L)
	c1 := 0.0
	c2 := 0.0
	c3 := 0.0
	ce := 0.0

	for _, inject := range sim.injections {
		bolus, continuous := inject.getAmount(sim.h)
		c1, c2, c3, ce = sim.model.rungeKutta(c1, c2, c3, ce, bolus, continuous, sim.h)
		datetime := time.Date(inject.time.Year(), inject.time.Month(), inject.time.Day(), inject.time.Hour(), inject.time.Minute(), inject.time.Second(), 0, time.Local)
		sim.c1 = append(sim.c1, Prediction{time: &datetime, concentration: c1 / (sim.model.v1 * 1000)})
		sim.ce = append(sim.ce, Prediction{time: &datetime, concentration: ce / (sim.model.v1 * 1000)})
	}
}

//AddInjection 投与
func (sim *Simulation) AddInjection(inTime time.Time, amount float64, unit Unit) *Simulation {
	//刻み時間に丸める
	sec := inTime.Minute()*60 + inTime.Second()
	sec = sec - sec%int(sim.timestep)
	inTime = time.Date(inTime.Year(), inTime.Month(), inTime.Day(), inTime.Hour(), 0, 0, 0, time.Local)
	inTime = inTime.Add(time.Duration(sec) * time.Second)
	additional := NewInjection(inTime, amount, &unit)

	//injectionsに加える
	for i, inject := range sim.injections {
		if inject.time.Equal(inTime) {
			switch unit.unitType {
			case UnitTypeBolus:
				sim.injections[i] = additional
			case UnitTypeContinuous:
			}
		}
	}
	return sim
}

//ShowResult 結果を表示
func (sim *Simulation) ShowResult() {
	for i, c1 := range sim.c1 {
		fmt.Println(c1.time.Format("15:04:05"), c1.concentration, sim.ce[i].concentration)
	}
}
