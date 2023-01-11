package keyvalue

import (
	"reflect"
	"strconv"
	"encoding/json"
	"sync"
)

	var data sync.Map

/*

*/	

func New (capacity int) {
	if data == nil {
		data = sync.Map{}
		lock = true
	}
}

/*

*/

func Set (key string, value string) {
	data.Store(key, value)
}

func Get (key string) string {
	d, ok := data.Load(key)
	if ok {
		return d.(string)
	}
	
	return nil
}

func SetJson (key string, value map[string]interface{}) bool {
	bytes, err := json.Marshal(value)
	if err != nil {
		return false
	}
		
	data.Store(key, value)
	return true
}

func GetJson (key string) map[string]interface{} {
	var mapData map[string]interface{}
	
	d, ok := data.Load(key)
	if ok {
		decoder := json.NewDecoder(bytes.NewReader(d.(string)))
		decoder.UseNumber()
		err := decoder.Decode(&mapData);
	}	
	return mapdata
}

func Delete (key string) {
	data.Delete(key)
}

func Empty (key string) bool {
	d, ok := data.Load(key)	
	return ok
}

func Check (key string) bool {
	return Empty (key)
}
