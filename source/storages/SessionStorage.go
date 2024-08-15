package storages

var SessionStorage Storage

func init() {

	SessionStorage = Storage{
		name: "sessionStorage",
	}

}
