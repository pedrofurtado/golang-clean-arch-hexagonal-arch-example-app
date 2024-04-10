package uuid_generators

import (
	"fmt"
)

type EmptyUUIDGenerator struct{}

func(u EmptyUUIDGenerator) NewUUID() string {
	return ""
}

func Init() EmptyUUIDGenerator {
	fmt.Println("UUID Generator | Init | Implementation: Empty")
	return EmptyUUIDGenerator{}
}
