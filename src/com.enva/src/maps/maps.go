package maps

import "fmt"



//							  Key     Value
func CreateMap() (theMap map[string] string) {
    theMap = makeMap()


	for key,value := range theMap {
		fmt.Println(key + "=>" + value)
	}
	return
}
/**
 * Create map in a simple way
 */
func makeMap() (theMap map[string] string)  {
    myMap := map[string]string {
    	"one":"one value",
    	"two":"one value",
    	"three":"one value",
    	"four":"one value",
	}
	return myMap
}

func UpdateMapValue(key string,value string)  {
	theMap := makeMap();
	theMap[key]=value;
}

func RemoveMapValue(key string)  {
	theMap := makeMap();
	delete(theMap,key)
}

func keyExist(key string) (val string) {
	theMap := makeMap();
	if value,exists := theMap[key]; exists {
		val = value
	}
	return
}