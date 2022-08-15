package primitives

import "testing"

func TestUserFlag_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		userFlag UserFlag
		want     bool
	}{
		{
			name:     "UserFlagNil",
			userFlag: UserFlagNil,
			want:     false,
		},
		{
			name:     "UserFlagDiscordEmployee",
			userFlag: UserFlagDiscordEmployee,
			want:     true,
		},
		{
			name:     "UserFlagPartneredServerOwner",
			userFlag: UserFlagPartneredServerOwner,
			want:     true,
		},
		{
			name:     "UserFlagHypeSquadEvents",
			userFlag: UserFlagHypeSquadEvents,
			want:     true,
		},
		{
			name:     "UserFlagBugHunterLevel1",
			userFlag: UserFlagBugHunterLevel1,
			want:     true,
		},
		{
			name:     "UserFlagHouseBravery",
			userFlag: UserFlagHouseBravery,
			want:     true,
		},
		{
			name:     "UserFlagHouseBrilliance",
			userFlag: UserFlagHouseBrilliance,
			want:     true,
		},
		{
			name:     "UserFlagHouseBalance",
			userFlag: UserFlagHouseBalance,
			want:     true,
		},
		{
			name:     "UserFlagEarlySupporter",
			userFlag: UserFlagEarlySupporter,
			want:     true,
		},
		{
			name:     "UserFlagTeamUser",
			userFlag: UserFlagTeamUser,
			want:     true,
		},
		{
			name:     "UserFlagBugHunterLevel2",
			userFlag: UserFlagBugHunterLevel2,
			want:     true,
		},
		{
			name:     "UserFlagVerifiedBot",
			userFlag: UserFlagVerifiedBot,
			want:     true,
		},
		{
			name:     "UserFlagEarlyVerifiedBotDeveloper",
			userFlag: UserFlagEarlyVerifiedBotDeveloper,
			want:     true,
		},
		{
			name:     "UserFlagDiscordCertifiedModerator",
			userFlag: UserFlagDiscordCertifiedModerator,
			want:     true,
		},
		{
			name:     "UserFlagAll",
			userFlag: UserFlagAll,
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.userFlag.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserFlag_Contains(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		userFlag UserFlag
		flags    UserFlag
		want     bool
	}{
		{
			name:     "UserFlagPartneredServerOwner|UserFlagHypeSquadEvents",
			userFlag: UserFlagAll,
			flags:    UserFlagPartneredServerOwner | UserFlagHypeSquadEvents,
			want:     true,
		},
		{
			name:     "UserFlagAll",
			userFlag: UserFlagAll,
			flags:    UserFlagAll,
			want:     true,
		},
		{
			name:     "UserFlagBugHunterLevel1",
			userFlag: UserFlagBugHunterLevel1,
			flags:    UserFlagPartneredServerOwner,
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.userFlag.Contains(tt.flags); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPremiumType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		premiumType PremiumType
		want        bool
	}{
		{
			name:        "PremiumTypeNil",
			premiumType: PremiumTypeNil,
			want:        false,
		},
		{
			name:        "PremiumTypeNitroClassic",
			premiumType: PremiumTypeNitroClassic,
			want:        true,
		},
		{
			name:        "PremiumTypeNitro",
			premiumType: PremiumTypeNitro,
			want:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.premiumType.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
