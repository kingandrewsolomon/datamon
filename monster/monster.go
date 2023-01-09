package monster

import "math/rand"

type Datamon struct {
	Name          string
	Age           int
	Is_fainted    bool
	Times_fainted int
	Health        int
	Hunger        int
	Strength      float32
	Stamina       float32
	Speed         int
}

func NewDatamon() *Datamon {
	return nil
}

func (d *Datamon) BreedDatamon() *Datamon {
	return nil
}

func (d *Datamon) Attack(target *Datamon) {
	damage := int(d.Strength * d.Stamina * 10 * (rand.Float32() + 0.5))
	if damage < 1 {
		damage = 1
	}
	target.Health -= damage
}

func (d *Datamon) Grow() {
	d.Age += 1
}

func (d *Datamon) CheckStatus() bool {
	return !d.Is_fainted
}
