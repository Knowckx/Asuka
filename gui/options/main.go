package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MainWin struct {
	*walk.MainWindow
	AssetsNow      *walk.NumberEdit
	AssetsPredict  *walk.NumberEdit
	OptionsNow     *walk.NumberEdit
	OptionsLever   *walk.NumberEdit
	OptionsPredict *walk.TextEdit
}

var mainWin = new(MainWin)

func main() {

	wgs := []Widget{}
	wgs = append(wgs, NumberEditors()...)
	wgs = append(wgs, NewButton())
	wgs = append(wgs, NewTextEdit(&mainWin.OptionsPredict))

	err := GenMainWin(wgs)
	if err != nil {
		log.Fatal(err)
	}
}

func GenMainWin(wgs []Widget) error {
	win := MainWindow{
		AssignTo: &mainWin.MainWindow,
		Title:    "Options cala",
		MinSize:  Size{300, 500},
		Layout:   VBox{},
		Children: wgs,
		OnKeyDown: func(key walk.Key) { // 无效 - 听说是因为这个窗口没有成为焦点 va 它永远被其他组件所遮
			fmt.Println("GenMainWin key is --", key)
			if key == walk.KeyReturn {
			}
		},
	}
	_, err := win.Run()
	return err

}

// 输入 input 4个条件数字
func NumberEditors() []Widget {
	ins := []Widget{}
	ins = PatchNumEditWgs(ins, "Assets Now:", &mainWin.AssetsNow)
	ins = PatchNumEditWgs(ins, "Assets Predict:", &mainWin.AssetsPredict)
	ins = PatchNumEditWgs(ins, "Options Now:", &mainWin.OptionsNow)
	ins = PatchNumEditWgs(ins, "Options Lever:", &mainWin.OptionsLever)
	return ins
}

func NewButton() PushButton {
	btn := PushButton{}
	btn.Text = "Cal"
	btn.OnClicked = MainCalc
	return btn
}

func NewTextEdit(te **walk.TextEdit) TextEdit {
	out := TextEdit{
		AssignTo:  te,
		OnKeyDown: KeyClearNum,
	}
	return out
}

// 开始计算
func MainCalc() {
	err := CheckInput()
	if err != nil {
		return
	}
	// 100 --> 90   rate = -10%
	rate := mainWin.AssetsPredict.Value()/mainWin.AssetsNow.Value() - 1
	// level -5 : -10% * -5 = 50%
	reRate := rate * mainWin.OptionsLever.Value()
	rst := mainWin.OptionsNow.Value() * (1.0 + reRate)
	// format
	sres := fmt.Sprintf("%.3f", rst)
	mainWin.OptionsPredict.SetText(sres)
}

func CheckInput() error {
	err := fmt.Errorf("input find empty")
	if mainWin.AssetsPredict.Value() == 0 || mainWin.AssetsNow.Value() == 0 {
		return err
	}

	if mainWin.OptionsNow.Value() == 0 || mainWin.OptionsLever.Value() == 0 {
		return err
	}
	return nil
}

func ClearInputNums() {
	mainWin.AssetsPredict.SetValue(0.0)
	mainWin.AssetsNow.SetValue(0.0)
	mainWin.OptionsNow.SetValue(0.0)
	mainWin.OptionsLever.SetValue(0.0)
	mainWin.OptionsPredict.SetText("")
}

// ctrl 清数据
func KeyClearNum(key walk.Key) {
	if key == walk.KeyControl {
		ClearInputNums()
		mainWin.MainWindow.SetFocus() //这样就可以调用主界面的快捷键
	}
}
