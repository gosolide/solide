// this file is auto-generated by github.com/grexie/isolates, DO NOT EDIT

package fs

import (
  isolates "github.com/grexie/isolates"
)

var _ = isolates.RegisterRuntime("fs", "/Users/tim/src/grexie/solid/pkg/runtime/fs/stats.go", func (in isolates.FunctionArgs) (*isolates.Value, error) {
  return nil, nil
})

func (s *stats) V8FuncIsDirectory(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := s.IsDirectory()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8FuncIsFile(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := s.IsFile()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8FuncIsSymbolicLink(in isolates.FunctionArgs) (*isolates.Value, error) {
  result := s.IsSymbolicLink()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetContentType(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.ContentType()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetAtimeMs(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.AccessedMilli()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetMtimeMs(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.ModifiedMilli()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetCtimeMs(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.ChangedMilli()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetBirthtimeMs(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.CreatedMilli()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetAtime(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Accessed()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetMtime(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Modified()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetCtime(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Changed()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetBirthtime(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Created()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetSize(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Size()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetUid(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Uid()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetGid(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Gid()
  return in.Context.Create(in.ExecutionContext, result)
}

func (s *stats) V8GetMode(in isolates.GetterArgs) (*isolates.Value, error) {
  result := s.Mode()
  return in.Context.Create(in.ExecutionContext, result)
}