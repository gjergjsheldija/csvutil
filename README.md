# CSV Util

CSV Util is a wrapper for the csv package in go's standard library, designed to facilitate easy map access to csv files with header rows.  
This small library is based on the work of [CSV Map](https://github.com/peterdeka/go-csv-map) and of 
[CSV MAP](https://github.com/andrewcharlton/csvmap/tree/master)

## Installation

This package can be installed with the go get command

```
go get github.com/gjergjsheldija/csvutil
```

## Documentation

Where possible, the API has been designed to stick as closely to that of the original csv package as possible, with the exception that maps are returned instead of slices.

### Reading
``` go
func ExampleReader() {

	in := `name,alias,superpower
Logan,Wolverine,"Super healing and adamantium claws"
Charles Xavier,Professor X,Telepathy
`

	r := csvutil.NewReader(strings.NewReader(in))
    r.Columns, _ = reader.ReadHeader()
    
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Name:", record["name"])
		fmt.Println("Alias:", record["alias"])
		fmt.Println("Superpower:", record["superpower"])
		fmt.Println("")
	}

	// Output:
	// Name: Logan
	// Alias: Wolverine
	// Superpower: Super healing and adamantium claws
	//
	// Name: Charles Xavier
	// Alias: Professor X
	// Superpower: Telepathy
	//
}
```
### Writing

``` go
func ExampleWriter() {

	headers := []string{"Name", "Alias", "Superpower"}
	data := []map[string]string{
		{"Name": "Logan", "Alias": "Wolverine", "Superpower": "Super healing"},
		{"Name": "Charles Xavier", "Alias": "Professor X", "Superpower": "Telepathy"},
	}

	out := &bytes.Buffer{}
	w := csvutil.NewWriter(out, headers)

	err := w.WriteAll(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())

	// Output:
	// Name,Alias,Superpower
	// Logan,Wolverine,Super healing
	// Charles Xavier,Professor X,Telepathy
	//

}
```