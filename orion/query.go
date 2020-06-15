package orion

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/marrbor/go-fiware-api/datamodel"
	"github.com/marrbor/goutil"
)

var (
	IncompatibleQueryError = fmt.Errorf("incompatible query")
	InvalidLatLngError     = fmt.Errorf("invalid latitude or longitude")
)

// Query is a structure for handling query for fiware-orion.
type Query struct {
	queries map[string]string
}

// SetToURL sets query attribute for given url instance.
func (q *Query) SetToURL(u *url.URL) {
	uq := u.Query() // load current query set.
	for k, v := range q.queries {
		uq.Set(k, v) // add queries into current query set.
	}
	u.RawQuery = uq.Encode() // update query.
}

// SetToRequest sets query attribute for given http request.
func (q *Query) SetToRequest(r *http.Request) {
	u := r.URL
	q.SetToURL(u)
}

// IsExists returns whether the given key has been already existing or not.
func (q *Query) IsExists(k string) bool {
	_, ok := q.queries[k]
	return ok
}

// SetQuery sets given query strings (key & value) into this instance.
// If the same key has been already exists, replace it.
func (q *Query) SetQuery(k, v string) *Query {
	q.RemoveQuery(k)
	q.queries[k] = v
	return q
}

// RemoveQuery removes given key from this instance.
func (q *Query) RemoveQuery(k string) *Query {
	if q.IsExists(k) {
		delete(q.queries, k)
	}
	return q
}

// SetIDQuery convert given ID list into query parameter and sets them into this instance.
// Use for GetEntityList
// id A comma-separated list of elements. Retrieve entities whose ID matches one of the elements in the list.
// Incompatible with idPattern.
// Example: Boe_Idearium.
func (q *Query) SetIDQuery(list []string) error {
	if q.IsExists("idPattern") {
		return IncompatibleQueryError
	}
	q.SetQuery("id", strings.Join(list, ","))
	return nil
}

// SetTypeQuery convert given type list into query parameter and sets them into this instance.
// type comma-separated list of elements. Retrieve entities whose type matches one of the elements in the list.
// Incompatible with typePattern.
// Example: Room.
func (q *Query) SetTypeQuery(list []string) error {
	if q.IsExists("typePattern") {
		return IncompatibleQueryError
	}
	q.SetQuery("type", strings.Join(list, ","))
	return nil
}

// SetIdPatternQuery sets given string as idPattern query strings into this instance.
// idPattern A correctly formated regular expression. Retrieve entities whose ID matches the regular expression.
// Incompatible with id.
// Example: Bode_.*.
func (q *Query) SetIDPatternQuery(re string) error {
	if q.IsExists("id") {
		return IncompatibleQueryError
	}
	if _, err := regexp.Compile(re); err != nil {
		return err
	}
	q.SetQuery("idPattern", re)
	return nil
}

// SetTypePatternQuery sets given string as typePattern query strings into this instance.
// typePattern A correctly formated regular expression. Retrieve entities whose type matches the regular expression.
// Incompatible with type.
// Example: Room_.*.
func (q *Query) SetTypePatternQuery(re string) error {
	if q.IsExists("type") {
		return IncompatibleQueryError
	}
	if _, err := regexp.Compile(re); err != nil {
		return err
	}
	q.SetQuery("typePattern", re)
	return nil
}

// SetQQuery convert given type list into query parameter and sets them into this instance.
// q A query expression, composed of a list of statements separated by ;, i.e., q=statement1;statement2;statement3.
// See [Simple Query Language specification](https://jsapi.apiary.io/previews/null/reference/entities/list-entities/list-entities#simple_query_language).
// Example: temperature>40.
func (q *Query) SetQQuery(list []string) *Query {
	return q.SetQuery("q", strings.Join(list, ";"))
}

// SetMQQuery convert given type list into query parameter and sets them into this instance.
// mq A query expression for attribute metadata, composed of a list of statements separated by ;, i.e., mq=statement1;statement2;statement3.
// See [Simple Query Language specification](https://jsapi.apiary.io/previews/null/reference/entities/list-entities/list-entities#simple_query_language).
// Example: temperature.accuracy<0.9.
func (q *Query) SetMQQuery(list []string) *Query {
	return q.SetQuery("mq", strings.Join(list, ";"))
}

// SetGeoRelQuery sets given string as georel query strings into this instance.
// georel Spatial relationship between matching entities and a reference shape.
// See [Geographical Queries](https://jsapi.apiary.io/previews/null/reference/entities/list-entities/list-entities#geographical_queries).
// Example: near.
// TODO check whether given string is suit to parameter or not.
func (q *Query) SetGeorelQuery(gq string) error {
	q.SetQuery("georel", gq)
	return nil
}

// SetGeometryQuery sets given string as geometry query strings into this instance.
// geometry Geografical area to which the query is restricted.
// See [Geographical Queries](https://jsapi.apiary.io/previews/null/reference/entities/list-entities/list-entities#geographical_queries).
// Example: point.
// TODO check whether given string is suit to parameter or not.
func (q *Query) SetGeometryQuery(gm string) error {
	q.SetQuery("geometry", gm)
	return nil
}

// SetCoordsQuery sets
// coords List of latitude-longitude pairs of coordinates separated by ';'. See Geographical Queries. Example: 41.390205,2.154007;48.8566,2.3522.
func (q *Query) SetCoordsQuery(l *[]datamodel.LatLng) error {
	var qs []string
	for _, ll := range *l {
		if !goutil.IsValidLatitude(ll.Latitude) {
			return InvalidLatLngError
		}
		if !goutil.IsValidLongitude(ll.Longitude) {
			return InvalidLatLngError
		}
		qs = append(qs, ll.String())
	}
	q.SetQuery("coords", strings.Join(qs, ";"))
	return nil
}

// SetLimit sets limits the number of entities to be retrieved.
// Example: 20.
func (q *Query) SetLimit(limit int) *Query {
	return q.SetQuery("limit", fmt.Sprintf("%d", limit))
}

// SetOffset sets the offset from where entities are retrieved
// Example: 20.
func (q *Query) SetOffset(offset int) *Query {
	return q.SetQuery("offset", fmt.Sprintf("%d", offset))
}

// SetAttrs sets attrs Comma-separated list of attribute names whose data are to be included in the response.
// The attributes are retrieved in the order specified by this parameter.
// If this parameter is not included, the attributes are retrieved in arbitrary order.
// See "Filtering out attributes and metadata" section for more detail.
// Example: seatNumber.
func (q *Query) SetAttrs(attrs []string) *Query {
	return q.SetQuery("attrs", strings.Join(attrs, ","))
}

// SetMetadata sets metadata A list of metadata names to include in the response.
// See "Filtering out attributes and metadata" section for more detail
// Example: accuracy.
func (q *Query) SetMetadata(md []string) *Query {
	return q.SetQuery("metadata", strings.Join(md, ","))
}

// orderBy	Criteria for ordering results.
// See "Ordering Results" section for details.
// Example: temperature,!speed.
func (q *Query) SetOrderBy(ob []string) *Query {
	return q.SetQuery("orderBy", strings.Join(ob, ","))
}

// SetOptions sets query options.
// Possible values:  count , keyValues , values , unique .
func (q *Query) SetOptions(list []Option) *Query {
	var opts []string
	for _, o := range list {
		opts = append(opts, o.value)
	}
	q.SetQuery("options", strings.Join(opts, ","))
	return q
}

// NewQuery returns new (empty) Query instance.
func NewQuery() *Query {
	return &Query{queries: make(map[string]string)}
}

// NewKeyValueQuery returns new query set keyvalue mode query.
func NewKeyValuesQuery() *Query {
	return NewQuery().SetOptions([]Option{QueryOptions.KeyValues})
}

///// Query Option
type (
	Option struct{ value string }
)

var (
	// QueryOptions holds possible query option value.
	QueryOptions = struct {
		Count     Option
		KeyValues Option
		Values    Option
		Unique    Option
		Append    Option
	}{
		Count:     Option{"count"},
		KeyValues: Option{"keyValues"},
		Values:    Option{"values"},
		Unique:    Option{"unique"},
		Append:    Option{"append"},
	}
)
