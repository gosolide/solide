//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface UIKit.UIDropInteraction : objc.NSObject
*/

type UIDropInteraction struct {
  *objc.NSObject

}

func (r *UIDropInteraction) InitWithDelegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropInteraction) Init() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropInteraction) New() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropInteraction) Delegate() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropInteraction) AllowsSimultaneousDropSessions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIDropInteraction) SetAllowsSimultaneousDropSessions() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}
