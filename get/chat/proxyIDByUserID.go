package get

import (
	elastic "github.com/olivere/elastic"
	getUser "github.com/sea350/ustart_go/get/user"
)

//ProxyIDByUserID ...
func ProxyIDByUserID(eclient *elastic.Client, userID string) (string, error) {

	// ctx := context.Background()

	// termQuery := elastic.NewTermQuery("DocID", strings.ToLower(userID))
	// searchResult, err := eclient.Search().
	// 	Index(globals.ProxyMsgIndex).
	// 	Query(termQuery).
	// 	Do(ctx)

	// var proxyID string
	// if err != nil {
	// 	return proxyID, err
	// }

	// if searchResult.TotalHits() == 0 {
	// 	proxy := types.ProxyMessages{DocID: userID, Class: 1}
	// 	result, err := eclient.Index().
	// 		Index(globals.ProxyMsgIndex).
	// 		Type(globals.ProxyMsgType).
	// 		BodyJson(proxy).
	// 		Do(ctx)

	// 	return result.Id, err
	// }
	// if searchResult.TotalHits() > 1 {
	// 	/*
	// 		for _, element := range searchResult.Hits.Hits {
	// 			prx, _ := ProxyMsgByID(eclient, element.Id)
	// 			log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 			log.Println(prx)

	// 		}
	// 	*/
	// 	return proxyID, errors.New("multiple proxies found")
	// }

	// for _, element := range searchResult.Hits.Hits {

	// 	proxyID = element.Id
	// 	break
	// }

	usr, err := getUser.UserByID(eclient, userID)

	return usr.ProxyMessagesID, err

}
