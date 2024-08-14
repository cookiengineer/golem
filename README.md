
# Golem UI

Opinionated Web UI framework for Go that is made for
stateless HTML in both Web Pages and Web Views.


# Limitations

- `RemoveEventListener` cannot be used from WebASM, because `syscall/js` loses the reference to a `js.Func` when using `(js.Value).Call(...)`


# Opinions

- Static elements are never removed from the DOM
- Static elements can have DOM event listeners
- Elements with an `id` property are static elements

- Dynamic elements can be removed from the DOM
- Dynamic elements cannot have DOM event listeners

- Web Forms are static elements
- Web Forms with `enctype="application/json"` use a REST compatible API endpoint
- Web Forms with an `action` URL are automatically validateable
- Web Forms are serializable via `json.Marshal()` into a `struct`
- Web Forms are validateable via an `interface` with `Validate() bool, err`

- Web Clients are REST API clients
- Web Clients are validateable via an `interface` with `Validate() bool, err`
- Web Clients are routable via a `map[URL]struct` for each supported route


# Layouts

- [ ] fixed content, single sidebar
- [ ] scroll content, single sidebar
- [ ] fixed content, dual sidebar
- [ ] scroll content, dual sidebar


# Usage

TBD

# License

MIT
