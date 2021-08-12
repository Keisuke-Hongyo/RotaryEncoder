package Encoder

import (
	"machine"
)

// エンコーダのパターンによる増減
var pattern []int32 = []int32{0, -1, 1, 0, 1, 0, 0, -1, -1, 0, 0, 1, 0, 1, -1, 0}

type Encoder struct {
	phaseA machine.Pin // エンコーダA相の入力端子
	phaseB machine.Pin // エンコーダB相の入力端子
	oldEnc int32       // エンコーダパターン格納変数(過去)
	newEnc int32       // エンコーダパターン格納変数(現在)
	Cnt    int32       // エンコーダのカウント数
}

// New : エンコーダ初期設定 */
func New(pinA machine.Pin, pinB machine.Pin) Encoder {

	en := Encoder{}

	en.phaseA = pinA
	en.phaseB = pinB
	en.phaseA.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	en.phaseB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	/* 初期の設定*/
	en.newEnc = 0
	en.oldEnc = 0
	en.Cnt = 0
	if en.phaseA.Get() {
		en.oldEnc += 2
	}

	if en.phaseB.Get() {
		en.oldEnc += 1
	}
	/* 初期設定用の構造体*/
	return en
}

// Counter : エンコーダのカウント処理
func (en *Encoder) Counter() {
	en.newEnc = 0
	if en.phaseA.Get() {
		en.newEnc += 2
	}

	if en.phaseB.Get() {
		en.newEnc += 1
	}

	enc := en.oldEnc*4 + en.newEnc

	en.oldEnc = en.newEnc

	en.Cnt += pattern[enc]
}
