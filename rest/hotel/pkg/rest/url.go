package rest

import (
	"fmt"
	"strings"
)

// FilterOptions ...
type FilterOptions struct {
	Field    string
	Operator string
	Values   []string
}

// ToString ...
func (fo *FilterOptions) ToString() string {
	return fmt.Sprintf("%s%s", fo.Operator, strings.Join(fo.Values, ","))
}
