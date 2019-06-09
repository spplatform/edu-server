package server

import "testing"

func TestCalculateHash(t *testing.T) {
	tests := []struct {
		name string
		base string
		want string
	}{
		{
			name: "one",
			base: "one",
			want: "fe05bcdcdc4928012781a5f1a2a77cbb5398e106",
		},
		{
			name: "two",
			base: "two",
			want: "ad782ecdac770fc6eb9a62e44f90873fb97fb26b",
		},
		{
			name: "some",
			base: "5fgioklw4yjh0t4rhwy89sgjeio",
			want: "776be58ec8d1739c6638db01c21bb4f4e226f392",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateHash(tt.base); got != tt.want {
				t.Errorf("wrong hash %v, want %v", got, tt.want)
			}
		})
	}
}
