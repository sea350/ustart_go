package uses

import (
	elastic "gopkg.in/olivere/elastic.v5"
)

//MultiWildCardQuery ... takes in an existing bool query, a field, array of strings to add, and a should boolead and then adds each string to the query as a wildcard
//if should == false then runs MUST
func MultiWildCardQuery(query *elastic.BoolQuery, field string, searchTerms []string, should bool) *elastic.BoolQuery {
	for _, element := range searchTerms {
		if should {
			query = query.Should(elastic.NewWildcardQuery(field, `*`+element+`*`))
		} else {
			query = query.Must(elastic.NewWildcardQuery(field, `*`+element+`*`))
		}
	}
	return query
}
