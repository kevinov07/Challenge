package constants

const ZINCSEARCH_ENDPOINT = "api/email-indexer/_search"
const ZINCSEARCH_PORT = "http://localhost:4080/"
const USERNAME = "admin"
const PASSWORD = "Complexpass#123"

const SERVER = "http://localhost:3000/"

var SEARCH_QUERY string = `{
    "search_type": "match",
    "query": {
        "term": "%s"
    },
    "sort_fields": ["date"],
    "from": 0,
    "max_results": %d,
    "_source": []
}`

var EMAIL_QUERY string = `{
	"search_type": "matchphrase",
	"query": {
		"term": "%s",
		"field": "message_id"
	},
	"sort_fields": [],
	"from": 0,
	"max_results": 1,
	"_source": []
}`
