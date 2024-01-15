//js:package native/macos/core-services
package core_services

//go:generate go run github.com/grexie/isolates/codegen

/*
struct CoreServices.AEDesc
*/

type AEDesc struct {
  DescriptorType uint `v8:"descriptorType"`
  DataHandle **OpaqueAEDataStorageType `v8:"dataHandle"`
}