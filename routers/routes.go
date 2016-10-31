package routers

import "github.com/apisearch/importer/routers/api/v1"

var routes = Routes{
	Route{
		"Healthz",
		"GET",
		"/api/v1/status/healthz",
		v1.Healthz,
	},
	Route{
		"GetSettingsByUserId",
		"GET",
		"/api/v1/settings/{userId}",
		v1.GetSettingsById,
	},
	Route{
		"CreateNewSettings",
		"POST",
		"/api/v1/settings/{userId}",
		v1.CreateSettings,
	},
	Route{
		"DeleteSettings",
		"DELETE",
		"/api/v1/settings/{userId}",
		v1.DeleteSettings,
	},
}
