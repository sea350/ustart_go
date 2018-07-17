chmod +x ./userPutMapping.sh
USER='elastic'
PASS='elasticpassword'
URL='localhost'
PORT='9200/'
JSON='Content-Type:application/json'
INDEX='test-user_data/USER'
CURL_BASE='http://'$USER':'$PASS'@'$URL':'$PORT

 
#curl -XPOST "$CURL_BASE""$TEST_USER"'/0?pretty=true' -H "$JSON" -d "$USER_TEMPLATE"
# curl -XGET $CURL_BASE'/_cat/indices?v&pretty=true&s=index'
echo $INDEX
curl -X PUT  "$CURL_BASE""$INDEX"'/_mapping/_doc' -H 'Content-Type: application/json' -d'
{
  "properties": {
    "USER": {
      "properties": {
        "type":"nested"
        "LoginWarnings": { 
          "type": "text"
        }
      }
    } 
    }
  }
}
'
