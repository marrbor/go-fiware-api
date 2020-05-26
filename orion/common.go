// common definition about orion access
package orion

type (
	Http struct {
		Url string `json:"url"`
	}
)

func GenEntities(entities []string, id bool, pattern bool) *[]map[string]string {
	e := "id"
	if !id {
		e = "type"
	}
	if pattern {
		e += "Pattern"
	}

	es := make([]map[string]string, 0)
	for _, entity := range entities {
		es = append(es, map[string]string{e: entity})
	}
	return &es
}
