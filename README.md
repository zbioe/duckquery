# DUCKQUERY

## USAGE

### Example
```
package main

import (
	"fmt"
	"encoding/json"
	"github.com/iuryfukuda/duckquery"
)

func main() {
	o := duckquery.Query("Eduardo Campos")
	b, _ := json.Marshal(o)
	fmt.Printf("Output:%s", b)
}

```

## OUTPUT 


### EXAMPLE

```
{
	"abstract": "Eduardo Henrique Accioly Campos was a Brazilian congressman and governor. Born and raised in Recife, in the Northeast Brazil, he graduated in Economics from the Recife's Federal University of Pernambuco. Campos' maternal grandfather, the governor of the Brazilian state, Pernambuco, made him his Financial Secretary. Campos became a federal congressman in Brazil and got Pernambuco federal money for a shipyard, railways and an oil refinery. Later, as Brazil's Minister for Science and Technology, he supported stem-cell research. He served two terms as governor of his home state, Pernambuco. He helped hospitals, secondary schools, wind power, farms, poor people and anti-crime data-mining. In his 2014 campaign for president of Brazil he criticized the incumbent and her Workers' Party and positioned himself as the business-friendly leader of the Brazilian Socialist Party. For outdoor rallies and local radio interviews, he criss-crossed the country by rented jet.",
	"labels": [{
			"label": "Born",
			"value": "Aug 10, 1965, Recife, Brazil"
		},
		{
			"label": "Died",
			"value": "Aug 13, 2014, Santos, Brazil"
		},
		{
			"label": "Political party",
			"value": "Socialist Party"
		},
		{
			"label": "Alma mater",
			"value": "Federal University of Pernambuco"
		},
		{
			"label": "Religion",
			"value": "Roman Catholicism"
		},
		{
			"label": "Preceded by",
			"value": "Mendonça Filho"
		},
		{
			"label": "Succeeded by",
			"value": "João Lyra Neto"
		}
	],
	"topics": [
		"Federal University of Pernambuco alumni",
		"Victims of aviation accidents or incidents in Brazil",
		"Members of the Legislative Assembly of Pernambuco",
		"Governors of Pernambuco",
		"Brazilian Socialist Party politicians",
		"Recipients of the Great Cross of the National Order of Scientific Merit (Brazil)",
		"Government ministers of Brazil",
		"Brazilian Roman Catholics"
	]
}
```

### SCHEMA

```
{
	"abstract": {
		"type": "string",
		"required": false
	},
	"labels": {
		"type": "array",
		"required": false,
		"format": {
			"type": "object",
			"required": true,
			"format": {
				"label": {
					"type": "string",
					"required": true
				},
				"value": {
					"type": "string",
					"required": true
				}
			}
		}
	},
	"topics": {
		"type": "array",
		"required": false,
		"format": {
			"type": "string"
		}
	}
}
```
