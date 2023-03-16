package hapara

import "peer-coding/cmd/db"

func DoWork(id string) string {

	//Implement the cache system in `cache.go` and wire it into
	//this middleware
	return db.GetById(id)
}
