//js:package native/ios/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

import (
  "github.com/gosolid/solid/pkg/runtime/native/ios/objc"
  "fmt"
)

/*
interface Metal.MTLSharedEventHandle : objc.NSObject
*/

type MTLSharedEventHandle struct {
  *objc.NSObject

}

func (r *MTLSharedEventHandle) Label() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}
