package middleware

import (
	"myapp/data"

	"github.com/IrakliGiorgadze/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
