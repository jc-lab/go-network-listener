package retainer

var retainedObjects = make(map[uintptr]interface{})

func Retain(ptr uintptr, obj interface{}) uintptr {
	retainedObjects[ptr] = obj
	return ptr
}

func Get(ptr uintptr) interface{} {
	return retainedObjects[ptr]
}

func Release(ptr uintptr) {
	delete(retainedObjects, ptr)
}
