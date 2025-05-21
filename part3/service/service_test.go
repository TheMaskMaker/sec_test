package service

import (
	"io"
	"net/http/httptest"
	"part3/database"
	"testing"
)

type MockDb struct{}

func (m *MockDb) Get(key string) (string, bool) {
	return "value", true
}
func (m *MockDb) Update(key string, value string) error {
	return nil
}
func (m *MockDb) Delete(key string) error {
	return nil
}

var _ database.DataKind = &MockDb{}

func TestHelloName(t *testing.T) {

}

func TestGetItem(t *testing.T) {
	t.Error("foo")
	req := httptest.NewRequest("GET", "/api/123", nil)
	w := httptest.NewRecorder()
	m := MockDb{}
	s := New(&m)
	s.GetItem(w, req)
	res := w.Result()
	if res.StatusCode != 200 {
		t.Errorf("Wrong Status Code: %v", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	t.Errorf(string(data))
}
