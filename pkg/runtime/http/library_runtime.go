// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package http

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("http", "library.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "Http", func (in isolates.FunctionArgs) (*HttpBase, error) {
    return NewHttp(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "Http", constructor); err != nil {
    return nil, err
  }

var Module *isolates.Module
Exports := in.Args[1]
Require := in.Args[2]
if m, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(Module)); err != nil {
  return nil, err
} else {
  Module = m.Interface().(*isolates.Module)
}


  {
    fnName := "createDefaultHttp"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
rin := isolates.RuntimeFunctionArgs{FunctionArgs: in, Module: Module, Exports: Exports, Require: Require}

      if result, err := NewDefaultHttp(rin); err != nil {
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

func (h *HttpBase) V8GetNet(in isolates.GetterArgs) (*isolates.Value, error) {
  result := h.Net()
  return in.Context.Create(in.ExecutionContext, result)
}

func (h *HttpBase) V8GetServer(in isolates.GetterArgs) (*isolates.Value, error) {
  result := h.Server()
  return in.Context.Create(in.ExecutionContext, result)
}

func (h *HttpBase) V8GetAgent(in isolates.GetterArgs) (*isolates.Value, error) {
  result := h.Agent()
  return in.Context.Create(in.ExecutionContext, result)
}

func (h *HttpBase) V8GetGlobalAgent(in isolates.GetterArgs) (*isolates.Value, error) {
  result := h.GlobalAgent()
  return in.Context.Create(in.ExecutionContext, result)
}