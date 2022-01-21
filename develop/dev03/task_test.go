package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sortLines(t *testing.T) {
	type args struct {
		input      string
		parameters Flag
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				input:      "bb kk\naa ll\ncc",
				parameters: Flag{},
			},
			want: []string{"aa ll", "bb kk", "cc"},
		},
		{
			name: "Test 2",
			args: args{
				input:      "bb kk\naa ll\ncc aa",
				parameters: Flag{columnNumber: 2},
			},
			want: []string{"cc aa", "bb kk", "aa ll"},
		},
		{
			name: "Test 3",
			args: args{
				input:      "bb kk aa\naa ll aa\ncc aa aa",
				parameters: Flag{columnNumber: 3},
			},
			want: []string{"aa ll aa", "bb kk aa", "cc aa aa"},
		},
		{
			name: "Test 4",
			args: args{
				input:      "10 bb kk aa\n2 aa ll aa\n3 cc aa aa",
				parameters: Flag{isNumericSort: true},
			},
			want: []string{"2 aa ll aa", "3 cc aa aa", "10 bb kk aa"},
		},
		{
			name: "Test 5",
			args: args{
				input:      "bb kk aa\naa ll aa\ncc aa aa",
				parameters: Flag{isReverse: true},
			},
			want: []string{"cc aa aa", "bb kk aa", "aa ll aa"},
		},
		{
			name: "Test 6",
			args: args{
				input:      "bb kk aa\naa ll aa\ncc aa aa",
				parameters: Flag{columnNumber: 3, isReverse: true},
			},
			want: []string{"cc aa aa", "bb kk aa", "aa ll aa"},
		},
		{
			name: "Test 7",
			args: args{
				input:      "bb kk aa\naa ll aa\ncc aa aa\nbb kk aa",
				parameters: Flag{isOnlyUnique: true},
			},
			want: []string{"aa ll aa", "bb kk aa", "cc aa aa"},
		},
		{
			name: "Test 8",
			args: args{
				input:      "10.1 bb kk aa\n2.3 aa ll aa\n3.5 cc aa aa",
				parameters: Flag{isNumericSort: true},
			},
			want: []string{"2.3 aa ll aa", "3.5 cc aa aa", "10.1 bb kk aa"},
		},
		{
			name: "Test 8",
			args: args{
				input:      "10.1 bb kk aa\n2.3 aa ll aa\n3.5 cc aa aa\n10.1 bb kk aa",
				parameters: Flag{isOnlyUnique: true, isNumericSort: true},
			},
			want: []string{"2.3 aa ll aa", "3.5 cc aa aa", "10.1 bb kk aa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortLines(tt.args.input, tt.args.parameters)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
