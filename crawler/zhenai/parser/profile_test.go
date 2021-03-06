package parser

import (
	"io/ioutil"
	"learngo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := parseProfile(contents, "", "", "", "", "", "", "", "", "", "")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}

	profile := result.Items[0].Payload.(model.Profile)

	expected := model.Profile{
		Name:       "",
		Gender:     "",
		Age:        "",
		Height:     "",
		Weight:     "",
		Income:     "",
		Marriage:   "",
		Education:  "",
		Occupation: "",
		Hokou:      "",
		Xinzuo:     "",
		House:      "",
		Car:        "",
	}

	if profile != expected {
		t.Errorf("expected %v, but was %v", expected, profile)
	}
}
