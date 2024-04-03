package utils

import "testing"

func TestFeatureFlagRateOriginCodeField(t *testing.T) {
	shouldSeeFeature := Channel{18, 1940}
	shouldSeeFeatureAsWell := Channel{27, 2}
	shouldNotSeeFeature := Channel{17, 1939}

	testCases := []struct {
		name     string
		simulate Channel
		expected bool
	}{
		{
			name:     "Channel Dynamics see the field rateOrigincCode",
			simulate: shouldSeeFeature,
			expected: true,
		},
		{
			name:     "Channel Dynamics PJ see the field rateOriginCode",
			simulate: shouldSeeFeatureAsWell,
			expected: true,
		},
		{
			name:     "Channel WithoutPermission don't see the field rateOriginCode",
			simulate: shouldNotSeeFeature,
			expected: false,
		},
	}

	ftManager := NewFeatureFlagsManager()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ftManager.RateOriginCodeField(tc.simulate)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v for simulate request %+v", tc.expected, result, tc.simulate)
			}
		})
	}
}
