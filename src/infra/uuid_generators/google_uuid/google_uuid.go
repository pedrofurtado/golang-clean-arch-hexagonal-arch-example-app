package uuid_generators

import (
	"fmt"

	"github.com/google/uuid"
)

type GoogleUUIDGenerator struct{}

func(u GoogleUUIDGenerator) NewUUID() string {
	return uuid.New().String()
}

func Init() GoogleUUIDGenerator {
	fmt.Println("UUID Generator | Init | Implementation: Google")
	return GoogleUUIDGenerator{}
}
