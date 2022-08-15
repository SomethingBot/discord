package primitives

import "testing"

func TestPresenceStatus_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		status PresenceStatus
		want   bool
	}{
		{
			name:   "PresenceStatusIdle",
			status: PresenceStatusIdle,
			want:   true,
		},
		{
			name:   "PresenceStatusDnd",
			status: PresenceStatusDnd,
			want:   true,
		},
		{
			name:   "PresenceStatusOnline",
			status: PresenceStatusOnline,
			want:   true,
		},
		{
			name:   "PresenceStatusOffline",
			status: PresenceStatusOffline,
			want:   true,
		},
		{
			name:   "PresenceStatusNil",
			status: PresenceStatusNil,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.status.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		activityType ActivityType
		want         bool
	}{
		{
			name:         "ActivityTypeGame",
			activityType: ActivityTypeGame,
			want:         true,
		},
		{
			name:         "ActivityTypeStreaming",
			activityType: ActivityTypeStreaming,
			want:         true,
		},
		{
			name:         "ActivityTypeListening",
			activityType: ActivityTypeListening,
			want:         true,
		},
		{
			name:         "ActivityTypeWatching",
			activityType: ActivityTypeWatching,
			want:         true,
		},
		{
			name:         "ActivityTypeCustom",
			activityType: ActivityTypeCustom,
			want:         true,
		},
		{
			name:         "ActivityTypeCompeting",
			activityType: ActivityTypeCompeting,
			want:         true,
		},
		{
			name:         "ActivityTypeInvalid",
			activityType: ActivityTypeInvalid,
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.activityType.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityFlags_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		activityFlags ActivityFlag
		want          bool
	}{
		{
			name:          "ActivityFlagInstance",
			activityFlags: ActivityFlagInstance,
			want:          true,
		},
		{
			name:          "ActivityFlagJoin",
			activityFlags: ActivityFlagJoin,
			want:          true,
		},
		{
			name:          "ActivityFlagSpectate",
			activityFlags: ActivityFlagSpectate,
			want:          true,
		},
		{
			name:          "ActivityFlagJoinRequest",
			activityFlags: ActivityFlagJoinRequest,
			want:          true,
		},
		{
			name:          "ActivityFlagSync",
			activityFlags: ActivityFlagSync,
			want:          true,
		},
		{
			name:          "ActivityFlagPlay",
			activityFlags: ActivityFlagPlay,
			want:          true,
		},
		{
			name:          "ActivityFlagAll",
			activityFlags: ActivityFlagAll,
			want:          true,
		},
		{
			name:          "ActivityFlagSync|ActivityFlagPlay",
			activityFlags: ActivityFlagSync | ActivityFlagPlay,
			want:          true,
		},
		{
			name:          "ActivityFlagInvalid",
			activityFlags: ActivityFlagInvalid,
			want:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.activityFlags.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityFlags_Contains(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		activityFlags ActivityFlag
		want          bool
	}{
		{
			name:          "ActivityFlagInstance",
			activityFlags: ActivityFlagInstance,
			want:          true,
		},
		{
			name:          "ActivityFlagJoin",
			activityFlags: ActivityFlagJoin,
			want:          true,
		},
		{
			name:          "ActivityFlagSpectate",
			activityFlags: ActivityFlagSpectate,
			want:          true,
		},
		{
			name:          "ActivityFlagJoinRequest",
			activityFlags: ActivityFlagJoinRequest,
			want:          true,
		},
		{
			name:          "ActivityFlagSync",
			activityFlags: ActivityFlagSync,
			want:          true,
		},
		{
			name:          "ActivityFlagPlay",
			activityFlags: ActivityFlagPlay,
			want:          true,
		},
		{
			name:          "ActivityFlagInvalid",
			activityFlags: ActivityFlagInvalid,
			want:          false,
		},
		{
			name:          "ActivityFlagPlay|ActivityFlagSync",
			activityFlags: ActivityFlagPlay | ActivityFlagSync,
			want:          true,
		},
		{
			name:          "ActivityFlagAll",
			activityFlags: ActivityFlagAll,
			want:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := ActivityFlagAll.Contains(tt.activityFlags); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
