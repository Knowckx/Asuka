// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var te1 *walk.NumberEdit
var te2 *walk.NumberEdit
var te3 *walk.TextEdit

func main() {

	wgs := []Widget{}
	wgs = append(wgs, NewNumberEditor()...)
	wgs = append(wgs, NewButton())
	wgs = append(wgs, NewTextEdit(&te3))

	if _, err := (MainWindow{
		Title:    "Walk Clipboard Example",
		MinSize:  Size{300, 300},
		Layout:   VBox{},
		Children: wgs,
	}).Run(); err != nil {
		log.Fatal(err)
	}
}

func NewButton() PushButton {
	btn := PushButton{}
	btn.Text = "Cal"
	btnFunc := MainCalc
	btn.OnClicked = btnFunc
	return btn

}

func NewTextEdit(te **walk.TextEdit) TextEdit {
	out := TextEdit{
		AssignTo: te,
	}
	return out
}

func NewNumberEditor() []Widget {
	w1 := Label{
		Text: "Weight:",
	}
	w2 := NumberEdit{
		Suffix:         " g",
		Decimals:       2,
		AssignTo:       &te1,
		OnValueChanged: walk.EventHandler(MainCalc),
		OnKeyPress:     walk.KeyEventHandler(ClearNums),
	}

	w3 := Label{
		Text: "Price:",
	}
	w4 := NumberEdit{
		Suffix:         " y",
		Decimals:       2,
		AssignTo:       &te2,
		OnValueChanged: walk.EventHandler(MainCalc),
		OnKeyPress:     walk.KeyEventHandler(ClearNums),
	}
	return []Widget{w1, w2, w3, w4}
}

type Cal struct {
	Name   string
	Weight float64
	Price  float64
	Res    float64
}

func BeKey(key walk.Key) {
	if key == walk.KeySpace {
		MainCalc()
	}
}

func MainCalc() {
	if te1.Value() == 0 || te2.Value() == 0 {
		return
	}
	rst := 500 / te1.Value() * te2.Value()
	sres := fmt.Sprintf("%.2f", rst)
	log.Print(sres)
	te3.SetText(sres)
	// te3.SetValue(rst)
}
func ClearNums(key walk.Key) {
	log.Print(key)
	if key == walk.KeySpace {
		te1.SetValue(0.0)
		te2.SetValue(0.0)
	}
}
