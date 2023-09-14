package globalFunction

import (
	"fmt"
	"github.com/nulla-vis/golang-fiber-template/core/helper"
	"reflect"
	"strings"
)

func GetMessage(code string, replacements interface{}) map[string]interface{}{
	langEn := helper.LangEn()
	langId := helper.LangId()

	mapMessage := map[string]interface{} {
		"code": code,
		"en": code,
		"id": code,
	}

	msgEn := langEn[code]
	msgId := langId[code]

	if msgEn != nil {
		mapMessage["en"] = fmt.Sprint(msgEn)
	}

	if msgId != nil {
		mapMessage["id"] = fmt.Sprint(msgId)
	}



	replacementType := reflect.TypeOf(replacements).Kind()

	if replacements == "" {
		//do nothing
	} else if replacementType == reflect.String {
		mapMessage["en"] = strings.Replace(fmt.Sprint(mapMessage["en"]), "%s", fmt.Sprint(replacements), 1)
		mapMessage["id"] = strings.Replace(fmt.Sprint(mapMessage["id"]), "%s", fmt.Sprint(replacements), 1)
	} else if replacementType == reflect.Slice {
		s := reflect.ValueOf(replacements)
		ret := make([]interface{}, s.Len())
		for i:=0; i<s.Len(); i++ {
			ret[i] = s.Index(i).Interface()
			mapMessage["en"] = strings.Replace(fmt.Sprint(mapMessage["en"]), "%s", fmt.Sprint(ret[i]), 1)
			mapMessage["id"] = strings.Replace(fmt.Sprint(mapMessage["id"]), "%s", fmt.Sprint(ret[i]), 1)
		}
	}

	return mapMessage
}
// Convert data from database (if the value in byte slice)
func ConvertByteSlicesToStrings(data interface{}) {
    // Check if data is a slice of maps
    maps, isSliceOfMaps := data.([]map[string]interface{})
    if !isSliceOfMaps {
        // Data is not in the expected format
        return
    }

    // Iterate through the maps
    for _, m := range maps {
        for key, value := range m {
            byteSlice, isByteSlice := value.([]byte)
            if isByteSlice {
                // Convert the byte slice to a string
                m[key] = string(byteSlice)
            }
        }
    }
}