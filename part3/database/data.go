package database

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"sync"
)

//Interface check
var _ DataKind = &Data{}

type DataKind interface {
	Update(key, value string) error
	Delete(key string) error
	Get(key string) (string, bool)
}

type Data struct {
	Filename string
	Stuff    map[string]string
	FileLock sync.Mutex
}

//Creates a new DataKind of type Data
func Initialize(filename string) (DataKind, error) {
	d := Data{
		Filename: filename,
		Stuff:    make(map[string]string),
	}

	rawData, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("Error reading file:" + err.Error())
	}
	var temp []map[string]string

	// Parse the JSON
	err = json.Unmarshal(rawData, &temp)
	if err != nil {
		return nil, errors.New("Error parsing JSON:" + err.Error())
	}

	if len(temp) != 1 {
		return nil, errors.New("Malformed Data, array of size: " + strconv.Itoa(len(temp)))
	}
	d.Stuff = temp[0]
	return &d, nil
}

func (d *Data) Get(key string) (string, bool) {
	value, ok := d.Stuff[key]
	return value, ok
}

func (d *Data) Update(key, value string) error {
	d.FileLock.Lock()
	defer d.FileLock.Unlock()
	d.Stuff[key] = value
	file, err := os.Create(d.Filename)
	if err != nil {
		//For prod I'd handle this graecfully but I want this to panic for my testing
		panic(err)
		//return err
	}
	err = json.NewEncoder(file).Encode([]map[string]string{d.Stuff})
	defer file.Close()
	if err != nil {
		//For prod I'd handle this graecfully but I want this to panic for my testing
		panic(err)
		//return err
	}
	return nil
}

func (d *Data) Delete(key string) error {
	d.FileLock.Lock()
	defer d.FileLock.Unlock()
	delete(d.Stuff, key)
	file, err := os.Create(d.Filename)
	if err != nil {
		//For prod I'd handle this graecfully but I want this to panic for my testing
		panic(err)
		//return err
	}
	err = json.NewEncoder(file).Encode([]map[string]string{d.Stuff})
	defer file.Close()
	if err != nil {
		//For prod I'd handle this graecfully but I want this to panic for my testing
		panic(err)
		//return err
	}
	return nil
}
