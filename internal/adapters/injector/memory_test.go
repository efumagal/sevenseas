package injector

import (
	"strings"
	"testing"

	"github.com/efumagal/sevenseas/internal/adapters/repository"
	"github.com/efumagal/sevenseas/internal/core/services"
	"github.com/stretchr/testify/assert"
)

var data = `{
	"AEAJM": {
		"name": "Ajman",
		"city": "Ajman",
		"country": "United Arab Emirates",
		"alias": [],
		"regions": [],
		"coordinates": [
		  55.5136433,
		  25.4052165
		],
		"province": "Ajman",
		"timezone": "Asia/Dubai",
		"unlocs": [
		  "AEAJM"
		],
		"code": "52000"
	  },
	  "AEAUH": {
		"name": "Abu Dhabi",
		"coordinates": [
		  54.37,
		  24.47
		],
		"city": "Abu Dhabi",
		"province": "Abu Z¸aby [Abu Dhabi]",
		"country": "United Arab Emirates",
		"alias": [],
		"regions": [],
		"timezone": "Asia/Dubai",
		"unlocs": [
		  "AEAUH"
		],
		"code": "52001"
	  }
}`

func TestAddFromStream(t *testing.T) {
	store := repository.NewPortMemoryRepository()
	svc := services.NewPortService(store)
	pfs := NewPortStreamService(svc)
	inserted, err := pfs.InjectStream(strings.NewReader(data))
	assert.Nil(t, err)
	assert.Equal(t, 2, inserted)
}
