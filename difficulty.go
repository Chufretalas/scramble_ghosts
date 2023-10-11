package main

import (
	"fmt"
	"time"
)

func NewDefaultDifficulty() Difficulty {
	return Difficulty{Level: 0, EnemiesPerSpawn: 1, EnemySpawnDelay: time.Millisecond * 120, EnemySpeedMult: 1, DWSpawnChance: 550, DWSpeedMult: 1, ShouldIncrease: true, ArcshotDelay: time.Second * 7, ArcshotShots: 4}
}

type Difficulty struct {
	Level           int // this is what keeps track of the current difficulty level
	EnemiesPerSpawn int
	EnemySpawnDelay time.Duration
	EnemySpeedMult  float64
	DWSpawnChance   int32 // the higher the number, the rarer the spawn
	DWSpeedMult     float64
	ShouldIncrease  bool
	ArcshotDelay    time.Duration // how long between spawns //TODO: increase this
	ArcshotShots    int           //how many times arcshot shoots in one pass through the screen //TODO: increase this, maybe
}

func (d *Difficulty) Reset() {
	d.Level = 0
	d.EnemiesPerSpawn = 1
	d.EnemySpawnDelay = time.Millisecond * 120
	d.EnemySpeedMult = 1
	d.DWSpawnChance = 550
	d.DWSpeedMult = 1
	d.ShouldIncrease = true
	d.ArcshotDelay = time.Second * 7
	d.ArcshotShots = 4
}

func (d *Difficulty) Increase() {
	d.Level++

	if d.Level%12 == 0 && d.EnemiesPerSpawn < 3 {
		d.EnemiesPerSpawn++
		return
	}

	if d.Level%5 == 0 && d.EnemySpeedMult < 1.3 {
		d.EnemySpeedMult += .05
		return
	}

	if d.Level%6 == 0 && d.DWSpeedMult < 1.3 {
		d.DWSpeedMult += .05
		return
	}

	if d.Level%8 == 0 && d.DWSpawnChance > 400 {
		d.DWSpawnChance -= 25
		return
	}

	if d.EnemySpawnDelay > time.Millisecond*60 {
		d.EnemySpawnDelay -= time.Duration(float32(d.EnemySpawnDelay) * 0.05)
	}
}

func (d Difficulty) String() string {
	return fmt.Sprintf("Lvl: %v | EPS: %v\nESD: %v | ESM: %v\nDWSC: %v | DWSM: %v", d.Level, d.EnemiesPerSpawn, d.EnemySpawnDelay, d.EnemySpeedMult, d.DWSpawnChance, d.DWSpeedMult)
}
