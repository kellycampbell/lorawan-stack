// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

var GatewayBrandFieldPathsNested = []string{
	"id",
	"logos",
	"name",
	"url",
}

var GatewayBrandFieldPathsTopLevel = []string{
	"id",
	"logos",
	"name",
	"url",
}
var GatewayModelFieldPathsNested = []string{
	"brand_id",
	"id",
	"name",
}

var GatewayModelFieldPathsTopLevel = []string{
	"brand_id",
	"id",
	"name",
}
var GatewayVersionIdentifiersFieldPathsNested = []string{
	"brand_id",
	"firmware_version",
	"hardware_version",
	"model_id",
}

var GatewayVersionIdentifiersFieldPathsTopLevel = []string{
	"brand_id",
	"firmware_version",
	"hardware_version",
	"model_id",
}
var GatewayRadioFieldPathsNested = []string{
	"chip_type",
	"enable",
	"frequency",
	"rssi_offset",
	"tx_configuration",
	"tx_configuration.max_frequency",
	"tx_configuration.min_frequency",
	"tx_configuration.notch_frequency",
}

var GatewayRadioFieldPathsTopLevel = []string{
	"chip_type",
	"enable",
	"frequency",
	"rssi_offset",
	"tx_configuration",
}
var GatewayVersionFieldPathsNested = []string{
	"clock_source",
	"ids",
	"ids.brand_id",
	"ids.firmware_version",
	"ids.hardware_version",
	"ids.model_id",
	"photos",
	"radios",
}

var GatewayVersionFieldPathsTopLevel = []string{
	"clock_source",
	"ids",
	"photos",
	"radios",
}
var GatewayClaimAuthenticationCodeFieldPathsNested = []string{
	"secret",
	"secret.key_id",
	"secret.value",
	"valid_from",
	"valid_to",
}

var GatewayClaimAuthenticationCodeFieldPathsTopLevel = []string{
	"secret",
	"valid_from",
	"valid_to",
}
var GatewayFieldPathsNested = []string{
	"antennas",
	"attributes",
	"auto_update",
	"claim_authentication_code",
	"claim_authentication_code.secret",
	"claim_authentication_code.secret.key_id",
	"claim_authentication_code.secret.value",
	"claim_authentication_code.valid_from",
	"claim_authentication_code.valid_to",
	"contact_info",
	"created_at",
	"deleted_at",
	"description",
	"downlink_path_constraint",
	"enforce_duty_cycle",
	"frequency_plan_id",
	"frequency_plan_ids",
	"gateway_server_address",
	"ids",
	"ids.eui",
	"ids.gateway_id",
	"lbs_lns_secret",
	"lbs_lns_secret.key_id",
	"lbs_lns_secret.value",
	"location_public",
	"name",
	"require_authenticated_connection",
	"schedule_anytime_delay",
	"schedule_downlink_late",
	"status_public",
	"target_cups_key",
	"target_cups_key.key_id",
	"target_cups_key.value",
	"target_cups_uri",
	"update_channel",
	"update_location_from_status",
	"updated_at",
	"version_ids",
	"version_ids.brand_id",
	"version_ids.firmware_version",
	"version_ids.hardware_version",
	"version_ids.model_id",
}

var GatewayFieldPathsTopLevel = []string{
	"antennas",
	"attributes",
	"auto_update",
	"claim_authentication_code",
	"contact_info",
	"created_at",
	"deleted_at",
	"description",
	"downlink_path_constraint",
	"enforce_duty_cycle",
	"frequency_plan_id",
	"frequency_plan_ids",
	"gateway_server_address",
	"ids",
	"lbs_lns_secret",
	"location_public",
	"name",
	"require_authenticated_connection",
	"schedule_anytime_delay",
	"schedule_downlink_late",
	"status_public",
	"target_cups_key",
	"target_cups_uri",
	"update_channel",
	"update_location_from_status",
	"updated_at",
	"version_ids",
}
var GatewaysFieldPathsNested = []string{
	"gateways",
}

var GatewaysFieldPathsTopLevel = []string{
	"gateways",
}
var GetGatewayRequestFieldPathsNested = []string{
	"field_mask",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
}

var GetGatewayRequestFieldPathsTopLevel = []string{
	"field_mask",
	"gateway_ids",
}
var GetGatewayIdentifiersForEUIRequestFieldPathsNested = []string{
	"eui",
}

var GetGatewayIdentifiersForEUIRequestFieldPathsTopLevel = []string{
	"eui",
}
var ListGatewaysRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"deleted",
	"field_mask",
	"limit",
	"order",
	"page",
}

var ListGatewaysRequestFieldPathsTopLevel = []string{
	"collaborator",
	"deleted",
	"field_mask",
	"limit",
	"order",
	"page",
}
var CreateGatewayRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"gateway",
	"gateway.antennas",
	"gateway.attributes",
	"gateway.auto_update",
	"gateway.claim_authentication_code",
	"gateway.claim_authentication_code.secret",
	"gateway.claim_authentication_code.secret.key_id",
	"gateway.claim_authentication_code.secret.value",
	"gateway.claim_authentication_code.valid_from",
	"gateway.claim_authentication_code.valid_to",
	"gateway.contact_info",
	"gateway.created_at",
	"gateway.deleted_at",
	"gateway.description",
	"gateway.downlink_path_constraint",
	"gateway.enforce_duty_cycle",
	"gateway.frequency_plan_id",
	"gateway.frequency_plan_ids",
	"gateway.gateway_server_address",
	"gateway.ids",
	"gateway.ids.eui",
	"gateway.ids.gateway_id",
	"gateway.lbs_lns_secret",
	"gateway.lbs_lns_secret.key_id",
	"gateway.lbs_lns_secret.value",
	"gateway.location_public",
	"gateway.name",
	"gateway.require_authenticated_connection",
	"gateway.schedule_anytime_delay",
	"gateway.schedule_downlink_late",
	"gateway.status_public",
	"gateway.target_cups_key",
	"gateway.target_cups_key.key_id",
	"gateway.target_cups_key.value",
	"gateway.target_cups_uri",
	"gateway.update_channel",
	"gateway.update_location_from_status",
	"gateway.updated_at",
	"gateway.version_ids",
	"gateway.version_ids.brand_id",
	"gateway.version_ids.firmware_version",
	"gateway.version_ids.hardware_version",
	"gateway.version_ids.model_id",
}

var CreateGatewayRequestFieldPathsTopLevel = []string{
	"collaborator",
	"gateway",
}
var UpdateGatewayRequestFieldPathsNested = []string{
	"field_mask",
	"gateway",
	"gateway.antennas",
	"gateway.attributes",
	"gateway.auto_update",
	"gateway.claim_authentication_code",
	"gateway.claim_authentication_code.secret",
	"gateway.claim_authentication_code.secret.key_id",
	"gateway.claim_authentication_code.secret.value",
	"gateway.claim_authentication_code.valid_from",
	"gateway.claim_authentication_code.valid_to",
	"gateway.contact_info",
	"gateway.created_at",
	"gateway.deleted_at",
	"gateway.description",
	"gateway.downlink_path_constraint",
	"gateway.enforce_duty_cycle",
	"gateway.frequency_plan_id",
	"gateway.frequency_plan_ids",
	"gateway.gateway_server_address",
	"gateway.ids",
	"gateway.ids.eui",
	"gateway.ids.gateway_id",
	"gateway.lbs_lns_secret",
	"gateway.lbs_lns_secret.key_id",
	"gateway.lbs_lns_secret.value",
	"gateway.location_public",
	"gateway.name",
	"gateway.require_authenticated_connection",
	"gateway.schedule_anytime_delay",
	"gateway.schedule_downlink_late",
	"gateway.status_public",
	"gateway.target_cups_key",
	"gateway.target_cups_key.key_id",
	"gateway.target_cups_key.value",
	"gateway.target_cups_uri",
	"gateway.update_channel",
	"gateway.update_location_from_status",
	"gateway.updated_at",
	"gateway.version_ids",
	"gateway.version_ids.brand_id",
	"gateway.version_ids.firmware_version",
	"gateway.version_ids.hardware_version",
	"gateway.version_ids.model_id",
}

var UpdateGatewayRequestFieldPathsTopLevel = []string{
	"field_mask",
	"gateway",
}
var ListGatewayAPIKeysRequestFieldPathsNested = []string{
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
	"limit",
	"page",
}

var ListGatewayAPIKeysRequestFieldPathsTopLevel = []string{
	"gateway_ids",
	"limit",
	"page",
}
var GetGatewayAPIKeyRequestFieldPathsNested = []string{
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
	"key_id",
}

var GetGatewayAPIKeyRequestFieldPathsTopLevel = []string{
	"gateway_ids",
	"key_id",
}
var CreateGatewayAPIKeyRequestFieldPathsNested = []string{
	"expires_at",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
	"name",
	"rights",
}

var CreateGatewayAPIKeyRequestFieldPathsTopLevel = []string{
	"expires_at",
	"gateway_ids",
	"name",
	"rights",
}
var UpdateGatewayAPIKeyRequestFieldPathsNested = []string{
	"api_key",
	"api_key.created_at",
	"api_key.expires_at",
	"api_key.id",
	"api_key.key",
	"api_key.name",
	"api_key.rights",
	"api_key.updated_at",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
}

var UpdateGatewayAPIKeyRequestFieldPathsTopLevel = []string{
	"api_key",
	"gateway_ids",
}
var ListGatewayCollaboratorsRequestFieldPathsNested = []string{
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
	"limit",
	"page",
}

var ListGatewayCollaboratorsRequestFieldPathsTopLevel = []string{
	"gateway_ids",
	"limit",
	"page",
}
var GetGatewayCollaboratorRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.organization_ids",
	"collaborator.ids.organization_ids.organization_id",
	"collaborator.ids.user_ids",
	"collaborator.ids.user_ids.email",
	"collaborator.ids.user_ids.user_id",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
}

var GetGatewayCollaboratorRequestFieldPathsTopLevel = []string{
	"collaborator",
	"gateway_ids",
}
var SetGatewayCollaboratorRequestFieldPathsNested = []string{
	"collaborator",
	"collaborator.ids",
	"collaborator.ids.ids",
	"collaborator.ids.ids.organization_ids",
	"collaborator.ids.ids.organization_ids.organization_id",
	"collaborator.ids.ids.user_ids",
	"collaborator.ids.ids.user_ids.email",
	"collaborator.ids.ids.user_ids.user_id",
	"collaborator.rights",
	"gateway_ids",
	"gateway_ids.eui",
	"gateway_ids.gateway_id",
}

var SetGatewayCollaboratorRequestFieldPathsTopLevel = []string{
	"collaborator",
	"gateway_ids",
}
var GatewayAntennaFieldPathsNested = []string{
	"attributes",
	"gain",
	"location",
	"location.accuracy",
	"location.altitude",
	"location.latitude",
	"location.longitude",
	"location.source",
}

var GatewayAntennaFieldPathsTopLevel = []string{
	"attributes",
	"gain",
	"location",
}
var GatewayStatusFieldPathsNested = []string{
	"advanced",
	"antenna_locations",
	"boot_time",
	"ip",
	"metrics",
	"time",
	"versions",
}

var GatewayStatusFieldPathsTopLevel = []string{
	"advanced",
	"antenna_locations",
	"boot_time",
	"ip",
	"metrics",
	"time",
	"versions",
}
var GatewayConnectionStatsFieldPathsNested = []string{
	"connected_at",
	"downlink_count",
	"last_downlink_received_at",
	"last_status",
	"last_status.advanced",
	"last_status.antenna_locations",
	"last_status.boot_time",
	"last_status.ip",
	"last_status.metrics",
	"last_status.time",
	"last_status.versions",
	"last_status_received_at",
	"last_uplink_received_at",
	"protocol",
	"round_trip_times",
	"round_trip_times.count",
	"round_trip_times.max",
	"round_trip_times.median",
	"round_trip_times.min",
	"sub_bands",
	"uplink_count",
}

var GatewayConnectionStatsFieldPathsTopLevel = []string{
	"connected_at",
	"downlink_count",
	"last_downlink_received_at",
	"last_status",
	"last_status_received_at",
	"last_uplink_received_at",
	"protocol",
	"round_trip_times",
	"sub_bands",
	"uplink_count",
}
var GatewayRadio_TxConfigurationFieldPathsNested = []string{
	"max_frequency",
	"min_frequency",
	"notch_frequency",
}

var GatewayRadio_TxConfigurationFieldPathsTopLevel = []string{
	"max_frequency",
	"min_frequency",
	"notch_frequency",
}
var GatewayConnectionStats_RoundTripTimesFieldPathsNested = []string{
	"count",
	"max",
	"median",
	"min",
}

var GatewayConnectionStats_RoundTripTimesFieldPathsTopLevel = []string{
	"count",
	"max",
	"median",
	"min",
}
var GatewayConnectionStats_SubBandFieldPathsNested = []string{
	"downlink_utilization",
	"downlink_utilization_limit",
	"max_frequency",
	"min_frequency",
}

var GatewayConnectionStats_SubBandFieldPathsTopLevel = []string{
	"downlink_utilization",
	"downlink_utilization_limit",
	"max_frequency",
	"min_frequency",
}
