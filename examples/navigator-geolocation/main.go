package main

import "gooey"
import "gooey/navigator"
import "strconv"
import "time"

func main() {

	element_latitude := gooey.Document.QuerySelector("#latitude")
	element_longitude := gooey.Document.QuerySelector("#longitude")
	element_altitude := gooey.Document.QuerySelector("#altitude")
	element_accuracy := gooey.Document.QuerySelector("#accuracy")
	element_error := gooey.Document.QuerySelector("#error")

	navigator.Geolocation.GetCurrentPosition(func(position navigator.GeolocationPosition) {

		element_latitude.SetInnerHTML(strconv.FormatFloat(position.Coords.Latitude, 'g', -1, 64))
		element_longitude.SetInnerHTML(strconv.FormatFloat(position.Coords.Longitude, 'g', -1, 64))
		element_altitude.SetInnerHTML(strconv.FormatFloat(position.Coords.Altitude, 'g', -1, 64))
		element_accuracy.SetInnerHTML(strconv.FormatFloat(position.Coords.Accuracy, 'g', -1, 64) + " meters")

	}, func(err navigator.GeolocationPositionError) {

		switch err {
		case navigator.GeolocationPositionErrorUnknown:
			element_error.SetInnerHTML("Unknown Error")
			break
		case navigator.GeolocationPositionErrorPermissionDenied:
			element_error.SetInnerHTML("Permission Denied")
			break
		case navigator.GeolocationPositionErrorPositionUnavailable:
			element_error.SetInnerHTML("Position Unavailable")
			break
		case navigator.GeolocationPositionErrorTimeout:
			element_error.SetInnerHTML("Timeout")
			break
		}

	})

	for true {

		// Do Nothing
		time.Sleep(1 * time.Second)

	}

}
