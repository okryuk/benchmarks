package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func addSJSON(doc, patch []byte) []byte {
	fname := gjson.Get(string(patch), "0.value")
	final, _ := sjson.Set(string(doc), "data.firstname", fname.String())
	return []byte(final)
}

func replaceSJSON(doc, patch []byte) []byte {
	lname := gjson.Get(string(patch), "0.value")
	final, _ := sjson.Set(string(doc), "data.lastname", lname.String())
	return []byte(final)
}

func removeSJSON(doc, patch []byte) []byte {
	final, _ := sjson.Delete(string(doc), "data.lastname")
	return []byte(final)
}

func opJSONPATCH(doc, op []byte) []byte {
	patch, _ := jsonpatch.DecodePatch(op)
	final, _ := patch.Apply(doc)
	return final
}

// -----------------------------------------------------------------------------
// Test implementation for path op, extraction to make a proper benchmark for
// SJSON package vs JSONPATCH, that already have that logic and run it every time.

// uniSJSONwithRawMessage accepts multiple operations at once and selects the
// appropriate JSON patch logic. Stdlib json Unmarshalling is used with
// RawMessage for optimizations.
func uniSJSONwithRawMessage(doc, op []byte) []byte {
	var ops []json.RawMessage
	err := json.Unmarshal(op, &ops)
	if err != nil {
		log.Fatalln("error unmarshalling op json:", err)
	}

	for _, o := range ops {
		g := gjson.Get(string(o), "op")
		switch g.String() {
		case "add":
			rawPath := gjson.Get(string(o), "path")
			path := getPath(rawPath.String())
			value := gjson.Get(string(o), "value")
			result, _ := sjson.Set(string(doc), path, value.String())
			return []byte(result)
		case "replace":
			rawPath := gjson.Get(string(o), "path")
			path := getPath(rawPath.String())
			value := gjson.Get(string(o), "value")
			result, _ := sjson.Set(string(doc), path, value.String())
			return []byte(result)
		case "remove":
			rawPath := gjson.Get(string(o), "path")
			path := getPath(rawPath.String())
			result, _ := sjson.Delete(string(doc), path)
			return []byte(result)
		}
	}
	return nil
}

func getPath(path string) string {
	p := strings.Split(path[1:], "/")
	expPath := strings.Join(p, ".")
	return expPath
}

// uniSJSONOptim accepts multiple operations at once and selects the
// appropriate JSON patch logic. No Stdlib json Unmarshalling is used.
// GJSON and SJSON packages implement all the functionality.
func uniSJSONOptim(doc, op []byte) []byte {

	arrOp := gjson.Get(string(op), "#")

	for i := int64(0); i < arrOp.Int(); i++ {
		oneOp := gjson.Get(string(op), fmt.Sprint(i))
		g := gjson.Get(oneOp.String(), "op")

		switch g.String() {
		case "add":
			rawPath := gjson.Get(oneOp.String(), "path")
			path := getPath(rawPath.String())
			value := gjson.Get(oneOp.String(), "value")
			result, _ := sjson.Set(string(doc), path, value.String())
			return []byte(result)
		case "replace":
			rawPath := gjson.Get(oneOp.String(), "path")
			path := getPath(rawPath.String())
			value := gjson.Get(oneOp.String(), "value")
			result, _ := sjson.Set(string(doc), path, value.String())
			return []byte(result)
		case "remove":
			rawPath := gjson.Get(oneOp.String(), "path")
			path := getPath(rawPath.String())
			result, _ := sjson.Delete(string(doc), path)
			return []byte(result)
		}
	}

	return nil
}
