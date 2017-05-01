package micro_web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	v "github.com/laux-development/micro_view"
)

type testProfileModel struct{}

func (pm testProfileModel) GithubName() string {
	return "bleeppurple"
}

func (pm testProfileModel) FirstName() string {
	return "james"
}

func (pm testProfileModel) LastName() string {
	return "laux"
}

func (pm testProfileModel) Address() string {
	return "a nice road, in a nice place, goodpostcode"
}

func (pm testProfileModel) Company() string {
	return "Laux Development"
}

func (pm testProfileModel) JobRole() string {
	return "Developer"
}

func (pm testProfileModel) DOB() string {
	return "20/01/1990"
}

func (pm testProfileModel) Gender() string {
	return "male"
}

type testIncompleteProfileModel struct{}

func (pm testIncompleteProfileModel) GithubName() string {
	return "bleeppurple"
}

func (pm testIncompleteProfileModel) FirstName() string {
	return "james"
}

func (pm testIncompleteProfileModel) LastName() string {
	return "laux"
}

func (pm testIncompleteProfileModel) Address() string {
	return "a nice road, in a nice place, goodpostcode"
}

func (pm testIncompleteProfileModel) Company() string {
	return "Laux Development"
}

func (pm testIncompleteProfileModel) JobRole() string {
	return "Developer"
}

func (pm testIncompleteProfileModel) DOB() string {
	return "20/01/1990"
}

type testProfileHTTPHandler struct{}

func (h *testProfileHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// create profile model
	m := testProfileModel{}

	// create view model
	v := v.NewView()

	// create web
	web := NewWeb(v, m)
	web.Profile(w, r)
}

type testIncompleteProfileHTTPHandler struct{}

func (h *testIncompleteProfileHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// create profile model
	m := testIncompleteProfileModel{}

	// create view model
	v := v.NewView()

	// create web
	web := NewWeb(v, m)
	web.Profile(w, r)
}

func TestProfile(t *testing.T) {

	//if err := os.Chdir("../micro_view"); err != nil {
	//	panic(err)
	//}

	// Create server with register token handler
	h := &testProfileHTTPHandler{}
	server := httptest.NewServer(h)
	defer server.Close()

	// Make a test request
	resp, err := http.Get(server.URL)
	if err != nil {
		fmt.Printf("error: %v \n", err)
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("expected to parse response but got : %s", err)
	}

	if !strings.Contains(string(body), "goodpostcode") {
		t.Errorf("expected Profile to be rendered as HTML but got : %s", string(body))
	}
}

func TestIncompleteProfile(t *testing.T) {

	//if err := os.Chdir("../micro_view"); err != nil {
	//	panic(err)
	//}

	// Create server with register token handler
	h := &testIncompleteProfileHTTPHandler{}
	server := httptest.NewServer(h)
	defer server.Close()

	// Make a test request
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("response error: %v \n", err)
	}

	if resp.StatusCode == 200 {
		t.Errorf("Received non-200 response: %d\n", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("expected to parse response but got : %s", err)
	}

	if strings.Contains(string(body), "goodpostcode") {
		t.Errorf("expected Profile to be rendered as HTML but got : %s", string(body))
	}

}
