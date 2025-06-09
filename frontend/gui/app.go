package gui

import (
	"box/internal/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func StartGUI(boxService service.BoxService) {
	myApp := app.New()
	myWindow := myApp.NewWindow("Box Manager")

	// 创建主界面
	mainContent := CreateMainUI(myWindow, boxService)
	myWindow.SetContent(mainContent)

	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
