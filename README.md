
# Gooey

<p align="center">
    <img width="256" height="256" src="https://raw.githubusercontent.com/cookiengineer/gooey/master/assets/gooey.jpg">
</p>

Gooey (GUI) is a Pure Go Web UI framework made for stateless HTML in Web Views.
It bridges the gaps between Go, WebASM, Browser APIs, unified App Layouts and reusable Themes.


# Opinions

**HTML Elements**

- Static elements can never be removed from the DOM
- Static elements can have DOM event listeners
- Static elements always have an `id` property
- Dynamic elements can be removed from the DOM
- Dynamic elements cannot have DOM event listeners

**App Layout**

- App Layout always consists of `header`, `main`, and `footer` elements
- App Views are represented by different `main > section[data-view=...]` elements
- App Views can contain `aside` elements representing the sidebar

**Web Clients**

- Web Clients are REST API clients using either [fetch](/source/fetch) or [xhr](/source/xhr)
- Web Clients are fulfill an `interface` with `Validate() bool, err`
- Web Clients are routable via a `map[URL]struct` for each supported route

**Web Forms**

Note: This is currently work-in-progress

- Web Forms are static elements
- Web Forms with `enctype="application/json"` use a REST compatible API endpoint
- Web Forms with an `action` URL are automatically validateable
- Web Forms are serializable via `json.Marshal()` into a `struct`
- Web Forms are fulfill an `interface` with `Validate() bool, err`


# Bindings

- [x] [Document](/source/Document.go)
- [x] [Screen](/source/Screen.go)
- [x] [ScreenOrientation](/source/ScreenOrientation.go)
- [x] [Window](/source/Window.go)

**animations**

- [x] [animations/CancelAnimationFrame](/source/animations/CancelAnimationFrame.go)
- [x] [animations/RequestAnimationFrame](/source/animations/RequestAnimationFrame.go)

**dom**

- [x] [dom/Element](/source/dom/Element.go)
- [x] [dom/Event](/source/dom/Event.go)
- [x] [dom/EventListener](/source/dom/EventListener.go)
- [x] [dom/EventPhase](/source/dom/EventPhase.go)
- [x] [dom/EventType](/source/dom/EventType.go)

**fetch**

Note: If you run into problems, use the [Synchronous XMLHttpRequest](/source/xhr/XMLHttpRequest_sync.go) APIs instead.

- [x] [fetch/Fetch](/source/fetch/Fetch.go) [2]
- [x] [fetch/Headers](/source/fetch/Headers.go)
- [x] [fetch/Request](/source/fetch/Request.go) (or `RequestInit` object)
- [x] [fetch/Response](/source/fetch/Response.go)

Fetch RequestInit Properties:

- [x] [fetch/Cache](/source/fetch/Cache.go)
- [x] [fetch/Credentials](/source/fetch/Credentials.go)
- [x] [fetch/Method](/source/fetch/Method.go)
- [x] [fetch/Mode](/source/fetch/Mode.go)
- [x] [fetch/Redirect](/source/fetch/Redirect.go)
- [x] `Referrer` has to be a `string` due to arbitrary URL values.
- [x] [fetch/ReferrerPolicy](/source/fetch/ReferrerPolicy.go)

**location**

- [x] [location/Location](/source/location/Location.go)

**navigator**

- [x] [navigator/Geolocation](/source/navigator/Geolocation.go)
- [x] [navigator/GeolocationPosition](/source/navigator/GeolocationPosition.go)
- [x] [navigator/GeolocationPositionError](/source/navigator/GeolocationPositionError.go)
- [x] [navigator/GeolocationPositionOptions](/source/navigator/GeolocationPositionOptions.go) [1]

**storages**

- [x] [storages/LocalStorage](/source/storages/LocalStorage.go)
- [x] [storages/SessionStorage](/source/storages/SessionStorage.go)

**timers**

- [x] [timers/ClearInterval](/source/timers/ClearInterval.go)
- [x] [timers/ClearTimeout](/source/timers/ClearTimeout.go)
- [x] [timers/SetInterval](/source/timers/SetInterval.go)
- [x] [timers/SetTimeout](/source/timers/SetTimeout.go)

**xhr**

- [x] [xhr/Method](/source/xhr/Method.go)
- [x] [xhr/XMLHttpRequest](/source/xhr/XMLHttpRequest.go) [2]
- [x] Synchronous [xhr/XMLHttpRequest](/source/xhr/XMLHttpRequest_sync.go)


--------

[1] This feature is implemented, but not supported across all Browsers. It is disabled to prevent WebASM runtime errors that are irrecoverable.

[2] This feature is implemented asynchronously and uses a go `chan`. It only works with `tinygo` as a compiler as of now. If your WebASM binary
    hangs when using this, use the synchronous XMLHttpRequest APIs instead.


# License

This project is licensed under the [MIT](./LICENSE_MIT.txt) license.

