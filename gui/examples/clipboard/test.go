// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main2() {
	var te1 *walk.TextEdit
	var te2 *walk.TextEdit

	// win := MainWindow{}
	// win.Title = "123"

	wgs := []Widget{NewWidget(te1), NewWidget(te2), NewButton()}

	if _, err := (MainWindow{
		Title:    "Walk Clipboard Example",
		MinSize:  Size{600, 400},
		Layout:   VBox{},
		Children: wgs,
	}).Run(); err != nil {
		log.Fatal(err)
	}
}

func NewButton() PushButton {
	btn := PushButton{}
	btn.Text = "Cal"
	btnFunc := func() {
		log.Print("Hello")
	}
	btn.OnClicked = btnFunc
	return btn

}

func NewWidget(te *walk.TextEdit) TextEdit {
	out := TextEdit{
		AssignTo: &te,
	}
	return out
}
