
# Golem UI

Opinionated Web UI framework for Go that is made for
stateless HTML in both Web Pages and Web Views.


# Opinions

**HTML Elements**

- Static elements are never removed from the DOM
- Static elements can have DOM event listeners
- Elements with an `id` property are static elements

- Dynamic elements can be removed from the DOM
- Dynamic elements cannot have DOM event listeners

**Web Forms**

- Web Forms are static elements
- Web Forms with `enctype="application/json"` use a REST compatible API endpoint
- Web Forms with an `action` URL are automatically validateable
- Web Forms are serializable via `json.Marshal()` into a `struct`
- Web Forms are validateable via an `interface` with `Validate() bool, err`

**Web Clients**

- Web Clients are REST API clients
- Web Clients are validateable via an `interface` with `Validate() bool, err`
- Web Clients are routable via a `map[URL]struct` for each supported route


# Implementations

- [x] [Document](/source/Document.go)
- [ ] [Screen](/source/Screen.go)
- [x] [ScreenOrientation](/source/ScreenOrientation.go)
- [x] [Window](/source/Window.go)

**animations**

- [x] [animations/CancelAnimationFrame(uint)](/source/timers/CancelAnimationFrame.go)
- [x] [animations/RequestAnimationFrame(func(timestamp float64)) uint](/source/timers/RequestAnimationFrame.go)

**dom**

- [x] [dom/Element](/source/dom/Element.go)
- [x] [dom/Event](/source/dom/Event.go)
- [x] [dom/EventListener](/source/dom/EventListener.go)
- [x] [dom/EventPhase](/source/dom/EventPhase.go)
- [x] [dom/EventType](/source/dom/EventType.go)

**timers**

- [x] [timers/ClearInterval](/source/timers/ClearInterval.go)
- [x] [timers/SetInterval](/source/timers/SetInterval.go)
- [x] [timers/ClearTimeout](/source/timers/ClearTimeout.go)
- [x] [timers/SetTimeout](/source/timers/SetTimeout.go)


# Layouts

- [ ] fixed content, single sidebar
- [ ] scroll content, single sidebar
- [ ] fixed content, dual sidebar
- [ ] scroll content, dual sidebar


# Usage

TBD


# License

MIT
