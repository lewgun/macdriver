package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

const (
	NSAlertStyleWarning       = 0
	NSAlertStyleInformational = 1
	NSAlertStyleCritical      = 2
)

const (
	NSAlertFirstButtonReturn  = 1000
	NSAlertSecondButtonReturn = 1001
	NSAlertThirdButtonReturn  = 1002
	//...
)

type NSAlert struct {
	objc.Object
}

var nsAlert = objc.Get("NSAlert")

func NSAlert_New() NSAlert {
	return NSAlert{nsAlert.Alloc().Init()}
}

func (i NSAlert) Title() string {
	return i.Get("window").Get("title").String()
}

func (i NSAlert) SetTitle(s string) {
	i.Get("window").Send("setTitle:", core.String(s))
}

func (i NSAlert) SetMessageText(s string) {
	i.Send("setMessageText:", core.String(s))
}

func (i NSAlert) SetAlertStyle(style int) {
	i.Set("alertStyle:", style)
}

func (i NSAlert) Icon() NSImage {
	return NSImage{i.Get("icon")}
}

func (i NSAlert) SetIcon(obj NSImage) {
	i.Set("icon:", obj)
}
func (i NSAlert) AddButtonWithTitle(s string) {
	i.Set("addButtonWithTitle:", core.String(s))
}

func (i NSAlert) RunModal() int64 {
	return i.Send("runModal").Int()
}

func (i NSAlert) ShowsSuppressionButton(show bool) {
	i.Set("showsSuppressionButton:", show)
}

func (i NSAlert) SuppressionButton() NSButton {
	return NSButton{i.Get("suppressionButton")}
}
