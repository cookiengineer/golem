package fetch

import "context"
import "io"

type Request struct {
	Method         Method            `json:"method"`
	Headers        map[string]string `json:"headers"`
	Body           io.Reader         `json:"body"`
	Mode           Mode              `json:"mode"`
	Credentials    Credentials       `json:"credentials"`
	Cache          Cache             `json:"cache"`
	Redirect       Redirect          `json:"redirect"`
	Referrer       string            `json:"referrer"`
	ReferrerPolicy string            `json:"referrerPolicy"`
	Integrity      string            `json:"integrity"`
	KeepAlive      bool              `json:"keepalive"`
	Signal         context.Context
}


func (request *Request) MapToJS() map[string]interface{} {

	mapped := make(map[string]interface{})

	if tmp := request.Method.String(); tmp != "" {
		mapped["method"] = request.Method.String()
	}

	if len(request.Headers) > 0 {

		mapped_headers := make(map[string]interface{})

		for key, val := range request.Headers {
			mapped_headers[key] = val
		}

	}

	if tmp := request.Mode.String(); tmp != "" {
		mapped["mode"] = request.Mode.String()
	}

	if tmp := request.Credentials.String(); tmp != "" {
		mapped["credentials"] = request.Credentials.String()
	}

	if tmp := request.Cache.String(); tmp != "default" && tmp != "" {
		mapped["cache"] = request.Cache.String()
	}

	if tmp := request.Redirect.String(); tmp != "follow" && tmp != "" {
		mapped["redirect"] = request.Redirect.String()
	}

	if request.Referrer != "" {
		mapped["referrer"] = request.Referrer
	}

	if request.ReferrerPolicy != "" {
		mapped["referrerPolicy"] = request.ReferrerPolicy
	}

	if request.Integrity != "" {
		mapped["integrity"] = request.Integrity
	}

	if request.KeepAlive != false {
		mapped["keepalive"] = request.KeepAlive
	}

	return mapped

}

