// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package fs

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("fs", "/Users/tim/src/grexie/solid/pkg/runtime/fs/library.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
var Module *isolates.Module
Exports := in.Args[1]
Require := in.Args[2]
if m, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(Module)); err != nil {
  return nil, err
} else {
  Module = m.Interface().(*isolates.Module)
}


  {
    fnName := "createLibrary"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
rin := isolates.RuntimeFunctionArgs{FunctionArgs: in, Module: Module, Exports: Exports, Require: Require}

      if result, err := createLibrary(rin); err != nil {
        return nil, err
      } else {
        return in.Context.Create(in.ExecutionContext, result)
      }
    }); err != nil {
      return nil, err
    } else if instance, err := fn.Call(in.ExecutionContext, nil); err != nil {
      return nil, err
    } else if err := in.Args[1].Set(in.ExecutionContext, "default", instance); err != nil {
      return nil, err
    }
  }

  return nil, nil
})

