package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func PatchNumEditWgs(ins []Widget, str string, edit **walk.NumberEdit) []Widget {
	wgs := GenNumEditWgs(str, edit)
	ins = append(ins, wgs...) //append返回的是新的
	return ins
}

func GenNumEditWgs(str string, edit **walk.NumberEdit) []Widget {
	w11 := Label{
		Text: str,
	}
	w12 := GenNumberEdit(edit)
	return []Widget{w11, w12}
}

func GenNumberEdit(edit **walk.NumberEdit) *NumberEdit {
	out := &NumberEdit{
		Decimals: 3,
		AssignTo: edit,
		// OnMouseDown: clearEdit,
	}
	return out
}

// // ctrl 清数据  需要根据xy坐标来确定是哪个edit要清空。好恶心
// func clearEdit(x, y int, button walk.MouseButton) {
// 	if key == walk.KeyControl {
// 		ClearInputNums()
// 		mainWin.MainWindow.SetFocus() //这样就可以调用主界面的快捷键
// 	}
// }
