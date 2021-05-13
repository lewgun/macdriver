package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSMenu struct {
	objc.Object
}

var nsMenu = objc.Get("NSMenu")

func NSMenu_New() NSMenu {
	return NSMenu{nsMenu.Alloc().Init()}
}

func NSMenu_Init(title string) NSMenu {
	return NSMenu{nsMenu.Alloc().Send("initWithTitle:", core.String(title))}
}

func (menu NSMenu) SetTitle(title string) {
	menu.Send("setTitle:", core.String(title))
}

func (menu NSMenu) Title() string {
	return menu.Send("title").String()
}

func (menu NSMenu) AddItem(item NSMenuItem) {
	menu.Send("addItem:", item)
}

func (menu NSMenu) RemoveItem(item NSMenuItem) {
	menu.Send("removeItem:", item)
}

func (menu NSMenu) SetAutoenablesItems(b bool) {
	menu.Set("autoenablesItems:", b)
}

func (menu NSMenu) AutoenablesItems() bool {
	return menu.Get("autoenablesItems").Bool()
}

func (menu NSMenu) ItemArray() []NSMenuItem {
	obj := menu.Get("itemArray")
	a := core.NSArray{obj}
	count := int(a.Count())
	items := make([]NSMenuItem, count)
	for i := 0; i < count; i++ {
		o := a.ObjectAtIndex(uint64(i))
		items[i] = NSMenuItem{o}
	}
	return items
}
func (menu NSMenu) NumberOfItems() uint64 {
	return menu.Get("numberOfItems").Uint()
}

func (menu NSMenu) RemoveAllItems(){
	menu.Send("removeAllItems:", item)
}