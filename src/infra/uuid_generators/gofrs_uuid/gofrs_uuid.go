package uuid_generators

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type GofrsUUIDGenerator struct{}

func(u GofrsUUIDGenerator) NewUUID() string {
	generated_uuid, _ := uuid.NewV4()
	return generated_uuid.String()
}

func Init() GofrsUUIDGenerator {
	fmt.Println("UUID Generator | Init | Implementation: Gofrs")
	return GofrsUUIDGenerator{}
}
