// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package util

import (
  isolates "github.com/grexie/isolates"
  time "time"
  reflect "reflect"
)

var _ = isolates.RegisterRuntime("util", "counter.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "Counter", func (in isolates.FunctionArgs) (*Counter, error) {
    var _size int64
    if v, err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&_size).Elem()); err != nil {
      return nil, err
    } else {
      _size = v.Interface().(int64)
    }

    var _duration time.Duration
    if v, err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&_duration).Elem()); err != nil {
      return nil, err
    } else {
      _duration = v.Interface().(time.Duration)
    }

    return NewCounter(_size, _duration), nil
  }); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "Counter", constructor); err != nil {
    return nil, err
  }

  return nil, nil
})

func (c *Counter) V8FuncClone(in isolates.FunctionArgs) (*isolates.Value, error) {
  if result, err := c.Clone(in.ExecutionContext); err != nil {
    return nil, err
  } else {
    return in.Context.Create(in.ExecutionContext, result)
  }
}

func (c *Counter) V8FuncFloorTime(in isolates.FunctionArgs) (*isolates.Value, error) {
  var now time.Time
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&now).Elem()); __err != nil {
    return nil, __err
  } else {
    now = v.Interface().(time.Time)
  }

  result := c.FloorTime(now)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncAdd(in isolates.FunctionArgs) (*isolates.Value, error) {
  var value int64
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&value).Elem()); __err != nil {
    return nil, __err
  } else {
    value = v.Interface().(int64)
  }

  var now time.Time
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&now).Elem()); __err != nil {
    return nil, __err
  } else {
    now = v.Interface().(time.Time)
  }

  result := c.Add(value, now)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncAddNow(in isolates.FunctionArgs) (*isolates.Value, error) {
  var value int64
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&value).Elem()); __err != nil {
    return nil, __err
  } else {
    value = v.Interface().(int64)
  }

  result := c.AddNow(value)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncAddDuration(in isolates.FunctionArgs) (*isolates.Value, error) {
  var from time.Time
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&from).Elem()); __err != nil {
    return nil, __err
  } else {
    from = v.Interface().(time.Time)
  }

  var to time.Time
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&to).Elem()); __err != nil {
    return nil, __err
  } else {
    to = v.Interface().(time.Time)
  }

  result := c.AddDuration(from, to)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncAddDurationNow(in isolates.FunctionArgs) (*isolates.Value, error) {
  var from time.Time
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&from).Elem()); __err != nil {
    return nil, __err
  } else {
    from = v.Interface().(time.Time)
  }

  result := c.AddDurationNow(from)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncSub(in isolates.FunctionArgs) (*isolates.Value, error) {
  var value int64
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&value).Elem()); __err != nil {
    return nil, __err
  } else {
    value = v.Interface().(int64)
  }

  var now time.Time
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&now).Elem()); __err != nil {
    return nil, __err
  } else {
    now = v.Interface().(time.Time)
  }

  result := c.Sub(value, now)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncSumNow(in isolates.FunctionArgs) (*isolates.Value, error) {
  var duration time.Duration
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&duration).Elem()); __err != nil {
    return nil, __err
  } else {
    duration = v.Interface().(time.Duration)
  }

  result := c.SumNow(duration)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncSum(in isolates.FunctionArgs) (*isolates.Value, error) {
  var duration time.Duration
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&duration).Elem()); __err != nil {
    return nil, __err
  } else {
    duration = v.Interface().(time.Duration)
  }

  var now time.Time
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&now).Elem()); __err != nil {
    return nil, __err
  } else {
    now = v.Interface().(time.Time)
  }

  result := c.Sum(duration, now)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncSumAllNow(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := c.SumAllNow()
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncSumAll(in isolates.FunctionArgs) (*isolates.Value, error) {
  var now time.Time
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&now).Elem()); __err != nil {
    return nil, __err
  } else {
    now = v.Interface().(time.Time)
  }

  result := c.SumAll(now)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncSumDurationNow(in isolates.FunctionArgs) (*isolates.Value, error) {
  var duration time.Duration
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&duration).Elem()); __err != nil {
    return nil, __err
  } else {
    duration = v.Interface().(time.Duration)
  }

  result := c.SumDurationNow(duration)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncSumAllDurationNow(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := c.SumAllDurationNow()
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncSumDuration(in isolates.FunctionArgs) (*isolates.Value, error) {
  var duration time.Duration
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&duration).Elem()); __err != nil {
    return nil, __err
  } else {
    duration = v.Interface().(time.Duration)
  }

  var now time.Time
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&now).Elem()); __err != nil {
    return nil, __err
  } else {
    now = v.Interface().(time.Time)
  }

  result := c.SumDuration(duration, now)
  return in.Context.Create(in.ExecutionContext, result)
}

func (c *Counter) V8FuncSumAllDuration(in isolates.FunctionArgs) (*isolates.Value, error) {
  var now time.Time
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&now).Elem()); __err != nil {
    return nil, __err
  } else {
    now = v.Interface().(time.Time)
  }

  result := c.SumAllDuration(now)
  return in.Context.Create(in.ExecutionContext, result)
}