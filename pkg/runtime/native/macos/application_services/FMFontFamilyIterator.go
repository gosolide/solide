//js:package native/macos/application-services
package application_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct ApplicationServices.FMFontFamilyIterator
*/

type FMFontFamilyIterator struct {
  Reserved [16]uint `v8:"reserved"`
}