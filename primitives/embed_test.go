package primitives

import (
	"testing"
)

func TestEmbedType_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		e    EmbedType
		want bool
	}{
		{
			name: "EmbedTypeNil",
			e:    EmbedTypeNil,
			want: false,
		},
		{
			name: "EmbedTypeRich",
			e:    EmbedTypeRich,
			want: true,
		},
		{
			name: "EmbedTypeImage",
			e:    EmbedTypeImage,
			want: true,
		},
		{
			name: "EmbedTypeVideo",
			e:    EmbedTypeVideo,
			want: true,
		},
		{
			name: "EmbedTypeGifv",
			e:    EmbedTypeGifv,
			want: true,
		},
		{
			name: "EmbedTypeArticle",
			e:    EmbedTypeArticle,
			want: true,
		},
		{
			name: "EmbedTypeLink",
			e:    EmbedTypeLink,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			if got := tt.e.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmbed_IsValid(t *testing.T) {
	t.Parallel()
	type fields struct {
		Title       string
		Description string
		Footer      EmbedFooter
		Author      EmbedAuthor
		Fields      []EmbedField
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		//todo: invalid utf8 check?
		{
			name: "ErrorEmbedTooLarge",
			fields: fields{
				Title:       "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
				Description: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
				Footer: EmbedFooter{
					Text: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
				},
				Author: EmbedAuthor{
					Name: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
				},
				Fields: []EmbedField{
					{
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					},
					{
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					},
					{
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					},
					{
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					}, {
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					}, {
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					}, {
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					}, {
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					}, {
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					}, {
						Name:  "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					},
				},
			},
			wantErr: ErrorEmbedTooLarge,
		},
		{
			name: "ErrorEmbedTitleTooLarge",
			fields: fields{
				Title: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
			},
			wantErr: ErrorEmbedTitleTooLarge,
		},
		{
			name: "ErrorEmbedDescriptionTooLarge",
			fields: fields{
				Description: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
			},
			wantErr: ErrorEmbedDescriptionTooLarge,
		},
		{
			name: "ErrorEmbedFieldsTooLarge",
			fields: fields{
				Fields: []EmbedField{
					{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
				},
			},
			wantErr: ErrorEmbedFieldsTooLarge,
		},
		{
			name: "ErrorEmbedFieldNameTooLarge",
			fields: fields{
				Fields: []EmbedField{
					{
						Name: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					},
				},
			},
			wantErr: ErrorEmbedFieldNameTooLarge,
		},
		{
			name: "ErrorEmbedFieldValueTooLarge",
			fields: fields{
				Fields: []EmbedField{
					{
						Value: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
					},
				},
			},
			wantErr: ErrorEmbedFieldValueTooLarge,
		},
		{
			name: "ErrorEmbedFooterTextTooLarge",
			fields: fields{
				Footer: EmbedFooter{
					Text: "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",
				},
			},
			wantErr: ErrorEmbedFooterTextTooLarge,
		},
		{
			name: "ErrorEmbedAuthorNameTooLarge",
			fields: fields{
				Author: EmbedAuthor{
					Name: "QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2QaMrZimxAM9Q5DzKyquYkLSp3xNtug5WY3Nxc3PqEnYApkcu3VuhLgPgh9MQjMXuVRtbisgQmCaX4z3DtQetPKuPWJ9jfcXCkQKsKDyNDwy5UhJjPxEXfbugSckoiog79WYAqsCQdixRtzUVP9FcuStpwqQArNvsXQts59zQVnygLDPZaQRafam7FheiVjkWjHf73NFNNHityy9HvqVtdVv4ekyKpx5hTtm9DzEjimvyzEgmXjvKXLCghkHYzJd2",
				},
			},
			wantErr: ErrorEmbedAuthorNameTooLarge,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			e := Embed{
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				Footer:      tt.fields.Footer,
				Author:      tt.fields.Author,
				Fields:      tt.fields.Fields,
			}
			if err := e.IsValid(); err != tt.wantErr {
				t.Errorf("IsValid() wrong error = (%v), wantErr (%v)", err, tt.wantErr)
			}
		})
	}
}