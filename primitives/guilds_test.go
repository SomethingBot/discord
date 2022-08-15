package primitives

import "testing"

func TestGuildScheduledEventPrivacyLevel_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                            string
		guildScheduledEventPrivacyLevel GuildScheduledEventPrivacyLevel
		want                            bool
	}{
		{
			name:                            "GuildScheduledEventPrivacyLevelGuildOnly",
			guildScheduledEventPrivacyLevel: GuildScheduledEventPrivacyLevelGuildOnly,
			want:                            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.guildScheduledEventPrivacyLevel.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuildScheduledEventStatus_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                      string
		guildScheduledEventStatus GuildScheduledEventStatus
		want                      bool
	}{
		{
			name:                      "GuildScheduledEventStatusNone",
			guildScheduledEventStatus: GuildScheduledEventStatusNone,
			want:                      false,
		},
		{
			name:                      "GuildScheduledEventStatusScheduled",
			guildScheduledEventStatus: GuildScheduledEventStatusScheduled,
			want:                      true,
		},
		{
			name:                      "GuildScheduledEventStatusActive",
			guildScheduledEventStatus: GuildScheduledEventStatusActive,
			want:                      true,
		},
		{
			name:                      "GuildScheduledEventStatusCompleted",
			guildScheduledEventStatus: GuildScheduledEventStatusCompleted,
			want:                      true,
		},
		{
			name:                      "GuildScheduledEventStatusCanceled",
			guildScheduledEventStatus: GuildScheduledEventStatusCanceled,
			want:                      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.guildScheduledEventStatus.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGuildScheduledEventEntityType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                          string
		guildScheduledEventEntityType GuildScheduledEventEntityType
		want                          bool
	}{
		{
			name:                          "GuildScheduledEventEntityTypeNone",
			guildScheduledEventEntityType: GuildScheduledEventEntityTypeNone,
			want:                          false,
		},
		{
			name:                          "GuildScheduledEventEntityTypeStageInstance",
			guildScheduledEventEntityType: GuildScheduledEventEntityTypeStageInstance,
			want:                          true,
		},
		{
			name:                          "GuildScheduledEventEntityTypeVoice",
			guildScheduledEventEntityType: GuildScheduledEventEntityTypeVoice,
			want:                          true,
		},
		{
			name:                          "GuildScheduledEventEntityTypeExternal",
			guildScheduledEventEntityType: GuildScheduledEventEntityTypeExternal,
			want:                          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.guildScheduledEventEntityType.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
