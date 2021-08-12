package main

import (
	"RotaryEncoder/Encoder"
	"machine"
	"time"
)

// エンコーダ処理用ルーチン
// Goroutineを使用して並行処理
func checkEncoder(ch chan<- int32) {
	en := Encoder.New(machine.D5, machine.D6)
	for {
		en.Counter()                 // エンコーダカウント処理
		ch <- en.Cnt / 4             // エンコーダの値をメインルーチンへ
		time.Sleep(time.Microsecond) // 1ms待機
	}
}

func blinked(ch chan<- int32, led machine.Pin) {
	for {
		ch <- 1
		led.High()
		time.Sleep(time.Millisecond * 500)
		led.Low()
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.Low()

	led2 := machine.D4
	led2.Configure(machine.PinConfig{Mode: machine.PinOutput})

	/* Goroutine for Encoder Count Process */
	encChan1 := make(chan int32, 1) // Goroutineとの通信チャネルを作成
	go checkEncoder(encChan1)       // Goroutineの開始

	chan2 := make(chan int32, 1)
	go blinked(chan2, led2)

	/* loop */
	for {
		select {

		/* Encoder goroutine */
		case cnt := <-encChan1: // エンコーダの値を取得
			if cnt >= 10 {
				led.High()
			} else {
				led.Low()
				break
			}
		/* LED Blink process */
		case <-chan2:
			break
		}
	}
}
