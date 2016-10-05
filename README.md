# Gofencer
## Go library for accessing the Fencer.io API

Download with ```go get github.com/fencerio/gofencer```

Consult the package and/or the API docs to establish which arguments are required. https://fencer.io/developers

## Tests

Without an API key there is only 18% test coverage. Tests can provide 80% coverage with an API key 
obtained from the Fencer.io service. Simply set it as an environment variable prior to calling ```go test```

```
export API_KEY='f7c04eaf-5510-5878-xxxxx-acad3c60f'
go test -cover
PASS
coverage: 80.7% of statements
ok      github.com/fencerio/gofencer    0.491s
```

## Use
An example of typical setup and use is included below.

```
package main 

import (
	"log"
	"github.com/fencerio/gofencer"
)

func main() {
	apiKey := "f7c04eaf-5510-5878-xxxxx-acad3c60f"
	accessKey := "4c89693a-02c4-xxxx-xxxx-cdf801237e3d"
	lat := 53.6750352651078
	lng := -2.4879334942895883

	f := new(gofencer.API)
	f.SetAPIKey(apiKey)
	f.SetAccessKey(accessKey)
	f.SetLatLng(lat, lng)

	// Example call to the API
	// Determine if passed lat/lng is inside geofence with
	response, err := f.PositionInside()
	if err != nil {
		log.Fatal(err)
	}

	// Just print it
	log.Println(response.Data.Inside) // true or false
}
```

## License

GPL license
