package errcode

import (
	"fmt"
	"sort"
	"sync"
)

var (
	errorCodeToDescriptors = map[ErrorCode]ErrorDescriptor{}
	//	idToDescriptors        = map[string]ErrorDescriptor{}
	groupToDescriptors = map[string][]ErrorDescriptor{}
)

//var nextCode = 1000
var registerLock sync.Mutex

// Register will make the passed-in error known to the environment and
// return a new ErrorCode
func Register(group string, errcode int, descriptor ErrorDescriptor) ErrorCode {
	registerLock.Lock()
	defer registerLock.Unlock()
	descriptor.Code = ErrorCode(errcode) //1000
	//	if _, ok := idToDescriptors[descriptor.Value]; ok {
	//		panic(fmt.Sprintf("ErrorValue %q is already registered", descriptor.Value))
	//	}
	if _, ok := errorCodeToDescriptors[descriptor.Code]; ok {
		panic(fmt.Sprintf("ErrorCode %v is already registered", descriptor.Code))
	}
	groupToDescriptors[group] = append(groupToDescriptors[group], descriptor)
	errorCodeToDescriptors[descriptor.Code] = descriptor
	//	idToDescriptors[descriptor.Value] = descriptor
	return descriptor.Code
}

type byValue []ErrorDescriptor

func (a byValue) Len() int           { return len(a) }
func (a byValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byValue) Less(i, j int) bool { return a[i].Value < a[j].Value }

// GetGroupNames returns the list of Error group names that are registered
func GetGroupNames() []string {
	keys := []string{}

	for k := range groupToDescriptors {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// GetErrorCodeGroup returns the named group of error descriptors
func GetErrorCodeGroup(name string) []ErrorDescriptor {
	desc := groupToDescriptors[name]
	sort.Sort(byValue(desc))
	return desc
}

// GetErrorAllDescriptors returns a slice of all ErrorDescriptors that are
// registered, irrespective of what group they're in
func GetErrorAllDescriptors() []ErrorDescriptor {
	result := []ErrorDescriptor{}

	for _, group := range GetGroupNames() {
		result = append(result, GetErrorCodeGroup(group)...)
	}
	sort.Sort(byValue(result))
	return result
}
