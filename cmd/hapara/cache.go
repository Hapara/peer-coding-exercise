package hapara

// / GetFromCache returns the value from the change or an error if the value is not in the cache
func GetFromCache(id string) (string, error) {
	return "", nil
}

// / PutInCache adds a value to the cache or an error if the cache can not be populated
// / the cache has a finite size, if putting the item in the cache would exceed the size limit
// / the drop the cache item that has aged the most
func PutInCache(id string, value string) error {
	return nil
}
