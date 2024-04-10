package uuid_generators

import (
	infra_adapters "my-app/src/config/infra_adapters"

	uuid_generator_google_uuid "my-app/src/infra/uuid_generators/google_uuid"
	uuid_generator_gofrs_uuid "my-app/src/infra/uuid_generators/gofrs_uuid"
	uuid_generator_empty_uuid "my-app/src/infra/uuid_generators/empty_uuid"
)

type GenericUUIDGenerator interface{
	NewUUID() string
}

func Init() GenericUUIDGenerator {
	switch infra_adapters.GetAdapters()["UUID_GENERATOR"] {
		case "google_uuid":
			return uuid_generator_google_uuid.Init()
		case "gofrs_uuid":
		return uuid_generator_gofrs_uuid.Init()
		default:
			return uuid_generator_empty_uuid.Init()
	}
}
