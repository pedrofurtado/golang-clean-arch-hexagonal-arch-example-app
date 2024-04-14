package uuid_generators

import (
	"os"
	"fmt"

	uuidGeneratorGoogleUUID "my-app/internal/infra/uuid_generators/google_uuid"
	uuidGeneratorGofrsUUID "my-app/internal/infra/uuid_generators/gofrs_uuid"
)

type GenericUUIDGenerator interface{
	NewUUID() string
}

func Init() GenericUUIDGenerator {
	switch os.Getenv("APP_ADAPTER_UUID_GENERATOR") {
		case "google_uuid":
			return uuidGeneratorGoogleUUID.Init()
		case "gofrs_uuid":
			return uuidGeneratorGofrsUUID.Init()
		default:
			err := "Must be defined a adapter for uuid generator"
			fmt.Println(err)
			panic(err)
	}
}
