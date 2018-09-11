//
// Copyright (c) 2018 ЗАО Геликон Про http://www.gelicon.biz
//
package main

import "time"

// Класс для реализации тиков в программе и периодических событий.

type Wait struct {
	mark time.Time
	min  time.Duration
	max  time.Duration
	tick time.Duration
}

func (w *Wait) Setup(tick time.Duration, bound time.Duration) {
	w.tick = tick
	w.min = -bound
	w.max = bound
}

func (w *Wait) Reset() {
	w.mark = time.Now()
}

func (w *Wait) Step() {
	w.mark = w.mark.Add(w.tick)
}

func (w *Wait) Correct() bool {
	if wait := time.Until(w.mark); wait >= w.min && wait <= w.max {
		return false
	} else {
		w.mark = time.Now()
		return true
	}
}

func (w *Wait) Wait() {
	wait := time.Until(w.mark)
	time.Sleep(wait)
}

func (w *Wait) After() bool {
	return time.Now().After(w.mark)
}
