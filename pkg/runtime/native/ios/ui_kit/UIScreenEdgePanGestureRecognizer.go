//js:package native/ios/ui-kit
package ui_kit

//go:generate go run github.com/grexie/isolates/codegen

import (
  "fmt"
)

/*
interface UIKit.UIScreenEdgePanGestureRecognizer : UIKit.UIPanGestureRecognizer
*/

type UIScreenEdgePanGestureRecognizer struct {
  *UIPanGestureRecognizer

}

func (r *UIScreenEdgePanGestureRecognizer) Edges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}

func (r *UIScreenEdgePanGestureRecognizer) SetEdges() (any, error) {
  return nil, fmt.Errorf("unimplemented")
}
