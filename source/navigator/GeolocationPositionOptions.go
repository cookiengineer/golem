package navigator

type GeolocationPositionOptions struct {
	MaximumAge         uint `json:"maximumAge"`
	Timeout            uint `json:"timeout"`
	EnableHighAccuracy bool `json:"enableHighAccuracy"`
}

