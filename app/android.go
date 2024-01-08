//go:build android
// +build android

package app // import "golang.org/x/mobile/app"

import fynemobile "fyne.io/fyne/v2/driver"

func RunOnJVM(f func(vm, env, ctx uintptr) error) error {
	return fynemobile.RunNative(func(ctx interface{}) error {
		data := ctx.(*fynemobile.AndroidContext)
		return f(data.VM, data.Env, data.Ctx)
	})
}
