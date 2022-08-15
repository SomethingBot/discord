package primitives

import "testing"

func TestPrivacyLevel_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		privacyLevel PrivacyLevel
		want         bool
	}{
		{
			name:         "PrivacyLevelNil",
			privacyLevel: PrivacyLevelNil,
			want:         false,
		},
		{
			name:         "PrivacyLevelPublic",
			privacyLevel: PrivacyLevelPublic,
			want:         true,
		},
		{
			name:         "PrivacyLevelGuildOnly",
			privacyLevel: PrivacyLevelGuildOnly,
			want:         true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.privacyLevel.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
