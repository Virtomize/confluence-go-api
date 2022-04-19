package goconfluence

import (
	"fmt"
	"github.com/perolo/gojson"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type ConfluenceTestType struct {
	APIEndpoint string
	File        string
	Method      string
	Type        string
	TypeFile    string
}

const (
	TestSpaceGetSpacesMocFileS = iota
	TestAddCategoryResponseType
)

var ConfluenceTest = []ConfluenceTestType{
	{APIEndpoint: "/rest/api/space", File: "mocks/spaces.json", Method: "GET", Type: "AllSpaces", TypeFile: "space-dtos.go"},
	{APIEndpoint: "/rest/extender/1.0/category/addSpaceCategory/space/ds/category/test", File: "mocks/extender-add.json", Method: "PUT", Type: "AddCategoryResponseType", TypeFile: "extender-dtos.go"},
}

func UpdateTests() {
	confClient, err := NewAPI("http://localhost:1990/confluence", "admin", "admin")
	confClient.Debug = true
	if err != nil {
		log.Fatal(err)
	}
	//Remove all old files
	for _, ctest := range ConfluenceTest {
		e := os.Remove(ctest.File)
		if e != nil {
			fmt.Printf("Expected? %s\n", e.Error())
		}
		e = os.Remove(ctest.TypeFile)
		if e != nil {
			fmt.Printf("Expected? %s\n", e.Error())
		}
	}
	for _, ctest := range ConfluenceTest {
		resp, err2 := confClient.SendGenericRequest(ctest.APIEndpoint, ctest.Method)
		if err2 != nil {
			log.Fatal(err2)
		}
		err3 := ioutil.WriteFile(ctest.File, resp, 0644)
		if err3 != nil {
			log.Fatal(err3.Error())
		}
		i := strings.NewReader(string(resp))
		res, err2 := gojson.Generate(i, gojson.ParseJson, ctest.Type, "goconfluence", []string{"json"}, false, true)
		if err2 != nil {
			log.Fatal(err2.Error())
		}
		err3 = ioutil.WriteFile(ctest.TypeFile, res, 0644)
		if err3 != nil {
			log.Fatal(err3.Error())
		}
	}
}
