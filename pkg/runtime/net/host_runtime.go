// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package net

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
  stream "github.com/gosolid/solid/pkg/runtime/stream"
)

var _ = isolates.RegisterRuntime("net", "host.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if _, err := in.Context.CreateWithName(in.ExecutionContext, "HostNetOptions", func (in isolates.FunctionArgs) (*HostNetOptions, error) {
    return NewHostNetOptions(in)
  }); err != nil {
    return nil, err
  }

  {
    fnName := "createHostNet"
    if fn, err := in.Context.CreateFunction(in.ExecutionContext, &fnName, func (in isolates.FunctionArgs) (*isolates.Value, error) {
      if result, err := NewHostNet(in.ExecutionContext); err != nil {
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

  if _, err := in.Context.CreateWithName(in.ExecutionContext, "HostSocketOptions", func (in isolates.FunctionArgs) (*HostSocketOptions, error) {
    return NewHostSocketOptions(in)
  }); err != nil {
    return nil, err
  }

  return nil, nil
})

func (n *HostNetOptions) V8FuncConnect(in isolates.FunctionArgs) (*isolates.Value, error) {
  var net Net
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&net).Elem()); __err != nil {
    return nil, __err
  } else {
    net = v.Interface().(Net)
  }

  var network string
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&network).Elem()); __err != nil {
    return nil, __err
  } else {
    network = v.Interface().(string)
  }

  var address string
  if v, __err := in.Arg(in.ExecutionContext, 2).Unmarshal(in.ExecutionContext, reflect.TypeOf(&address).Elem()); __err != nil {
    return nil, __err
  } else {
    address = v.Interface().(string)
  }

  if result, err := n.Connect(in, net, network, address); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (n *HostNetOptions) V8FuncListen(in isolates.FunctionArgs) (*isolates.Value, error) {
  var server Server
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&server).Elem()); __err != nil {
    return nil, __err
  } else {
    server = v.Interface().(Server)
  }

  var network string
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&network).Elem()); __err != nil {
    return nil, __err
  } else {
    network = v.Interface().(string)
  }

  var address string
  if v, __err := in.Arg(in.ExecutionContext, 2).Unmarshal(in.ExecutionContext, reflect.TypeOf(&address).Elem()); __err != nil {
    return nil, __err
  } else {
    address = v.Interface().(string)
  }

  if result, err := n.Listen(in, server, network, address); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (c *HostSocketOptions) V8FuncLocalAddress(in isolates.FunctionArgs) (*isolates.Value, error) {
  var this *SocketBase
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&this).Elem()); __err != nil {
    return nil, __err
  } else if v != nil {
    this = v.Interface().(*SocketBase)
  }

  if result, err := c.LocalAddress(in.ExecutionContext, this); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (c *HostSocketOptions) V8FuncRemoteAddress(in isolates.FunctionArgs) (*isolates.Value, error) {
  var this *SocketBase
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&this).Elem()); __err != nil {
    return nil, __err
  } else if v != nil {
    this = v.Interface().(*SocketBase)
  }

  if result, err := c.RemoteAddress(in.ExecutionContext, this); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (c *HostSocketOptions) V8FuncFinal(in isolates.FunctionArgs) (*isolates.Value, error) {
  var this stream.Writable
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&this).Elem()); __err != nil {
    return nil, __err
  } else {
    this = v.Interface().(stream.Writable)
  }

  callback := in.Arg(in.ExecutionContext, 1)
  if err := c.Final(in, this, callback); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}