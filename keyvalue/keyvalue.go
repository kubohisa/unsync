package keyvalue

import (
	"reflect"
	"strconv"
	"encoding/json"
)

	var data map[string]string
	var lock map[string]bool

/*

*/	

func New (capacity int) {
	if data == nil {
		data = make(map[string]string, capacity)
	}
	if lock == nil {
		lock = make(map[string]bool, capacity)
	}
}

/*

*/

func Set (key string, value string) bool {
	if (lock[key] == false) {
		data[key] = value
		lock[key] = true
		
		return true
	}
	return false
}

func Get (key string) string {
	nilCheck(key)
	
	return data[key]
}

func SetJson (key string, value map[string]interface{}) bool {
	if (lock[key] == false) {
		bytes, err := json.Marshal(value)
		if err != nil {
			return false
		}
		
		lock[key] = true
		data[key] = bytes
		return true
	}
	return false
}

func GetJson (key string) map[string]interface{} {
	nilCheck(key)
	
	var mapData map[string]interface{}
	
//	json.Unmarshal(data[key], &mapData)
	
	decoder := json.NewDecoder(bytes.NewReader(data[key]))
	decoder.UseNumber()
	err := decoder.Decode(&mapData);
	
	return mapData
}

func nilCheck(key string) string {
	value, isThere = data[key]
	if (isThere == nil) {
		data[key] = ""
		lock[key] = false
	}
}

/*
func GetInt (key string) bool, int {
	nilCheck(key)
	
	rt := reflect.TypeOf(data[key])
	if (rt.Kind() != reflect.Int) {
		return false, nil
	}
	
	toInt, _ = strconv.Atoi(data[key])
	return true, toInt
}
*/

func Delete (key string) {
	value, isThere = data[key]
	if (isThere != nil) {
		delete(data, string)
	}
	
	return
}

func Empty (key string) bool {
	value, isThere = data[key]
	if (isThere == nil || data[key] == "") {
		return false
	}
	
	return true
}

func Check (key string) bool {
	return Empty (key)
}


/*

*/

func Lock (key string) {
	lock[key] = true
	
	return
}

func Unlock (key string) {
	lock[key] = false
	
	return
}

func LockCheck (key string) {
	return lock[key]
}

