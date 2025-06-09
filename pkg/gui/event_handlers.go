package gui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func SetupEventHandlers(components *UIComponents, state *UIState, window fyne.Window) {
	// Box列表点击事件处理
	components.BoxList.OnSelected = func(id widget.ListItemID) {
		if id >= len(state.CurrentBoxes) {
			return
		}

		box := state.CurrentBoxes[id]
		if box.Name == ".." {
			log.Printf("返回上级目录%v", box)
			if state.CurrentBoxID != 0 {
				parentBox, err := state.BoxService.GetBoxByID(state.CurrentBoxID)
				log.Printf("获取父级Box%v", parentBox)
				if err == nil && parentBox.ParentID != nil {
					state.CurrentBoxID = uint(*parentBox.ParentID)
				} else {
					state.CurrentBoxID = 0
				}
			}
		} else {
			state.CurrentBoxID = uint(box.ID)
		}
		components.BoxList.Unselect(id)

		state.RefreshAll()
		components.BoxList.Refresh()
		components.FileList.Refresh()
	}

	// 添加Box按钮点击事件
	components.CreateBoxBtn.OnTapped = func() {
		nameEntry := widget.NewEntry()
		formDialog := dialog.NewForm(
			"新建Box",
			"创建",
			"取消",
			[]*widget.FormItem{
				{Text: "输入Box名称：", Widget: nameEntry},
			},
			func(submitted bool) {
				if !submitted {
					return
				}
				name := nameEntry.Text
				if name != "" {
					_, err := state.BoxService.CreateBox(name)
					if err != nil {
						dialog.ShowError(err, window)
					} else {
						state.RefreshAll()
						components.BoxList.Refresh()
					}
				}
			},
			window)
		formDialog.Show()
	}
}
