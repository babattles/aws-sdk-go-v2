// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DeletionMode string

// Enum values for DeletionMode
const (
	DeletionModeSoftDelete DeletionMode = "SoftDelete"
	DeletionModeHardDelete DeletionMode = "HardDelete"
)

// Values returns all known values for DeletionMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeletionMode) Values() []DeletionMode {
	return []DeletionMode{
		"SoftDelete",
		"HardDelete",
	}
}

type ExpirationTimeResponse string

// Enum values for ExpirationTimeResponse
const (
	ExpirationTimeResponseEnabled  ExpirationTimeResponse = "Enabled"
	ExpirationTimeResponseDisabled ExpirationTimeResponse = "Disabled"
)

// Values returns all known values for ExpirationTimeResponse. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExpirationTimeResponse) Values() []ExpirationTimeResponse {
	return []ExpirationTimeResponse{
		"Enabled",
		"Disabled",
	}
}

type TargetStore string

// Enum values for TargetStore
const (
	TargetStoreOnlineStore  TargetStore = "OnlineStore"
	TargetStoreOfflineStore TargetStore = "OfflineStore"
)

// Values returns all known values for TargetStore. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetStore) Values() []TargetStore {
	return []TargetStore{
		"OnlineStore",
		"OfflineStore",
	}
}

type TtlDurationUnit string

// Enum values for TtlDurationUnit
const (
	TtlDurationUnitSeconds TtlDurationUnit = "Seconds"
	TtlDurationUnitMinutes TtlDurationUnit = "Minutes"
	TtlDurationUnitHours   TtlDurationUnit = "Hours"
	TtlDurationUnitDays    TtlDurationUnit = "Days"
	TtlDurationUnitWeeks   TtlDurationUnit = "Weeks"
)

// Values returns all known values for TtlDurationUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TtlDurationUnit) Values() []TtlDurationUnit {
	return []TtlDurationUnit{
		"Seconds",
		"Minutes",
		"Hours",
		"Days",
		"Weeks",
	}
}
