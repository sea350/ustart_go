USER='elastic'
PASS='elasticpassword'
URL='ustart.today'
PORT='9200'
JSON='Content-Type:application/json'
INDEX='/test-event_data/EVENT'
CURL_BASE='http://'$USER':'$PASS'@'$URL':'$PORT

 


curl -X PUT  $CURL_BASE/INDEX/'_mapping/_doc' -H 'Content-Type: application/json' -d'
{
  "properties": {
    "EVENT\": {
      "properties": {
        "last": { 
          "type": "text"
        }
      }
    },
    "user_id": {
      "type": "keyword",
      "ignore_above": 100 
    }
  }
}
'
