package datamodel_test

import (
	"testing"

	"github.com/marrbor/go-fiware-api/datamodel"
	"github.com/stretchr/testify/assert"
)

func TestGenLocation(t *testing.T) {
	l, err := datamodel.GenLocation(datamodel.TypePoint)
	assert.NoError(t, err)
	p, ok := l.Location.(datamodel.Point)
	assert.True(t, ok)
	assert.EqualValues(t, datamodel.TypePoint, p.Type)
	assert.EqualValues(t, 2, len(p.Coordinates))

	l, err = datamodel.GenLocation(datamodel.TypeLineString)
	assert.NoError(t, err)
	ls, ok := l.Location.(datamodel.LineString)
	assert.True(t, ok)
	assert.EqualValues(t, datamodel.TypeLineString, ls.Type)
	assert.EqualValues(t, 2, len(ls.Coordinates))
	assert.EqualValues(t, 2, len(ls.Coordinates[0]))
	assert.EqualValues(t, 2, len(ls.Coordinates[1]))

	l, err = datamodel.GenLocation(datamodel.TypePolygon)
	assert.NoError(t, err)
	pg, ok := l.Location.(datamodel.Polygon)
	assert.True(t, ok)
	assert.EqualValues(t, datamodel.TypePolygon, pg.Type)
	assert.EqualValues(t, 1, len(pg.Coordinates))
	assert.EqualValues(t, 4, len(pg.Coordinates[0]))
	for i := 0; i < 4; i++ {
		assert.EqualValues(t, 2, len(pg.Coordinates[0][i]))
	}

	l, err = datamodel.GenLocation(datamodel.TypeMultiPoint)
	assert.NoError(t, err)
	mp, ok := l.Location.(datamodel.MultiPoint)
	assert.True(t, ok)
	assert.EqualValues(t, datamodel.TypeMultiPoint, mp.Type)
	assert.EqualValues(t, 1, len(mp.Coordinates))
	assert.EqualValues(t, 2, len(mp.Coordinates[0]))

	l, err = datamodel.GenLocation(datamodel.TypeMultiLineString)
	assert.NoError(t, err)
	mls, ok := l.Location.(datamodel.MultiLineString)
	assert.True(t, ok)
	assert.EqualValues(t, datamodel.TypeMultiLineString, mls.Type)
	assert.EqualValues(t, 1, len(mls.Coordinates))
	assert.EqualValues(t, 2, len(mls.Coordinates[0]))
	for i := 0; i < 2; i++ {
		assert.EqualValues(t, 2, len(mls.Coordinates[0][i]))
	}

	l, err = datamodel.GenLocation(datamodel.TypeMultiPolygon)
	assert.NoError(t, err)
	mpg, ok := l.Location.(datamodel.MultiPolygon)
	assert.True(t, ok)
	assert.EqualValues(t, datamodel.TypeMultiPolygon, mpg.Type)
	assert.EqualValues(t, 1, len(mpg.Coordinates))
	assert.EqualValues(t, 1, len(mpg.Coordinates[0]))
	for i := 0; i < 4; i++ {
		assert.EqualValues(t, 2, len(mpg.Coordinates[0][0][i]))
	}

	// invalid
	l, err = datamodel.GenLocation("loc")
	assert.Error(t, err)
}

func TestLatLng_String(t *testing.T) {
	l := datamodel.LatLng{
		Latitude:  35.1,
		Longitude: 135.98,
	}
	assert.EqualValues(t, "35.1,135.98", l.String())

	l = datamodel.LatLng{
		Latitude:  35.123,
		Longitude: 135.9876,
	}
	assert.EqualValues(t, "35.123,135.9876", l.String())

	l = datamodel.LatLng{
		Latitude:  35.123456,
		Longitude: 135.987654,
	}
	assert.EqualValues(t, "35.123456,135.987654", l.String())

	l = datamodel.LatLng{
		Latitude:  35.123456789,
		Longitude: 135.987654321,
	}
	assert.EqualValues(t, "35.123456789,135.987654321", l.String())
}

func TestXY_String(t *testing.T) {
	l := datamodel.XY{
		X: 0.1,
		Y: -9.87,
	}
	assert.EqualValues(t, "0.1,-9.87", l.String())

	l = datamodel.XY{
		X: 0.123,
		Y: -9.8765,
	}
	assert.EqualValues(t, "0.123,-9.8765", l.String())

	l = datamodel.XY{
		X: 0.12345,
		Y: -9.876543,
	}
	assert.EqualValues(t, "0.12345,-9.876543", l.String())

	l = datamodel.XY{
		X: 0.1234567,
		Y: -9.87654321,
	}
	assert.EqualValues(t, "0.1234567,-9.87654321", l.String())

	l = datamodel.XY{
		X: 0.123456789,
		Y: -9.8765432109,
	}
	assert.EqualValues(t, "0.123456789,-9.8765432109", l.String())

}
