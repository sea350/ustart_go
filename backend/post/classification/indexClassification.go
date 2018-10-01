package seniors

import (
	"context"
	"log"

	globals "github.com/sea350/ustart_go/backend/globals"
	types "github.com/sea350/ustart_go/backend/types"
	elastic "gopkg.in/olivere/elastic.v5"
)

func IndexClassification(eclient *elastic.Client, ID string, classificationStruct types.Classification) {
	ctx := context.Background()

	_, err := eclient.Index().
		Index(globals.ClassificationIndex).
		Type(globals.ClassificationType).
		Id(ID).
		BodyJson(classificationStruct).
		Do(ctx)

	if err != nil {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(err)
	}
}
