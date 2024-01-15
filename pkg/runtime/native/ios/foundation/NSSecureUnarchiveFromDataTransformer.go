//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface Foundation.NSSecureUnarchiveFromDataTransformer : Foundation.NSValueTransformer
*/

type NSSecureUnarchiveFromDataTransformer struct {
  *NSValueTransformer

}

func (r *NSSecureUnarchiveFromDataTransformer) AllowedTopLevelClasses() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}
