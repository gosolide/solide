//js:package native/macos/core-foundation
package core_foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreFoundation.CFRunLoopSourceContext
*/

type CFRunLoopSourceContext struct {
  Version int64 `v8:"version"`
  Info void `v8:"info"`
  Retain func(...any) (any, error) `v8:"retain"`
  Release func(...any) (any, error) `v8:"release"`
  CopyDescription func(...any) (any, error) `v8:"copyDescription"`
  Equal func(...any) (any, error) `v8:"equal"`
  Hash func(...any) (any, error) `v8:"hash"`
  Schedule func(...any) (any, error) `v8:"schedule"`
  Cancel func(...any) (any, error) `v8:"cancel"`
  Perform func(...any) (any, error) `v8:"perform"`
}