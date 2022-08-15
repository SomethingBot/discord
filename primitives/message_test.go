package primitives

import "testing"

func TestMessageType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		messageType MessageType
		want        bool
	}{
		{
			name:        "MessageTypeDefault",
			messageType: MessageTypeDefault,
			want:        true,
		},
		{
			name:        "MessageTypeRecipientAdd",
			messageType: MessageTypeRecipientAdd,
			want:        true,
		},
		{
			name:        "MessageTypeRecipientRemove",
			messageType: MessageTypeRecipientRemove,
			want:        true,
		},
		{
			name:        "MessageTypeCall",
			messageType: MessageTypeCall,
			want:        true,
		},
		{
			name:        "MessageTypeChannelNameChange",
			messageType: MessageTypeChannelNameChange,
			want:        true,
		},
		{
			name:        "MessageTypeChannelIconChange",
			messageType: MessageTypeChannelIconChange,
			want:        true,
		},
		{
			name:        "MessageTypeChannelPinnedMessage",
			messageType: MessageTypeChannelPinnedMessage,
			want:        true,
		},
		{
			name:        "MessageTypeGuildMemberJoin",
			messageType: MessageTypeGuildMemberJoin,
			want:        true,
		},
		{
			name:        "MessageTypeUserPremiumGuildSubscription",
			messageType: MessageTypeUserPremiumGuildSubscription,
			want:        true,
		},
		{
			name:        "MessageTypeUserPremiumGuildSubscriptionTier1",
			messageType: MessageTypeUserPremiumGuildSubscriptionTier1,
			want:        true,
		},
		{
			name:        "MessageTypeUserPremiumGuildSubscriptionTier2",
			messageType: MessageTypeUserPremiumGuildSubscriptionTier2,
			want:        true,
		},
		{
			name:        "MessageTypeUserPremiumGuildSubscriptionTier3",
			messageType: MessageTypeUserPremiumGuildSubscriptionTier3,
			want:        true,
		},
		{
			name:        "MessageTypeChannelFollowAdd",
			messageType: MessageTypeChannelFollowAdd,
			want:        true,
		},
		{
			name:        "MessageTypeGuildDiscoveryGracePeriodInitialWarning",
			messageType: MessageTypeGuildDiscoveryGracePeriodInitialWarning,
			want:        true,
		},
		{
			name:        "MessageTypeGuildDiscoveryGracePeriodFinalWarning",
			messageType: MessageTypeGuildDiscoveryGracePeriodFinalWarning,
			want:        true,
		},
		{
			name:        "MessageTypeThreadCreated",
			messageType: MessageTypeThreadCreated,
			want:        true,
		},
		{
			name:        "MessageTypeReply",
			messageType: MessageTypeReply,
			want:        true,
		},
		{
			name:        "MessageTypeChatInputCommand",
			messageType: MessageTypeChatInputCommand,
			want:        true,
		},
		{
			name:        "MessageTypeThreadStarterMessage",
			messageType: MessageTypeThreadStarterMessage,
			want:        true,
		},
		{
			name:        "MessageTypeGuildInviteReminder",
			messageType: MessageTypeGuildInviteReminder,
			want:        true,
		},
		{
			name:        "MessageTypeContextMenuCommand",
			messageType: MessageTypeContextMenuCommand,
			want:        true,
		},
		{
			name:        "MessageTypeInvalid",
			messageType: MessageTypeInvalid,
			want:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.messageType.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageActivityType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		m    MessageActivityType
		want bool
	}{
		{
			name: "MessageActivityTypeNil",
			m:    MessageActivityTypeNil,
			want: false,
		},
		{
			name: "MessageActivityTypeJoin",
			m:    MessageActivityTypeJoin,
			want: true,
		},
		{
			name: "MessageActivityTypeSpectate",
			m:    MessageActivityTypeSpectate,
			want: true,
		},
		{
			name: "MessageActivityTypeListen",
			m:    MessageActivityTypeListen,
			want: true,
		},
		{
			name: "MessageActivityTypeJoinRequest",
			m:    MessageActivityTypeJoinRequest,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.m.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
