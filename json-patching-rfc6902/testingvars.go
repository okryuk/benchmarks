package main

var (
	DocStr = `{
		"properties": [{
			"name": "name",
			"type": "symbol",
			"updatedAt": "2023-03-12T20:10:00Z",
			"valueSymbol": "symbol"
		}, {
			"name": "license",
			"type": "qualification",
			"updatedAt": "2023-03-12T20:10:00Z",
			"valueQualification": {
				"number": "12345",
				"region": "QC"
			}
		}],
		"data": {
			"lastname": "Doe",
			"license": {
				"number": "12345",
				"region": "QC"
			}
		}
	}`

	DocB = []byte(DocStr)

	AddOp      = []byte(`[{"op": "add", "path": "/data/firstname", "value":"John"}]`)
	ReplaceOp  = []byte(`[{"op": "replace", "path": "/data/lastname", "value":"Smith"}]`)
	RemoveOp   = []byte(`[{"op": "remove", "path": "/data/lastname"}]`)
	MultipleOp = []byte(`[{"op": "add", "path": "/data/firstname", "value":"John"},{"op": "remove", "path": "/data/lastname"}]`)

	WantAddStr = `{
			"properties": [{
				"name": "name",
				"type": "symbol",
				"updatedAt": "2023-03-12T20:10:00Z",
				"valueSymbol": "symbol"
			}, {
				"name": "license",
				"type": "qualification",
				"updatedAt": "2023-03-12T20:10:00Z",
				"valueQualification": {
					"number": "12345",
					"region": "QC"
				}
			}],
			"data": {
				"firstname": "John",
				"lastname": "Doe",
				"license": {
					"number": "12345",
					"region": "QC"
				}
			}
		}`

	WantReplaceStr = `{
			"properties": [{
				"name": "name",
				"type": "symbol",
				"updatedAt": "2023-03-12T20:10:00Z",
				"valueSymbol": "symbol"
			}, {
				"name": "license",
				"type": "qualification",
				"updatedAt": "2023-03-12T20:10:00Z",
				"valueQualification": {
					"number": "12345",
					"region": "QC"
				}
			}],
			"data": {
				"lastname": "Smith",
				"license": {
					"number": "12345",
					"region": "QC"
				}
			}
		}`

	WantRemoveStr = `{
			"properties": [{
				"name": "name",
				"type": "symbol",
				"updatedAt": "2023-03-12T20:10:00Z",
				"valueSymbol": "symbol"
			}, {
				"name": "license",
				"type": "qualification",
				"updatedAt": "2023-03-12T20:10:00Z",
				"valueQualification": {
					"number": "12345",
					"region": "QC"
				}
			}],
			"data": {
				"license": {
					"number": "12345",
					"region": "QC"
				}
			}
		}`
)
