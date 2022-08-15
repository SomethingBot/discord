package primitives

import "testing"

func TestInviteTargetType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		i    InviteTargetType
		want bool
	}{
		{
			name: "InviteTargetTypeNil",
			i:    InviteTargetTypeNil,
			want: false,
		},
		{
			name: "InviteTargetTypeStream",
			i:    InviteTargetTypeStream,
			want: true,
		},
		{
			name: "InviteTargetTypeEmbeddedApplication",
			i:    InviteTargetTypeEmbeddedApplication,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.i.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
