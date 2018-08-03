package get

import (
	"context"
	"encoding/json"

	globals "github.com/sea350/ustart_go/globals"
	types "github.com/sea350/ustart_go/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

//ByID ...
func ByID(eclient *elastic.Client, imgID string) (types.Img, error) {
	ctx := context.Background()         //intialize context background
	var img types.Img                   //initialize type img
	searchResult, err := eclient.Get(). //Get returns doc type, index, etc.
						Index(globals.ImgIndex).
						Id(imgID).
						Do(ctx)

	if err != nil {
		return img, err
	}

	Err := json.Unmarshal(*searchResult.Source, &img) //unmarshal type RawMessage into img struct

	return img, Err

}
