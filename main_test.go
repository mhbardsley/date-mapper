package main

import (
	"testing"
	"time"
)

func TestMapDate(t *testing.T) {
	tests := []struct {
		x, a, b, c, d time.Time
		want          time.Time
		wantErr       bool
	}{
		{
			x:   time.Date(2020, 6, 15, 0, 0, 0, 0, time.UTC),
			a:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			b:   time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			c:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
			want: time.Date(2021, 6, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			x:   time.Date(2020, 6, 15, 0, 0, 0, 0, time.UTC),
			a:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			b:   time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			c:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   time.Date(2021, 6, 30, 0, 0, 0, 0, time.UTC),
			want: time.Date(2021, 3, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			x:   time.Date(2020, 6, 15, 0, 0, 0, 0, time.UTC),
			a:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			b:   time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			c:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
			want: time.Date(2021, 6, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			x:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			a:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			b:   time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			c:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
			want: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			x:   time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			a:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			b:   time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			c:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			d:   time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
			want: time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			x:     time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC),
			a:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			b:     time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			c:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			d:     time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
			want:  time.Time{},
			wantErr: true,
		},
		{
			x:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			a:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			b:     time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC),
			c:     time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			d:     time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC),
			want:  time.Time{},
			wantErr: true,
		},
	}

	for _, test := range tests {
		// Call mapDate with the test arguments
		got, err := mapDate(test.x, test.a, test.b, test.c, test.d)

		// Check for errors
		if err != nil {
			if !test.wantErr {
				t.Errorf("unexpected error: %v", err)
			}
			continue
		} else if test.wantErr {
			t.Error("expected error but got none")
			continue
		}

		// Compare the date parts of the results and expected values
		if got.Year() != test.want.Year() || got.Month() != test.want.Month() || got.Day() != test.want.Day() {
			t.Errorf("unexpected result: want %v, got %v", test.want.Format("2006-01-02"), got.Format("2006-01-02"))
		}
	}
}