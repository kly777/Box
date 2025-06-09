package gui

import (
	"box/internal/core/service"
	"box/internal/storage/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

type UIState struct {
	BoxService   service.BoxService
	Window       fyne.Window
	CurrentBoxID uint
	CurrentBox   *models.Box
	CurrentBoxes []models.Box
	currentFiles []models.File
}

func NewUIState(boxService service.BoxService, window fyne.Window) *UIState {
	return &UIState{
		BoxService: boxService,
		Window:     window,
	}
}

func (s *UIState) RefreshBoxes() {
	var newBoxes []models.Box
	var err error

	if s.CurrentBoxID == 0 {
		newBoxes, err = s.BoxService.GetRootBoxes()
	} else {
		newBoxes, err = s.BoxService.GetChildBoxes(s.CurrentBoxID)
	}

	if err != nil {
		dialog.ShowError(err, s.Window)
		return
	}

	if s.CurrentBoxID != 0 {
		newBoxes = append([]models.Box{{Name: ".."}}, newBoxes...)
	}

	s.CurrentBoxes = newBoxes
}

func (s *UIState) RefreshFiles() {
	if s.CurrentBoxID == 0 {
		s.currentFiles = nil
		return
	}

	files, err := s.BoxService.GetFilesInBox(s.CurrentBoxID)
	if err != nil {
		dialog.ShowError(err, s.Window)
		return
	}
	s.currentFiles = files
}

func (s *UIState) RefreshAll() {
	s.RefreshBoxes()
	s.RefreshFiles()
}

func (s *UIState) GetCurrentFiles() []models.File {
	return s.currentFiles
}
