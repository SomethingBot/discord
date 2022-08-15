package primitives

import "testing"

func TestTeamMembershipState_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                string
		teamMembershipState TeamMembershipState
		want                bool
	}{
		{
			name:                "TeamMembershipStateInvalid",
			teamMembershipState: TeamMembershipStateInvalid,
			want:                false,
		},
		{
			name:                "TeamMembershipStateInvited",
			teamMembershipState: TeamMembershipStateInvited,
			want:                true,
		},
		{
			name:                "TeamMembershipStateAccepted",
			teamMembershipState: TeamMembershipStateAccepted,
			want:                true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.teamMembershipState.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
