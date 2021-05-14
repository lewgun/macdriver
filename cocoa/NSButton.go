package cocoa

import (
	"github.com/lewgun/macdriver/objc"
	"github.com/progrium/macdriver/core"
)

type NSButton struct {
	NSView
}

func (i NSButton) State() int64 {
	return i.Get("state").Int()
}
func (i NSButton) SetTarget(obj objc.Object) {
	i.Set("target:", obj)
}

func (i NSButton) SetAction(sel objc.Selector) {
	i.Set("action:", obj)
}

func (i NSButton) Title() string {
	return i.Get("title").String()
}

func (i NSButton) SetTitle(s string) {
	i.Set("title:", core.String(s))
}
