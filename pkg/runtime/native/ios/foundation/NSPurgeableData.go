//js:package native/ios/foundation
package foundation

//go:generate go run github.com/grexie/isolates/codegen

/*
interface Foundation.NSPurgeableData : Foundation.NSMutableData
*/

type NSPurgeableData struct {
  *NSMutableData

}
