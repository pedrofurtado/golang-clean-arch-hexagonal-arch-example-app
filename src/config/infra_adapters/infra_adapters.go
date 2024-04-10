package config

func GetAdapters() map[string]string {
	return map[string]string{
		// Possible values: chi | gorilla_mux | julienschmidt_httprouter
		"HTTP_ROUTER": "julienschmidt_httprouter",

		// Possible values: google_uuid | gofrs_uuid (if blank: empty_uuid)
		"UUID_GENERATOR": "gofrs_uuid",
	}
}
