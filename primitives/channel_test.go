package primitives

import "testing"

func TestSystemChannelFlag_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		systemChannelFlag SystemChannelFlag
		want              bool
	}{
		{
			name:              "SystemChannelFlagNil",
			systemChannelFlag: SystemChannelFlagNil,
			want:              false,
		},
		{
			name:              "SystemChannelFlagSuppressJoinNotifications",
			systemChannelFlag: SystemChannelFlagSuppressJoinNotifications,
			want:              true,
		},
		{
			name:              "SystemChannelFlagPremiumSubscriptions",
			systemChannelFlag: SystemChannelFlagPremiumSubscriptions,
			want:              true,
		},
		{
			name:              "SystemChannelFlagSuppressGuildReminderNotifications",
			systemChannelFlag: SystemChannelFlagSuppressGuildReminderNotifications,
			want:              true,
		},
		{
			name:              "SystemChannelFlagALL",
			systemChannelFlag: SystemChannelFlagALL,
			want:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.systemChannelFlag.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSystemChannelFlag_Contains(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		systemChannelFlag SystemChannelFlag
		flags             SystemChannelFlag
		want              bool
	}{
		{
			name:              "SystemChannelFlagNil",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagNil,
			want:              false,
		},
		{
			name:              "SystemChannelFlagSuppressJoinNotifications",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagSuppressJoinNotifications,
			want:              true,
		},
		{
			name:              "SystemChannelFlagPremiumSubscriptions",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagPremiumSubscriptions,
			want:              true,
		},
		{
			name:              "SystemChannelFlagSuppressGuildReminderNotifications",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagSuppressGuildReminderNotifications,
			want:              true,
		},
		{
			name:              "SystemChannelFlagALL",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagALL,
			want:              true,
		},
		{
			name:              "SystemChannelFlagSuppressGuildReminderNotifications|SystemChannelFlagPremiumSubscriptions",
			systemChannelFlag: SystemChannelFlagALL,
			flags:             SystemChannelFlagSuppressGuildReminderNotifications | SystemChannelFlagPremiumSubscriptions,
			want:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.systemChannelFlag.Contains(tt.flags); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannelType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		channelType ChannelType
		want        bool
	}{
		{
			name:        "ChannelTypeGuildText",
			channelType: ChannelTypeGuildText,
			want:        true,
		},
		{
			name:        "ChannelTypeDM",
			channelType: ChannelTypeDM,
			want:        true,
		},
		{
			name:        "ChannelTypeGuildVoice",
			channelType: ChannelTypeGuildVoice,
			want:        true,
		},
		{
			name:        "ChannelTypeGroupDM",
			channelType: ChannelTypeGroupDM,
			want:        true,
		},
		{
			name:        "ChannelTypeGuildCategory",
			channelType: ChannelTypeGuildCategory,
			want:        true,
		},
		{
			name:        "ChannelTypeGuildNews",
			channelType: ChannelTypeGuildNews,
			want:        true,
		},
		{
			name:        "ChannelTypeGuildStore",
			channelType: ChannelTypeGuildStore,
			want:        true,
		},
		{
			name:        "ChannelTypeGuildNewsThread",
			channelType: ChannelTypeGuildNewsThread,
			want:        true,
		},
		{
			name:        "ChannelTypeGuildPublicThread",
			channelType: ChannelTypeGuildPublicThread,
			want:        true,
		},
		{
			name:        "ChannelTypeGuildPrivateThread",
			channelType: ChannelTypeGuildPrivateThread,
			want:        true,
		},
		{
			name:        "ChannelTypeGuildStageVoice",
			channelType: ChannelTypeGuildStageVoice,
			want:        true,
		},
		{
			name:        "ChannelTypeINVALID",
			channelType: ChannelTypeINVALID,
			want:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.channelType.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVoiceQualityMode_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		voiceQualityMode VoiceQualityMode
		want             bool
	}{
		{
			name:             "VoiceQualityModeAuto",
			voiceQualityMode: VoiceQualityModeAuto,
			want:             true,
		},
		{
			name:             "VoiceQualityModeFull",
			voiceQualityMode: VoiceQualityModeFull,
			want:             true,
		},
		{
			name:             "VoiceQualityModeInvalid",
			voiceQualityMode: VoiceQualityModeInvalid,
			want:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.voiceQualityMode.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOverwriteIDType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		overwriteIDType OverwriteIDType
		want            bool
	}{
		{
			name:            "OverwriteIDTypeRole",
			overwriteIDType: OverwriteIDTypeRole,
			want:            true,
		},
		{
			name:            "OverwriteIDTypeUser",
			overwriteIDType: OverwriteIDTypeUser,
			want:            true,
		},
		{
			name:            "OverwriteIDTypeInvalid",
			overwriteIDType: 0313,
			want:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.overwriteIDType.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
