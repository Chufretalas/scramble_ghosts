package main

import "time"

func NewDefaultDifficulty() Difficulty {
	return Difficulty{Level: 0, EnemiesPerSpawn: 1, EnemySpawnDelay: time.Millisecond * 120, EnemySpeedMult: 1, DWSpawnChance: 550, DWSpeedMult: 1, ShouldIncrease: true}
}

type Difficulty struct {
	Level           int // this is what keeps track of the current difficulty level
	EnemiesPerSpawn int
	EnemySpawnDelay time.Duration
	EnemySpeedMult  float64 //TODO: impl this
	DWSpawnChance   int     // the higher the number, the rarer the spawn //TODO: impl this
	DWSpeedMult     float64 //TODO: impl this
	ShouldIncrease  bool    //TODO: impl this
}

func (d *Difficulty) Reset() {
	d.Level = 0
	d.EnemiesPerSpawn = 1
	d.EnemySpawnDelay = time.Millisecond * 120
	d.EnemySpeedMult = 1
	d.DWSpawnChance = 550
	d.DWSpeedMult = 1
	d.ShouldIncrease = true
}

func (d *Difficulty) Increase() {
	d.Level++

	if d.Level%10 == 0 && d.EnemiesPerSpawn < 3 {
		d.EnemiesPerSpawn++
		return
	}

	if d.EnemySpawnDelay > time.Millisecond*40 {
		d.EnemySpawnDelay -= time.Duration(float32(d.EnemySpawnDelay) * 0.05)
	}
}
