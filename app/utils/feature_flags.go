package utils

type Channel struct {
	ChannelId int
	PartnerId int
}

type FeatureFlag func(channel Channel) bool

// todo: put this into a DB?
type FeatureFlags struct {
	RateOriginCodeField FeatureFlag
	UseKosmosV2Route    FeatureFlag
}

func NewFeatureFlagsManager() *FeatureFlags {
	//todo: make the map more dynamic
	validateFeatureFlagConditions := func(allowedChannels []Channel) FeatureFlag {
		return func(channel Channel) bool {
			for _, allowedChannel := range allowedChannels {
				if channel == allowedChannel {
					return true
				}
			}
			return false
		}
	}

	return &FeatureFlags{RateOriginCodeField: validateFeatureFlagConditions([]Channel{
		{27, 2}, {18, 1940},
	}),
		UseKosmosV2Route: validateFeatureFlagConditions([]Channel{
			{27, 2},
			{18, 1940},
			{5, 1940},
		})}
}
