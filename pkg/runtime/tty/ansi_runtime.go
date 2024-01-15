// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package tty

import (
  reflect "reflect"
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("tty", "ansi.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  if constructor, err := in.Context.CreateWithName(in.ExecutionContext, "ANSI", func (in isolates.FunctionArgs) (*AnsiEscapeSequence, error) {
    return NewAnsiEscapeSequence(), nil
  }); err != nil {
    return nil, err
  } else if instance, err := constructor.New(in.ExecutionContext); err != nil {
    return nil, err
  } else if err := in.Args[1].Set(in.ExecutionContext, "ansi", instance); err != nil {
    return nil, err
  }

  return nil, nil
})

func (e *AnsiEscapeSequence) V8GetEnabled(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Enabled()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEnable(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.Enable()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncDisable(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.Disable()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncHome(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.Home()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncStrip(in isolates.FunctionArgs) (*isolates.Value, error) {
  var s string
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&s).Elem()); __err != nil {
    return nil, __err
  } else {
    s = v.Interface().(string)
  }

  result := e.Strip(s)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncMove(in isolates.FunctionArgs) (*isolates.Value, error) {
  var line int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&line).Elem()); __err != nil {
    return nil, __err
  } else {
    line = v.Interface().(int)
  }

  var column int
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&column).Elem()); __err != nil {
    return nil, __err
  } else {
    column = v.Interface().(int)
  }

  result := e.Move(line, column)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncMoveUp(in isolates.FunctionArgs) (*isolates.Value, error) {
  var lines int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&lines).Elem()); __err != nil {
    return nil, __err
  } else {
    lines = v.Interface().(int)
  }

  result := e.MoveUp(lines)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncMoveDown(in isolates.FunctionArgs) (*isolates.Value, error) {
  var lines int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&lines).Elem()); __err != nil {
    return nil, __err
  } else {
    lines = v.Interface().(int)
  }

  result := e.MoveDown(lines)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncMoveRight(in isolates.FunctionArgs) (*isolates.Value, error) {
  var columns int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&columns).Elem()); __err != nil {
    return nil, __err
  } else {
    columns = v.Interface().(int)
  }

  result := e.MoveRight(columns)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncMoveLeft(in isolates.FunctionArgs) (*isolates.Value, error) {
  var columns int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&columns).Elem()); __err != nil {
    return nil, __err
  } else {
    columns = v.Interface().(int)
  }

  result := e.MoveLeft(columns)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncMoveToStartDown(in isolates.FunctionArgs) (*isolates.Value, error) {
  var lines int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&lines).Elem()); __err != nil {
    return nil, __err
  } else {
    lines = v.Interface().(int)
  }

  result := e.MoveToStartDown(lines)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncMoveToStartUp(in isolates.FunctionArgs) (*isolates.Value, error) {
  var lines int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&lines).Elem()); __err != nil {
    return nil, __err
  } else {
    lines = v.Interface().(int)
  }

  result := e.MoveToStartUp(lines)
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncMoveToStart(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.MoveToStart()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncRequestCursorPosition(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.RequestCursorPosition()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncMoveUpOne(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.MoveUpOne()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncSaveCursorPosition(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.SaveCursorPosition()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncRestoreCursorPosition(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.RestoreCursorPosition()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncSaveCursorPositionSCO(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.SaveCursorPositionSCO()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncRequestCursorPositionSCO(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.RequestCursorPositionSCO()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEraseInDisplay(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EraseInDisplay()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEraseFromCursorToEndOfScreen(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EraseFromCursorToEndOfScreen()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEraseFromCursorToStartOfScreen(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EraseFromCursorToStartOfScreen()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEraseEntireScreen(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EraseEntireScreen()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEraseSavedLines(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EraseSavedLines()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEraseInLine(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EraseInLine()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEraseFromCursorToEndOfLine(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EraseFromCursorToEndOfLine()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEraseFromCursorToStartOfLine(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EraseFromCursorToStartOfLine()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8FuncEraseLine(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := e.EraseLine()
  return in.Context.Create(in.ExecutionContext, result)
}

func (e *AnsiEscapeSequence) V8GetReset(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Reset()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBold(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Bold()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBoldReset(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BoldReset()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetDim(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Dim()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetDimReset(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.DimReset()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetItalic(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Italic()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetItalicReset(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.ItalicReset()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetUnderline(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Underline()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetUnderlineReset(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.UnderlineReset()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBlink(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Blink()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBlinkReset(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BlinkReset()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetReverse(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Reverse()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetReverseReset(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.ReverseReset()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetHidden(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Hidden()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetHiddenReset(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.HiddenReset()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetStrikethrough(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Strikethrough()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetStrikethroughReset(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.StrikethroughReset()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBlack(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Black()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBlack(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBlack()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBrightBlack(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BrightBlack()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBrightBlack(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBrightBlack()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetRed(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Red()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundRed(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundRed()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBrightRed(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BrightRed()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBrightRed(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBrightRed()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetGreen(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Green()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundGreen(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundGreen()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBrightGreen(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BrightGreen()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBrightGreen(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBrightGreen()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetYellow(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Yellow()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundYellow(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundYellow()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBrightYellow(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BrightYellow()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBrightYellow(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBrightYellow()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBlue(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Blue()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBlue(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBlue()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBrightBlue(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BrightBlue()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBrightBlue(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBrightBlue()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetMagenta(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Magenta()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundMagenta(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundMagenta()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBrightMagenta(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BrightMagenta()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBrightMagenta(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBrightMagenta()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetCyan(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Cyan()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundCyan(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundCyan()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBrightCyan(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BrightCyan()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBrightCyan(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBrightCyan()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetWhite(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.White()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundWhite(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundWhite()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBrightWhite(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BrightWhite()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundBrightWhite(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundBrightWhite()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetDefault(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.Default()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8GetBackgroundDefault(in isolates.GetterArgs) (*isolates.Value, error) {
  result := e.BackgroundDefault()
  return result.getterDecorator(in)

}

func (e *AnsiEscapeSequence) V8FuncIndexedColor(in isolates.FunctionArgs) (*isolates.Value, error) {
  var index int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&index).Elem()); __err != nil {
    return nil, __err
  } else {
    index = v.Interface().(int)
  }

  result := e.IndexedColor(index)
  return result.functionDecorator(in)

}

func (e *AnsiEscapeSequence) V8FuncBackgroundIndexedColor(in isolates.FunctionArgs) (*isolates.Value, error) {
  var index int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&index).Elem()); __err != nil {
    return nil, __err
  } else {
    index = v.Interface().(int)
  }

  result := e.BackgroundIndexedColor(index)
  return result.functionDecorator(in)

}

func (e *AnsiEscapeSequence) V8FuncColor(in isolates.FunctionArgs) (*isolates.Value, error) {
  var r int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&r).Elem()); __err != nil {
    return nil, __err
  } else {
    r = v.Interface().(int)
  }

  var g int
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&g).Elem()); __err != nil {
    return nil, __err
  } else {
    g = v.Interface().(int)
  }

  var b int
  if v, __err := in.Arg(in.ExecutionContext, 2).Unmarshal(in.ExecutionContext, reflect.TypeOf(&b).Elem()); __err != nil {
    return nil, __err
  } else {
    b = v.Interface().(int)
  }

  result := e.Color(r, g, b)
  return result.functionDecorator(in)

}

func (e *AnsiEscapeSequence) V8FuncBackgroundColor(in isolates.FunctionArgs) (*isolates.Value, error) {
  var r int
  if v, __err := in.Arg(in.ExecutionContext, 0).Unmarshal(in.ExecutionContext, reflect.TypeOf(&r).Elem()); __err != nil {
    return nil, __err
  } else {
    r = v.Interface().(int)
  }

  var g int
  if v, __err := in.Arg(in.ExecutionContext, 1).Unmarshal(in.ExecutionContext, reflect.TypeOf(&g).Elem()); __err != nil {
    return nil, __err
  } else {
    g = v.Interface().(int)
  }

  var b int
  if v, __err := in.Arg(in.ExecutionContext, 2).Unmarshal(in.ExecutionContext, reflect.TypeOf(&b).Elem()); __err != nil {
    return nil, __err
  } else {
    b = v.Interface().(int)
  }

  result := e.BackgroundColor(r, g, b)
  return result.functionDecorator(in)

}