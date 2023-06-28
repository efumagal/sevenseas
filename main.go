package main

import (
	_ "bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

type Value struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

var data = `{
	"AEAJM": {
		"name": "Ajman",
		"city": "Ajman",
		"country": "United Arab Emirates",
		"alias": [],
		"regions": [],
		"coordinates": [
		  55.5136433,
		  25.4052165
		],
		"province": "Ajman",
		"timezone": "Asia/Dubai",
		"unlocs": [
		  "AEAJM"
		],
		"code": "52000"
	  },
	  "AEAUH": {
		"name": "Abu Dhabi",
		"coordinates": [
		  54.37,
		  24.47
		],
		"city": "Abu Dhabi",
		"province": "Abu Z¸aby [Abu Dhabi]",
		"country": "United Arab Emirates",
		"alias": [],
		"regions": [],
		"timezone": "Asia/Dubai",
		"unlocs": [
		  "AEAUH"
		],
		"code": "52001"
	  }
}`

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func decodeStream(r io.Reader) error {
	var m runtime.MemStats
	dec := json.NewDecoder(r)
	t, err := dec.Token()
	if err != nil {
		return err
	}
	if t != json.Delim('{') {
		return fmt.Errorf("expected {, got %v", t)
	}
	for dec.More() {
		t, err := dec.Token()
		if err != nil {
			return err
		}
		key := t.(string)

		var value Value
		if err := dec.Decode(&value); err != nil {
			return err
		}
		log.Printf("key %q, value %#v\n", key, value)

		runtime.ReadMemStats(&m)
		log.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	}
	return nil
}

func encodeStream(r *os.File) error {
	var m runtime.MemStats
	enc := json.NewEncoder(r)

	r.Write([]byte{'{'})

	for i := 1; i <= 100000; i++ {
		r.WriteString(fmt.Sprintf("\"%d\": ", i))
		var value Value
		value.Code = fmt.Sprint(i)
		value.Name = fmt.Sprint(i)
		value.Alias = []string{}
		value.Regions = []string{}
		value.Coordinates = []float64{55.5136433, 25.4052165}

		if err := enc.Encode(&value); err != nil {
			return err
		}
		r.Write([]byte{','})

		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB\n", bToMb(m.Alloc))
	}
	r.Write([]byte{'}'})
	return nil
}

func main() {
	fmt.Println(decodeStream(strings.NewReader(data)))

	fileName := "ports.json"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
	}
	defer f.Close()

	fmt.Println(decodeStream(f))

	/*
		fileName = "ports_big.json"
		fw, err := os.Create(fileName)
		if err != nil {
			log.Fatalf("Error to read [file=%v]: %v", fileName, err.Error())
		}
		defer fw.Close()

		encodeStream(fw)
	*/
}
