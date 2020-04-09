package fynegui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

//ButtonEntry Widget
type ButtonEntry struct {
	widget.Entry
	confirmButton *widget.Button
}

func (e *ButtonEntry) onKey() {
	e.confirmButton.OnTapped()
}

//KeyDown func
func (e *ButtonEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn:
		e.onKey()
	default:
		e.Entry.KeyDown(key)
	}
}

//SetConfirmButton func
func (e *ButtonEntry) SetConfirmButton(b *widget.Button) {
	e.confirmButton = b
}

//NewButtonEntry func
func NewButtonEntry() *ButtonEntry {
	entry := &ButtonEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}
