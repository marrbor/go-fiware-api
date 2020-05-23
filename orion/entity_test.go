package orion_test

import (
	"fmt"
	"testing"

	"github.com/marrbor/go-fiware-datamodel/datamodel"
	"github.com/marrbor/go-orion-api/orionapi"
	"github.com/marrbor/goutil"
	"github.com/stretchr/testify/assert"
)

type (
	// https://fiware.github.io/specifications/ngsiv2/stable/cookbook/
	CookBookNormalized struct {
		Type         string       `json:"type"`
		ID           string       `json:"id"`
		Author       Author       `json:"author"`
		ItemReviewed ItemReviewed `json:"itemReviewed"`
		ReviewBody   ReviewBody   `json:"reviewBody"`
		ReviewRating ReviewRating `json:"reviewRating"`
	}

	Author struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	ItemReviewed struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}

	ReviewBody struct {
		Value string `json:"value"`
	}
	ReviewRating struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	}

	CookBookKeyValue struct {
		Type         string `json:"type"`
		ID           string `json:"id"`
		Author       string `json:"author"`
		ItemReviewed string `json:"itemReviewed"`
		ReviewBody   string `json:"reviewBody"`
		ReviewRating int    `json:"reviewRating"`
	}
)

func TestAccessor_CRUDEntity(t *testing.T) {
	err := orionapi.StartTestServer(t)
	assert.NoError(t, err)
	a := orionapi.NewAccessor(fmt.Sprintf("http://%s:%d", orionapi.Host, orionapi.Port))

	// normalized request
	n := CookBookNormalized{
		Type: "Review",
		ID:   "review-Elizalde-34",
		Author: Author{
			Type:  "Person",
			Value: "Client1234",
		},
		ItemReviewed: ItemReviewed{
			Type:  "Restaurant",
			Value: "",
		},
		ReviewBody: ReviewBody{
			Value: "Cheap and nice place to eat.",
		},
		ReviewRating: ReviewRating{
			Type:  "Rating",
			Value: 4,
		},
	}
	err = a.CreateEntity("", "", nil, &n)
	assert.NoError(t, err)

	// keyvalue request
	kv := CookBookKeyValue{
		Type:         "Review",
		ID:           "review-Elizalde-35",
		Author:       "client8921",
		ItemReviewed: "0115206c51f60b48b77e4c937835795c33bb953f",
		ReviewBody:   "Expensive but nice place to eat",
		ReviewRating: 8,
	}

	op := orionapi.NewKeyValuesQuery()
	err = a.CreateEntity("", "", op, &kv)
	assert.NoError(t, err)

	// GetEntityList with normalized
	var c = make([]CookBookNormalized, 0)
	err = a.GetEntityList("", "", nil, &c)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(c))
	j, err := goutil.JSONString(c)
	assert.NoError(t, err)
	t.Log(j)

	// GetEntityList with keyValue
	var ck = make([]CookBookKeyValue, 0)
	op = orionapi.NewKeyValuesQuery()
	err = a.GetEntityList("", "", op, &ck)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(ck))
	j2, err := goutil.JSONString(c)
	assert.NoError(t, err)
	t.Log(j2)

	// GetEntityList with Query
	ck = make([]CookBookKeyValue, 0)
	op = orionapi.NewKeyValuesQuery()
	err = op.SetIDQuery([]string{"review-Elizalde-34"})
	assert.NoError(t, err)
	err = a.GetEntityList("", "", op, &ck)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(ck))

	//
	ck = make([]CookBookKeyValue, 0)
	op = orionapi.NewKeyValuesQuery()
	err = op.SetTypeQuery([]string{"Review", "Person"})
	assert.NoError(t, err)
	err = a.GetEntityList("", "", op, &ck)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(ck))

	//
	ck = make([]CookBookKeyValue, 0)
	op = orionapi.NewKeyValuesQuery()
	err = op.SetIDPatternQuery(`review-Elizalde-[0-9]+$`)
	assert.NoError(t, err)
	err = a.GetEntityList("", "", op, &ck)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(ck))

	//
	ck = make([]CookBookKeyValue, 0)
	op = orionapi.NewKeyValuesQuery()
	err = op.SetTypePatternQuery(`^R.+$`)
	assert.NoError(t, err)
	err = a.GetEntityList("", "", op, &ck)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(ck))

	//
	ck = make([]CookBookKeyValue, 0)
	op = orionapi.NewKeyValuesQuery().SetQQuery([]string{"reviewRating>=7"})
	err = a.GetEntityList("", "", op, &ck)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(ck))
	/*
		// TODO Metadata Query test is not success (400 Bad Request)
		ck = make([]CookBookKeyValue, 0)
		op = orionapi.NewQuery().SetMQQuery([]string{"itemReviewed.type=Restaurant"})
		err = op.SetOptions([]orionapi.Option{orionapi.QueryOptions.KeyValues})
		assert.NoError(t, err)
		err = a.GetEntityList("", "", op, &ck)
		assert.NoError(t, err)
		assert.EqualValues(t, 1, len(ck))
	*/
	//
	ck = make([]CookBookKeyValue, 0)
	op = orionapi.NewKeyValuesQuery().SetLimit(1)
	err = a.GetEntityList("", "", op, &ck)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(ck))

	//
	ck = make([]CookBookKeyValue, 0)
	op = orionapi.NewKeyValuesQuery().SetOffset(1).SetAttrs([]string{"type", "id", "reviewRating"})
	err = a.GetEntityList("", "", op, &ck)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(ck))

	//
	ck = make([]CookBookKeyValue, 0)
	op = orionapi.NewKeyValuesQuery().SetOrderBy([]string{"reviewRating"}).SetAttrs([]string{"type", "id", "reviewRating"})
	err = a.GetEntityList("", "", op, &ck)
	assert.NoError(t, err)
	assert.EqualValues(t, 2, len(ck))

	// GetEntity
	var ce CookBookNormalized
	err = a.GetEntity("", "", "review-Elizalde-34", nil, &ce)
	assert.NoError(t, err)
	assert.EqualValues(t, "review-Elizalde-34", ce.ID)

	// GetEntityAttribute
	var au Author
	err = a.GetEntityAttribute("", "", "review-Elizalde-34", "author", nil, &au)
	assert.NoError(t, err)
	assert.EqualValues(t, "Client1234", au.Value)

	// update
	up := CookBookNormalized{
		Author:       Author{Value: "abcdefg"},
		ReviewBody:   ReviewBody{Value: "booooh!!"},
		ReviewRating: ReviewRating{Value: 1},
	}
	err = a.UpdateEntity("", "", "review-Elizalde-34", "Review", &up)
	assert.NoError(t, err)

	// GetEntityAttribute
	au = Author{}
	err = a.GetEntityAttribute("", "", "review-Elizalde-34", "author", nil, &au)
	assert.NoError(t, err)
	assert.EqualValues(t, "abcdefg", au.Value)

	// Delete
	err = a.DeleteEntity("", "", "review-Elizalde-34", "Review")
	assert.NoError(t, err)

	// Delete check
	c = make([]CookBookNormalized, 0)
	err = a.GetEntityList("", "", nil, &c)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, len(c))
	j, err = goutil.JSONString(c)
	assert.NoError(t, err)
	t.Log(j)

	// Delete
	err = a.DeleteEntity("", "", "review-Elizalde-35", "Review")
	assert.NoError(t, err)

	// GetEntityList with normalized
	c = make([]CookBookNormalized, 0)
	err = a.GetEntityList("", "", nil, &c)
	assert.NoError(t, err)
	assert.EqualValues(t, 0, len(c))

	// stop server
	err = orionapi.StopTestServer(t)
	assert.NoError(t, err)
	return

}

func TestErrorQuery(t *testing.T) {
	op := orionapi.NewQuery()
	err := op.SetTypeQuery([]string{"Review", "Person"})
	assert.NoError(t, err)
	err = op.SetTypePatternQuery(`^R.+$`)
	assert.Error(t, err, orionapi.IncompatibleQueryError.Error())
	err = op.SetIDQuery([]string{"Review", "Person"})
	assert.NoError(t, err)
	err = op.SetIDPatternQuery(`^i.+_.+$`)
	assert.Error(t, err, orionapi.IncompatibleQueryError.Error())
	l := datamodel.LatLng{
		Latitude:  -100,
		Longitude: 0,
	}
	err = op.SetCoordsQuery(&[]datamodel.LatLng{l})
	assert.Error(t, err, orionapi.InvalidLatLngError)

	op2 := orionapi.NewQuery()
	err = op2.SetTypePatternQuery(`^R.+$`)
	assert.NoError(t, err)
	err = op2.SetIDPatternQuery(`^i.+_.+$`)
	assert.NoError(t, err)

	err = op2.SetTypeQuery([]string{"Review", "Person"})
	assert.Error(t, err, orionapi.IncompatibleQueryError.Error())
	err = op2.SetIDQuery([]string{"Review", "Person"})
	assert.Error(t, err, orionapi.IncompatibleQueryError.Error())
}
