package shopify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gedex/inflector"
)

type Auth struct {
	Key    string
	Secret string
}

type Error struct {
	Errors string `json:"errors,omitempty"`
}

func (e *Error) Error() string {
	return e.Errors
}

type Product struct {
	Title string `json:"title,omitempty"`
}

type Store struct {
	auth *Auth
	Name string
}

func Authenticate() *Auth {
	key := os.Getenv("SHOPIFY_KEY")
	secret := os.Getenv("SHOPIFY_SECRET")
	return &Auth{key, secret}
}

func (a *Auth) NewStore(name string) *Store {
	return &Store{a, name}
}

func (s *Store) Create(in interface{}) error {
	name := strings.ToLower(inflector.Pluralize(structName(in)))

	return s.post(name, in)
}

func (s *Store) post(resource string, in interface{}) error {
	payload, err := json.Marshal(in)
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(payload)
	resp, err := http.Post(fmt.Sprintf("%s/admin/%s.json", s.url(), resource), "application/json", buff)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = getError(body); err != nil {
		return err
	}
	err = json.Unmarshal(body, in)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) url() string {
	return fmt.Sprintf("https://%s.myshopify.com", s.Name)
}

func getError(body []byte) error {
	e := &Error{}
	if err := json.Unmarshal(body, e); err != nil {
		return err
	}
	if e.Errors != "" {
		return e
	}
	return nil
}

func structName(in interface{}) string {
	if t := reflect.TypeOf(in); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
