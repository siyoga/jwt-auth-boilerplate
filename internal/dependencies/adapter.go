package dependencies

import (
	"github.com/siyoga/jwt-auth-boilerplate/internal/adapter"
	"github.com/siyoga/jwt-auth-boilerplate/internal/adapter/random"
	"github.com/siyoga/jwt-auth-boilerplate/internal/adapter/time"
)

func (d *dependencies) TimeAdapter() adapter.TimeAdapter {
	if d.timeAdapter == nil {
		d.timeAdapter = time.NewAdapter(
			d.cfg.Base.Locale,
		)
	}

	return d.timeAdapter
}

func (d *dependencies) RandomAdapter() adapter.RandomAdapter {
	if d.randomAdapter == nil {
		d.randomAdapter = random.NewAdapter()
	}
	return d.randomAdapter
}
