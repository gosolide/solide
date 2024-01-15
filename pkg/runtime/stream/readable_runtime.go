// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package stream

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("stream", "readable.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "Readable", func (in isolates.FunctionArgs) (*ReadableBase, error) {
    return NewReadable(in)
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "Readable", constructor); err != nil {
    return nil, err
  }

  return nil, nil
})

func (r *ReadableBase) V8GetIsPaused(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.IsPaused()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8FuncPause(in isolates.FunctionArgs) (*isolates.Value, error) {
  r.Pause(in.ExecutionContext)
  return nil, nil
}

func (r *ReadableBase) V8FuncResume(in isolates.FunctionArgs) (*isolates.Value, error) {
  r.Resume(in.ExecutionContext)
  return nil, nil
}

func (r *ReadableBase) V8FuncPipe(in isolates.FunctionArgs) (*isolates.Value, error) {
  destination := in.Arg(in.ExecutionContext, 0)
  options := in.Arg(in.ExecutionContext, 1)
  if result, err := r.Pipe(in.ExecutionContext, destination, options); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (r *ReadableBase) V8FuncReadV8(in isolates.FunctionArgs) (*isolates.Value, error) {
  var size int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&size).Elem()); __err != nil {
    return nil, __err
  } else {
    size = v.Interface().(int)
  }

  if err := r.ReadV8(in.ExecutionContext, size); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *ReadableBase) V8GetReadable(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.Readable()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8GetReadableAborted(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.ReadableAborted()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8GetReadableDidRead(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.ReadableDidRead()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8GetReadableEncoding(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.ReadableEncoding()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8GetReadableEnded(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.ReadableEnded()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8GetReadableFlowing(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.ReadableFlowing()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8GetReadableHighWaterMark(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.ReadableHighWaterMark()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8GetReadableLength(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.ReadableLength()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8GetReadableObjectMode(in isolates.GetterArgs) (*isolates.Value, error) {
  result := r.ReadableObjectMode()
  return in.Context.Create(in.ExecutionContext, result)
}

func (r *ReadableBase) V8FuncSetEncoding(in isolates.FunctionArgs) (*isolates.Value, error) {
  var args0 BufferEncoding
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); __err != nil {
    return nil, __err
  } else {
    args0 = v.Interface().(BufferEncoding)
  }

  if result, err := r.SetEncoding(args0); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (r *ReadableBase) V8FuncUnpipe(in isolates.FunctionArgs) (*isolates.Value, error) {
  destination := in.Arg(in.ExecutionContext, 0)
  if result, err := r.Unpipe(in.ExecutionContext, destination); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (r *ReadableBase) V8FuncUnshift(in isolates.FunctionArgs) (*isolates.Value, error) {
  args0 := in.Arg(in.ExecutionContext, 0)
  var args1 *BufferEncoding
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); __err != nil {
    return nil, __err
  } else if v != nil {
    args1 = v.Interface().(*BufferEncoding)
  }

  if err := r.Unshift(in.ExecutionContext, args0, args1); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (r *ReadableBase) V8FuncWrap(in isolates.FunctionArgs) (*isolates.Value, error) {
  args0 := in.Arg(in.ExecutionContext, 0)
  if result, err := r.Wrap(in.ExecutionContext, args0); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (r *ReadableBase) V8FuncCompose(in isolates.FunctionArgs) (*isolates.Value, error) {
  args0 := in.Arg(in.ExecutionContext, 0)
  args1 := in.Arg(in.ExecutionContext, 1)
  if result, err := r.Compose(in.ExecutionContext, args0, args1); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (r *ReadableBase) V8FuncIterator(in isolates.FunctionArgs) (*isolates.Value, error) {
  args0 := in.Arg(in.ExecutionContext, 0)
  if result, err := r.Iterator(in.ExecutionContext, args0); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (r *ReadableBase) V8FuncMap(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

      if result, err := r.Map(in.ExecutionContext, args0, args1); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncFilter(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

      if result, err := r.Filter(in.ExecutionContext, args0, args1); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncForEach(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

      if err := r.ForEach(in.ExecutionContext, args0, args1); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Undefined(in.ExecutionContext); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncToArray(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

      if result, err := r.ToArray(in.ExecutionContext, args0); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncSome(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

      if result, err := r.Some(in.ExecutionContext, args0, args1); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncFind(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

      if result, err := r.Find(in.ExecutionContext, args0, args1); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncEvery(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

      if result, err := r.Every(in.ExecutionContext, args0, args1); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncFlatMap(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

      if result, err := r.FlatMap(in.ExecutionContext, args0, args1); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncDrop(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 int
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      args0 = v.Interface().(int)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

      if result, err := r.Drop(in.ExecutionContext, args0, args1); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncTake(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 int
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else {
      args0 = v.Interface().(int)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

      if result, err := r.Take(in.ExecutionContext, args0, args1); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncAsIndexedPairs(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

      if result, err := r.AsIndexedPairs(in.ExecutionContext, args0); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncReduce(in isolates.FunctionArgs) (*isolates.Value, error) {
  if resolver, err := in.Context.NewResolver(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    in.Background(func(in isolates.FunctionArgs) {
    var args0 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args0).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args0 = v.Interface().(*isolates.Value)
    }

    var args1 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args1).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args1 = v.Interface().(*isolates.Value)
    }

    var args2 *isolates.Value
    if v, err := in.Arg(in.ExecutionContext, 2).Unmarshal(in.ExecutionContext, reflect.TypeOf(&args2).Elem()); err != nil {
      resolver.Reject(in.ExecutionContext, err)
      return
    } else if v != nil {
      args2 = v.Interface().(*isolates.Value)
    }

      if result, err := r.Reduce(in.ExecutionContext, args0, args1, args2); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else if result, err := in.Context.Create(in.ExecutionContext, result); err != nil {
        resolver.Reject(in.ExecutionContext, err)
      } else {
        resolver.Resolve(in.ExecutionContext, result)
      }
    })

    return resolver.Promise(in.ExecutionContext)
  }
}

func (r *ReadableBase) V8FuncPush(in isolates.FunctionArgs) (*isolates.Value, error) {
  chunk := in.Arg(in.ExecutionContext, 0)
  var encoding *BufferEncoding
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&encoding).Elem()); __err != nil {
    return nil, __err
  } else if v != nil {
    encoding = v.Interface().(*BufferEncoding)
  }

  if result, err := r.Push(in.ExecutionContext, chunk, encoding); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (c *Reader) V8FuncConstruct(in isolates.FunctionArgs) (*isolates.Value, error) {
  var this Stream
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&this).Elem()); __err != nil {
    return nil, __err
  } else {
    this = v.Interface().(Stream)
  }

  callback := in.Arg(in.ExecutionContext, 1)
  if err := c.Construct(in, this, callback); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (c *Reader) V8FuncRead(in isolates.FunctionArgs) (*isolates.Value, error) {
  var this Readable
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&this).Elem()); __err != nil {
    return nil, __err
  } else {
    this = v.Interface().(Readable)
  }

  if err := c.Read(in, this); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (p *PipeBase) V8FuncRemove(in isolates.FunctionArgs) (*isolates.Value, error) {
  if err := p.Remove(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    return nil, nil
  }
}

func (p *PipeBase) V8GetSource(in isolates.GetterArgs) (*isolates.Value, error) {
  result := p.Source()
  return in.Context.Create(in.ExecutionContext, result)
}

func (p *PipeBase) V8GetDestination(in isolates.GetterArgs) (*isolates.Value, error) {
  result := p.Destination()
  return in.Context.Create(in.ExecutionContext, result)
}

func (p *PipeBase) V8GetOnData(in isolates.GetterArgs) (*isolates.Value, error) {
  result := p.OnData()
  return in.Context.Create(in.ExecutionContext, result)
}

func (p *PipeBase) V8GetOnError(in isolates.GetterArgs) (*isolates.Value, error) {
  result := p.OnError()
  return in.Context.Create(in.ExecutionContext, result)
}

func (p *PipeBase) V8GetOnEnd(in isolates.GetterArgs) (*isolates.Value, error) {
  result := p.OnEnd()
  return in.Context.Create(in.ExecutionContext, result)
}