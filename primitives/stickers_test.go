package primitives

import "testing"

func TestStickerType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		stickerType StickerType
		want        bool
	}{
		{
			name:        "StickerTypeNil",
			stickerType: StickerTypeNil,
			want:        false,
		},
		{
			name:        "StickerTypeStandard",
			stickerType: StickerTypeStandard,
			want:        true,
		},
		{
			name:        "StickerTypeGuild",
			stickerType: StickerTypeGuild,
			want:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.stickerType.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStickerFormatType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		stickerFormatType StickerFormatType
		want              bool
	}{
		{
			name:              "StickerFormatTypeNil",
			stickerFormatType: StickerFormatTypeNil,
			want:              false,
		},
		{
			name:              "StickerFormatTypePNG",
			stickerFormatType: StickerFormatTypePNG,
			want:              true,
		},
		{
			name:              "StickerFormatTypeAPNG",
			stickerFormatType: StickerFormatTypeAPNG,
			want:              true,
		},
		{
			name:              "StickerFormatTypeLOTTIE",
			stickerFormatType: StickerFormatTypeLOTTIE,
			want:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.stickerFormatType.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
