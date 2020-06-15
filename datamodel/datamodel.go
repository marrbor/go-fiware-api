/*
 * data model common
 *
 * https://fiware-datamodels.readthedocs.io/en/latest/howto/index.html
 * https://github.com/fiware/data-models/common-schema.json #GSMA-Commons, #DateYearLess
 */
package datamodel

import (
	"fmt"
	"regexp"
	"time"
)

const (
	Type = ""

	// from https://fiware.github.io/data-models/common-schema.json
	MinLengthEntityIdentifier = 1
	MaxLengthEntityIdentifier = 256
	NgsiLdOffset              = len("urn:ngsi-ld::")
	AllowedIdentifierLength   = MaxLengthEntityIdentifier - NgsiLdOffset
)

////// Type: Identifier

// Identifier is an alias name for string reflected "EntityIdentifierType" of common-schema.json
type Identifier string

// IdentifierRegexp is an expression about EntityIdentifierType#anyOf#pattern of common-schema.json.
var IdentifierRegexp = regexp.MustCompile("^[\\w\\-\\.\\{\\}\\$\\+\\*\\[\\]` | ~^@!, :\\\\]+$")

// ValidateIdentifier validates whether given strings are suited for Identifier or not.
func ValidateIdentifier(id string) bool {
	return IdentifierRegexp.Match([]byte(id))
}

// GenNgsiLdID returns whether ID string is suited for NGSI-LD identifier.
func GenNgsiLdID(baseName, typeName string) Identifier {
	return Identifier(fmt.Sprintf("urn:ngsi-ld:%s:%s", typeName, baseName))
}

////// Type: Core
var (
	TooShortBaseNameError = fmt.Errorf("too short base name")
	TooShortTypeNameError = fmt.Errorf("too short type name")
	TooLongIDLengthError  = fmt.Errorf("too long id length")
	InvalidTypeNameError  = fmt.Errorf("invalid type name")
	TypeNameRegexp        = regexp.MustCompile(`^[A-Z]`) // Have to begin upper case.
)

// Core is a necessary item for data model entity. It comes from "GSMA-Commons".
type Core struct {
	// ID is a unique identified of the entity modelled. Have to follow the NGSI-LD rules.
	ID Identifier `json:"id"`
	// Type is the entity type, i.e. the type of Data Model, e.g. Alert.
	Type string `json:"type"`
	// DateModified : Last update timestamp of the entity.
	DateModified time.Time `json:"dateModified"`
	// DateCreated : The entity's creation timestamp.
	DateCreated time.Time `json:"dateCreated"`
}

// Stamp set DateModified, and if DateCreated has not set, set it.
func (c *Core) Stamp() {
	c.DateModified = time.Now()
	if c.DateCreated.IsZero() {
		c.DateCreated = c.DateModified
	}
}

// NewCore returns new Core instance. name means base of ID.
func NewCore(baseName, typeName string) (*Core, error) {
	if len(baseName) < MinLengthEntityIdentifier {
		return nil, TooShortBaseNameError
	}
	if len(typeName) < MinLengthEntityIdentifier {
		return nil, TooShortTypeNameError
	}
	if !ValidateTypeName(typeName) {
		return nil, InvalidTypeNameError
	}
	if AllowedIdentifierLength < (len(baseName) + len(typeName)) {
		return nil, TooLongIDLengthError
	}

	return &Core{
		ID:           GenNgsiLdID(baseName, typeName),
		Type:         typeName,
		DateModified: time.Time{},
		DateCreated:  time.Time{},
	}, nil
}

// ValidateTypeName validates given type name string is suit for type name.
func ValidateTypeName(tn string) bool {
	return TypeNameRegexp.Match([]byte(tn))
}

////// Type: CoreOpt

// CoreOpt is a optional item for date model entity.
// Rest of "GSMA-Commons" (without Core and "seeAlso" element).
type CoreOpt struct {
	Owner         []Identifier `json:"owner"`         // An array of URIs or pointers to NGSI entities representing the owner(s) of the entity.
	Source        string       `json:"source"`        // A pointer (eventually an URI) to the service providing the data.
	Name          string       `json:"name"`          // A mnemonic name given to the entity as per schema.org defined within the core context as https://uri.etsi.org/ngsi-ld/name
	AlternateName string       `json:"alternateName"` // An alternative mnemonic name given to the entity as per schema.org
	Description   string       `json:"description"`   // A textual description of the entity as per schema.org defined within the core context as https://uri.etsi.org/ngsi-ld/description
	DataProvider  string       `json:"dataProvider"`  // A name identifying the entity providing the data.
}

////// Type: DateYearLess

// DateYearLess is an expression registered as "DateYearLess" in common-schema.json.
type DateYearLess string

var DateYearLessRegexp = regexp.MustCompile("^--((0[13578]|1[02])-31|(0[1,3-9]|1[0-2])-30|(0\\d|1[0-2])-([0-2]\\d))$")

// ValidateDateYearLess returns whether ID string is suited for DateYearLess.
func ValidateDateYearLess(d string) bool {
	return DateYearLessRegexp.Match([]byte(d))
}
