package primitives

import "testing"

func TestPermissionFlag_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		permissionFlag PermissionFlag
		want           bool
	}{
		{
			name:           "PermissionFlagNil",
			permissionFlag: PermissionFlagNil,
			want:           false,
		},
		{
			name:           "PermissionFlagCreateInstantInvite",
			permissionFlag: PermissionFlagCreateInstantInvite,
			want:           true,
		},
		{
			name:           "PermissionFlagKickMembers",
			permissionFlag: PermissionFlagKickMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagBanMembers",
			permissionFlag: PermissionFlagBanMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagAdministrator",
			permissionFlag: PermissionFlagAdministrator,
			want:           true,
		},
		{
			name:           "PermissionFlagManageChannels",
			permissionFlag: PermissionFlagManageChannels,
			want:           true,
		},
		{
			name:           "PermissionFlagManageGuild",
			permissionFlag: PermissionFlagManageGuild,
			want:           true,
		},
		{
			name:           "PermissionFlagAddReactions",
			permissionFlag: PermissionFlagAddReactions,
			want:           true,
		},
		{
			name:           "PermissionFlagViewAuditLog",
			permissionFlag: PermissionFlagViewAuditLog,
			want:           true,
		},
		{
			name:           "PermissionFlagPrioritySpeaker",
			permissionFlag: PermissionFlagPrioritySpeaker,
			want:           true,
		},
		{
			name:           "PermissionFlagStream",
			permissionFlag: PermissionFlagStream,
			want:           true,
		},
		{
			name:           "PermissionFlagViewChannel",
			permissionFlag: PermissionFlagViewChannel,
			want:           true,
		},
		{
			name:           "PermissionFlagSendMessages",
			permissionFlag: PermissionFlagSendMessages,
			want:           true,
		},
		{
			name:           "PermissionFlagSendTTSMessage",
			permissionFlag: PermissionFlagSendTTSMessage,
			want:           true,
		},
		{
			name:           "PermissionFlagManageMessages",
			permissionFlag: PermissionFlagManageMessages,
			want:           true,
		},
		{
			name:           "PermissionFlagEmbedLinks",
			permissionFlag: PermissionFlagEmbedLinks,
			want:           true,
		},
		{
			name:           "PermissionFlagAttachFiles",
			permissionFlag: PermissionFlagAttachFiles,
			want:           true,
		},
		{
			name:           "PermissionFlagReadMessageHistory",
			permissionFlag: PermissionFlagReadMessageHistory,
			want:           true,
		},
		{
			name:           "PermissionFlagMentionEveryone",
			permissionFlag: PermissionFlagMentionEveryone,
			want:           true,
		},
		{
			name:           "PermissionFlagUseExternalEmojis",
			permissionFlag: PermissionFlagUseExternalEmojis,
			want:           true,
		},
		{
			name:           "PermissionFlagViewGuildInsights",
			permissionFlag: PermissionFlagViewGuildInsights,
			want:           true,
		},
		{
			name:           "PermissionFlagConnect",
			permissionFlag: PermissionFlagConnect,
			want:           true,
		},
		{
			name:           "PermissionFlagSpeak",
			permissionFlag: PermissionFlagSpeak,
			want:           true,
		},
		{
			name:           "PermissionFlagMuteMembers",
			permissionFlag: PermissionFlagMuteMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagDeafenMembers",
			permissionFlag: PermissionFlagDeafenMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagMoveMembers",
			permissionFlag: PermissionFlagMoveMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagUseVoiceActivityDetection",
			permissionFlag: PermissionFlagUseVoiceActivityDetection,
			want:           true,
		},
		{
			name:           "PermissionFlagChangeNickname",
			permissionFlag: PermissionFlagChangeNickname,
			want:           true,
		},
		{
			name:           "PermissionFlagManageNicknames",
			permissionFlag: PermissionFlagManageNicknames,
			want:           true,
		},
		{
			name:           "PermissionFlagManageRoles",
			permissionFlag: PermissionFlagManageRoles,
			want:           true,
		},
		{
			name:           "PermissionFlagManageWebhooks",
			permissionFlag: PermissionFlagManageWebhooks,
			want:           true,
		},
		{
			name:           "PermissionFlagManageEmojisAndStickers",
			permissionFlag: PermissionFlagManageEmojisAndStickers,
			want:           true,
		},
		{
			name:           "PermissionFlagUseApplicationCommands",
			permissionFlag: PermissionFlagUseApplicationCommands,
			want:           true,
		},
		{
			name:           "PermissionFlagRequestToSpeak",
			permissionFlag: PermissionFlagRequestToSpeak,
			want:           true,
		},
		{
			name:           "PermissionFlagSkippedNotReal",
			permissionFlag: PermissionFlagSkippedNotReal,
			want:           false,
		},
		{
			name:           "PermissionFlagManageThreads",
			permissionFlag: PermissionFlagManageThreads,
			want:           true,
		},
		{
			name:           "PermissionFlagCreatePublicThreads",
			permissionFlag: PermissionFlagCreatePublicThreads,
			want:           true,
		},
		{
			name:           "PermissionFlagCreatePrivateThreads",
			permissionFlag: PermissionFlagCreatePrivateThreads,
			want:           true,
		},
		{
			name:           "PermissionFlagUseExternalStickers",
			permissionFlag: PermissionFlagUseExternalStickers,
			want:           true,
		},
		{
			name:           "PermissionFlagSendMessageInThreads",
			permissionFlag: PermissionFlagSendMessageInThreads,
			want:           true,
		},
		{
			name:           "PermissionFlagStartEmbeddedActivities",
			permissionFlag: PermissionFlagStartEmbeddedActivities,
			want:           true,
		},
		{
			name:           "PermissionModerateMembers",
			permissionFlag: PermissionModerateMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagALL",
			permissionFlag: PermissionFlagALL,
			want:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.permissionFlag.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermissionFlag_Contains(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		permissionFlag PermissionFlag
		flags          PermissionFlag
		want           bool
	}{
		{
			name:           "PermissionFlagNil",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagNil,
			want:           true,
		},
		{
			name:           "PermissionFlagCreateInstantInvite",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagCreateInstantInvite,
			want:           true,
		},
		{
			name:           "PermissionFlagKickMembers",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagKickMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagBanMembers",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagBanMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagAdministrator",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagAdministrator,
			want:           true,
		},
		{
			name:           "PermissionFlagManageChannels",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagManageChannels,
			want:           true,
		},
		{
			name:           "PermissionFlagManageGuild",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagManageGuild,
			want:           true,
		},
		{
			name:           "PermissionFlagAddReactions",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagAddReactions,
			want:           true,
		},
		{
			name:           "PermissionFlagViewAuditLog",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagViewAuditLog,
			want:           true,
		},
		{
			name:           "PermissionFlagPrioritySpeaker",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagPrioritySpeaker,
			want:           true,
		},
		{
			name:           "PermissionFlagStream",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagStream,
			want:           true,
		},
		{
			name:           "PermissionFlagViewChannel",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagViewChannel,
			want:           true,
		},
		{
			name:           "PermissionFlagSendMessages",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagSendMessages,
			want:           true,
		},
		{
			name:           "PermissionFlagSendTTSMessage",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagSendTTSMessage,
			want:           true,
		},
		{
			name:           "PermissionFlagManageMessages",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagManageMessages,
			want:           true,
		},
		{
			name:           "PermissionFlagEmbedLinks",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagEmbedLinks,
			want:           true,
		},
		{
			name:           "PermissionFlagAttachFiles",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagAttachFiles,
			want:           true,
		},
		{
			name:           "PermissionFlagReadMessageHistory",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagReadMessageHistory,
			want:           true,
		},
		{
			name:           "PermissionFlagMentionEveryone",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagMentionEveryone,
			want:           true,
		},
		{
			name:           "PermissionFlagUseExternalEmojis",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagUseExternalEmojis,
			want:           true,
		},
		{
			name:           "PermissionFlagViewGuildInsights",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagViewGuildInsights,
			want:           true,
		},
		{
			name:           "PermissionFlagConnect",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagConnect,
			want:           true,
		},
		{
			name:           "PermissionFlagSpeak",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagSpeak,
			want:           true,
		},
		{
			name:           "PermissionFlagMuteMembers",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagMuteMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagDeafenMembers",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagDeafenMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagMoveMembers",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagMoveMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagUseVoiceActivityDetection",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagUseVoiceActivityDetection,
			want:           true,
		},
		{
			name:           "PermissionFlagChangeNickname",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagChangeNickname,
			want:           true,
		},
		{
			name:           "PermissionFlagManageNicknames",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagManageNicknames,
			want:           true,
		},
		{
			name:           "PermissionFlagManageRoles",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagManageRoles,
			want:           true,
		},
		{
			name:           "PermissionFlagManageWebhooks",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagManageWebhooks,
			want:           true,
		},
		{
			name:           "PermissionFlagManageEmojisAndStickers",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagManageEmojisAndStickers,
			want:           true,
		},
		{
			name:           "PermissionFlagUseApplicationCommands",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagUseApplicationCommands,
			want:           true,
		},
		{
			name:           "PermissionFlagRequestToSpeak",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagRequestToSpeak,
			want:           true,
		},
		{
			name:           "PermissionFlagSkippedNotReal",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagSkippedNotReal,
			want:           true,
		},
		{
			name:           "PermissionFlagManageThreads",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagManageThreads,
			want:           true,
		},
		{
			name:           "PermissionFlagCreatePublicThreads",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagCreatePublicThreads,
			want:           true,
		},
		{
			name:           "PermissionFlagCreatePrivateThreads",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagCreatePrivateThreads,
			want:           true,
		},
		{
			name:           "PermissionFlagUseExternalStickers",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagUseExternalStickers,
			want:           true,
		},
		{
			name:           "PermissionFlagSendMessageInThreads",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagSendMessageInThreads,
			want:           true,
		},
		{
			name:           "PermissionFlagStartEmbeddedActivities",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagStartEmbeddedActivities,
			want:           true,
		},
		{
			name:           "PermissionModerateMembers",
			permissionFlag: PermissionModerateMembers,
			want:           true,
		},
		{
			name:           "PermissionFlagALL",
			permissionFlag: PermissionFlagALL,
			flags:          PermissionFlagALL,
			want:           true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.permissionFlag.Contains(tt.flags); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
