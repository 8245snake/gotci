package gotci

//Model モデル
type Model struct {
	name string
	//速度定数(h^-1)
	k10 float64
	k12 float64
	k13 float64
	k21 float64
	k31 float64
	ke0 float64
	//体重(kg)
	weight float64
	//分布容積(L)
	v1 float64
	v2 float64
	v3 float64
}

//NewModel モデル作成
func NewModel(name string, k10, k12, k13, k21, k31, ke0, _v1, weight float64) Model {
	v1 := _v1 * weight
	return Model{
		name:   name,
		k10:    k10,
		k12:    k12,
		k13:    k13,
		k21:    k21,
		k31:    k31,
		ke0:    ke0,
		weight: weight,
		v1:     v1,
		v2:     k12 * v1 / k21,
		v3:     k13 * v1 / k31,
	}
}

//中央区画の速度変化
func (m *Model) deltaC1(c1, c2, c3, h, bolus, cont float64) float64 {
	return (-c1*(m.k10+m.k12+m.k13) + c2*m.k21 + c3*m.k31) + cont/h
}

//抹消区画(fast)の速度変化
func (m *Model) deltaC2(c1, c2 float64) float64 {
	return (c1*m.k12 - c2*m.k21)
}

//抹消区画(slow)の速度変化
func (m *Model) deltaC3(c1, c3 float64) float64 {
	return (c1*m.k13 - c3*m.k31)
}

//効果部位の速度変化
func (m *Model) deltaCe(c1, ce float64) float64 {
	return (c1*m.ke0 - ce*m.ke0)
}

//rungeKutta ルンゲクッタ法でh(hour)だけ未来の値を予測する
func (m *Model) rungeKutta(c1, c2, c3, ce, bolus, cont, h float64) (float64, float64, float64, float64) {
	d1c1 := m.deltaC1(c1, c2, c3, h, bolus, cont)
	d1c2 := m.deltaC2(c1, c2)
	d1c3 := m.deltaC3(c1, c3)
	d1ce := m.deltaCe(c1, ce)

	d2c1 := m.deltaC1(c1+d1c1*0.5, c2+d1c2*0.5, c3+d1c3*0.5, h, bolus, cont)
	d2c2 := m.deltaC2(c1+d1c1*0.5, c2+d1c2*0.5)
	d2c3 := m.deltaC3(c1+d1c1*0.5, c3+d1c3*0.5)
	d2ce := m.deltaCe(c1+d1c1*0.5, ce+d1ce*0.5)

	d3c1 := m.deltaC1(c1+d2c1*0.5, c2+d2c2*0.5, c3+d2c3*0.5, h, bolus, cont)
	d3c2 := m.deltaC2(c1+d2c1*0.5, c2+d2c2*0.5)
	d3c3 := m.deltaC3(c1+d2c1*0.5, c3+d2c3*0.5)
	d3ce := m.deltaCe(c1+d2c1*0.5, ce+d2ce*0.5)

	d4c1 := m.deltaC1(c1+d3c1, c2+d3c2, c3+d3c3, h, bolus, cont)
	d4c2 := m.deltaC2(c1+d3c1, c2+d3c2)
	d4c3 := m.deltaC3(c1+d3c1, c3+d3c3)
	d4ce := m.deltaCe(c1+d3c1, ce+d3ce)

	c1 += (d1c1+2*d2c1+2*d3c1+d4c1)/6.0 + bolus/h
	c2 += (d1c2 + 2*d2c2 + 2*d3c2 + d4c2) / 6.0
	c3 += (d1c3 + 2*d2c3 + 2*d3c3 + d4c3) / 6.0
	ce += (d1ce + 2*d2ce + 2*d3ce + d4ce) / 6.0

	return c1, c2, c3, ce
}
