package geolocation

var Geolocation geolocation

type Geolocation struct {
}

func init() {

	Gelocation = geolocation{
		// TODO
	}

}

func (geolocation *geolocation) GetCurrentPosition(onsuccess func(GeolocationPosition), onerror func(GeolocationPositionError), options map[string]bool) {

	// TODO: options

}

func (geolocation *geolocation) WatchPosition(onsuccess func(GeolocationPosition), onerror func(GeolocationPositionError), options map[string]bool) {

	// TODO

}
