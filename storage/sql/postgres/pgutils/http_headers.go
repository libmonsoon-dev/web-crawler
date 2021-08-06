package pgutils

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/lib/pq/hstore"
)

func HeadersToHstore(h http.Header) (output hstore.Hstore) {
	output.Map = make(map[string]sql.NullString)

	for k, v := range h {
		var value sql.NullString
		if len(v) != 0 {
			value = sql.NullString{
				String: strings.Join(v, ";"),
				Valid:  true,
			}
		}
		output.Map[k] = value
	}

	return
}

func HstoreToHeaders(h hstore.Hstore) (output http.Header) {
	output = make(http.Header)

	for k, v := range h.Map {
		var value []string
		if v.Valid {
			value = strings.Split(v.String, ";")
		}

		output[k] = value
	}

	return
}
