package InformaCast

import "fmt"

type QueryParameters struct {
	MaxCount int64
}

func (q *QueryParameters) Compile() (s string) {

	if q.MaxCount == 0 {
		q.MaxCount = 10000
	}

	// start query compile
	s += "?maxCount=" + fmt.Sprint(q.MaxCount)


	return s
}


// HttpParameters contains standard properties used by a http request
type HttpParameters struct {
	Server   string
	Username string
	Password string
	Method string
}


type ResponseManager struct {
	Total uint64
	Previous uint64
	Next string //url to loop
	Data []struct{}
}