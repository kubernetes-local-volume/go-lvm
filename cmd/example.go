package main

import (
	"fmt"
	"github.com/nak3/go-lvm"
)

func main() {
	valid := lvm.VgNameValidate("fedora")
	fmt.Printf("%#v\n", valid)

	vglist := lvm.ListVgNames()
	lvm.ListVgUUIDs()
	lvm.Open()
	a := &lvm.VgObject{}
	//a.vgt = LvmVgOpen(vglist[1], "r")
	a.Vgt = lvm.VgOpen(vglist[1], "w")

	fmt.Printf("size: %d GiB\n", uint64(a.GetSize())/1024/1024/1024)

	fmt.Printf("pvlist: %#v\n", a.ListPVs())

	fmt.Printf("listLVs: %#v\n", a.ListLVs())

	// TODO /1024
	fmt.Printf("Free size: %d KiB\n", uint64(a.GetFreeSize())/1024/1024)

	l := &lvm.LvObject{}

	l = a.CreateLvLinear("foo", int64(a.GetFreeSize())/1024/1024/2)

	fmt.Printf("LV UUID: %#v\n", l.GetUuid())

	l.AddTag("Demo_tag")

	//	time.Sleep(10 * time.Second) // 3秒休む
	l.RemoveTag("Demo_tag")

	l.Remove()

}
