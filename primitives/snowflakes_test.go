package primitives

import (
	"reflect"
	"testing"
	"time"
)

func TestSnowflake_IsValid(t *testing.T) {
	tests := []struct {
		name      string
		snowflake Snowflake
		want      bool
	}{
		{
			name:      "from docs",
			snowflake: 175928847299117063,
			want:      true,
		},
		{
			name:      "nil snowflake",
			snowflake: 0,
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.snowflake.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_Timestamp(t *testing.T) {
	tests := []struct {
		name      string
		snowflake Snowflake
		want      time.Time
	}{
		{
			name:      "from docs",
			snowflake: 175928847299117063,
			want:      time.Date(2016, time.April, 30, 11, 18, 25, 796*1e6, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.snowflake.Timestamp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_InternalWorkerID(t *testing.T) {
	tests := []struct {
		name      string
		snowflake Snowflake
		want      uint8
	}{
		{
			name:      "from docs",
			snowflake: 175928847299117063,
			want:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.snowflake.InternalWorkerID(); got != tt.want {
				t.Errorf("InternalWorkerID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_InternalProcessID(t *testing.T) {
	tests := []struct {
		name      string
		snowflake Snowflake
		want      uint8
	}{
		{
			name:      "from docs",
			snowflake: 175928847299117063,
			want:      0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.snowflake.InternalProcessID(); got != tt.want {
				t.Errorf("InternalProcessID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSnowflake_Increment(t *testing.T) {
	tests := []struct {
		name      string
		snowflake Snowflake
		want      uint16
	}{
		{
			name:      "from docs",
			snowflake: 175928847299117063,
			want:      7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.snowflake.Increment(); got != tt.want {
				t.Errorf("Increment() = %v, want %v", got, tt.want)
			}
		})
	}
}
