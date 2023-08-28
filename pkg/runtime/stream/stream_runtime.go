// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package stream

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("stream", "/Users/tim/src/grexie/solid/pkg/runtime/stream/stream.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  {
    fnName := "pipeline"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
  args := make([]*isolates.Value, len(in.Args)-0)
      for i, arg := range in.Args[0:] {
        args[i] = arg
      }

      result := Pipeline(in.ExecutionContext, args...)
      return in.Context.Create(in.ExecutionContext, result)
    }); err != nil {
      return nil, err
    } else if err := in.Args[1].Set(in.ExecutionContext, "pipeline", fn); err != nil {
      return nil, err
    }
  }

  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "Stream", func (in isolates.FunctionArgs) (*StreamBase, error) {
    return NewStream(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "Stream", constructor); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "default", constructor); err != nil {
    return nil, err
  }

  return nil, nil
})

func (s *StreamBase) V8FuncDestroy(in isolates.FunctionArgs) (*isolates.Value, error) {
  err := in.Arg(in.ExecutionContext, 0)
  if err := s.Destroy(in.ExecutionContext, err); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (s *StreamBase) V8GetClosed(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Closed()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *StreamBase) V8GetDestroyed(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Destroyed()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *StreamBase) V8GetErrored(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Errored()
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Closer) V8FuncDestroy(in isolates.FunctionArgs) (*isolates.Value, error) {
  var this Stream
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&this).Elem()); __err != nil {
    return nil, __err
  } else {
    this = v.Interface().(Stream)
  }

  err := in.Arg(in.ExecutionContext, 1)
  callback := in.Arg(in.ExecutionContext, 2)
  if err := c.Destroy(in, this, err, callback); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}