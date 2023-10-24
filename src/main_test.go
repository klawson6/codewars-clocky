package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDegreeToTimeInt(t *testing.T) {
	type args struct {
		degrees int
	}
	tests := []struct {
		name          string
		args          []args
		expectedTimes []string
	}{
		{
			name:          "When input is 0, time resolves to 12:00",
			args:          []args{{degrees: 0}},
			expectedTimes: []string{"12:00"},
		},
		{
			name:          "When input is 360, time resolves to 12:00",
			args:          []args{{degrees: 360}},
			expectedTimes: []string{"12:00"},
		},
		{
			name:          "When 0 < input < 360, time resolves correctly",
			args:          []args{{degrees: 1}, {degrees: 20}, {degrees: 35}, {degrees: 67}, {degrees: 90}, {degrees: 111}, {degrees: 122}, {degrees: 163}, {degrees: 180}, {degrees: 208}, {degrees: 223}, {degrees: 250}, {degrees: 270}, {degrees: 285}, {degrees: 302}, {degrees: 330}, {degrees: 359}},
			expectedTimes: []string{"12:02", "12:40", "01:10", "02:14", "03:00", "03:42", "04:04", "05:26", "06:00", "06:56", "07:26", "08:20", "09:00", "09:30", "10:04", "11:00", "11:58"},
		},
		{
			name:          "When 360 < input < 0, time resolves correctly",
			args:          []args{{degrees: -1}, {degrees: -20}, {degrees: -35}, {degrees: -67}, {degrees: -90}, {degrees: -111}, {degrees: -122}, {degrees: -163}, {degrees: -180}, {degrees: -208}, {degrees: -223}, {degrees: -250}, {degrees: -270}, {degrees: -285}, {degrees: -302}, {degrees: -330}, {degrees: -359}},
			expectedTimes: []string{"11:58", "11:20", "10:50", "09:46", "09:00", "08:18", "07:56", "06:34", "06:00", "05:04", "04:34", "03:40", "03:00", "02:30", "01:56", "01:00", "12:02"},
		},
		{
			name:          "When input > 360, time resolves correctly",
			args:          []args{{degrees: 720}, {degrees: 1111}, {degrees: 245774}},
			expectedTimes: []string{"12:00", "01:02", "08:28"},
		},
		{
			name:          "When input < 0, time resolves correctly",
			args:          []args{{degrees: -720}, {degrees: -1111}, {degrees: -245774}},
			expectedTimes: []string{"12:00", "10:58", "03:32"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, in := range tt.args {
				result := DegreeToTimeInt(in.degrees)
				assert.Equal(t, tt.expectedTimes[i], result)
			}
		})
	}
}

func TestDegreeToTimeFloat64(t *testing.T) {
	type args struct {
		degrees float64
	}
	tests := []struct {
		name          string
		args          []args
		expectedTimes []string
	}{
		{
			name:          "When input is 0, time resolves to 12:00",
			args:          []args{{degrees: 0}},
			expectedTimes: []string{"12:00"},
		},
		{
			name:          "When input is 360, time resolves to 12:00",
			args:          []args{{degrees: 360}},
			expectedTimes: []string{"12:00"},
		},
		//{
		//	name:          "When 0 < input < 360, time resolves correctly",
		//	args:          []args{{degrees: 0.25}, {degrees: 0.74}, {degrees: 0.75}, {degrees: 133.456}, {degrees: 279.60733}, {degrees: 354.9993333}},
		//	expectedTimes: []string{"12:01", "12:01", "12:02", "04:27", "09:19", "11:50"},
		//},
		//{
		//	name:          "When 360 < input < 0, time resolves correctly",
		//	args:          []args{{degrees: -0.25}, {degrees: -0.74}, {degrees: -0.75}, {degrees: -133.456}, {degrees: -279.60733}, {degrees: -354.9993333}},
		//	expectedTimes: []string{"11:59", "11:59", "11.58", "07:33", "02:41", "01:10"},
		//},
		{
			name:          "When input > 360, time resolves correctly",
			args:          []args{{degrees: 720}, {degrees: 1111}, {degrees: 245774}},
			expectedTimes: []string{"12:00", "01:02", "08:28"},
		},
		//{
		//	name:          "When input < 0, time resolves correctly",
		//	args:          []args{{degrees: -720}, {degrees: -1111}, {degrees: -245774}},
		//	expectedTimes: []string{"12:00", "10:58", "03:32"},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, in := range tt.args {
				result := DegreeToTimeFloat64(in.degrees)
				assert.Equal(t, tt.expectedTimes[i], result)
			}
		})
	}
}
