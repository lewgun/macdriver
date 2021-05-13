package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSAlert struct {
	objc.Object
}

var nsAlert = objc.Get("NSAlert")

func NSAlert_New() NSAlert {
	return NSAlert{nsAlert.Alloc().Init()}
}


//
//func (i NSAlert) Title() string {
//	return i.Get("title").String()
//}

func (i NSAlert) SetTitle(s string) {
	i.Get("window").Send("setTitle:", core.String(s))
}

func (i NSAlert) SetMessageText(s string) {
	i.Send("setMessageText:" ,core.String(s))
}


func (i NSAlert) Icon() NSImage {
	return NSImage{i.Get("icon")}
}

func (i NSAlert) SetIcon(obj NSImage) {
	i.Set("icon:", obj)
}
func (i NSAlert)AddButtonWithTitle(s string) {
	i.Set("addButtonWithTitle:" ,core.String(s))
}

func (i NSAlert) RunModal() {
	i.Send("runModal")
}