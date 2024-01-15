//js:package native/macos/metal
package metal

//go:generate go run github.com/grexie/isolates/codegen

/*
struct Metal.MTLAccelerationStructureInstanceDescriptor
*/

type MTLAccelerationStructureInstanceDescriptor struct {
  TransformationMatrix MTLPackedFloat4x3 `v8:"transformationMatrix"`
  Options int `v8:"options"`
  Mask uint `v8:"mask"`
  IntersectionFunctionTableOffset uint `v8:"intersectionFunctionTableOffset"`
  AccelerationStructureIndex uint `v8:"accelerationStructureIndex"`
}