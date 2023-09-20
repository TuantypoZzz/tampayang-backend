// myelasticsearch/elasticsearch.go

package elasticsearchLib

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/nulla-vis/golang-fiber-template/core/helper"
)

var EsClient *elasticsearch.Client

func InitElasticsearchClient() error {
	// Your configuration logic here
	configuration := helper.ConfigJson()
	scheme := fmt.Sprint(configuration["elasticsearch"].(map[string]interface{})["server"].(map[string]interface{})["scheme"])
	host := fmt.Sprint(configuration["elasticsearch"].(map[string]interface{})["server"].(map[string]interface{})["host"])
	port := fmt.Sprint(configuration["elasticsearch"].(map[string]interface{})["server"].(map[string]interface{})["port"])
	username := fmt.Sprint(configuration["elasticsearch"].(map[string]interface{})["auth"].(map[string]interface{})["username"])
	password := fmt.Sprint(configuration["elasticsearch"].(map[string]interface{})["auth"].(map[string]interface{})["password"])

	esConfig := elasticsearch.Config{
		Addresses: []string{scheme + "://" + host + ":" + port},
		Username:  username,
		Password:  password,
	}

	client, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		return err
	}

	EsClient = client
	return nil
}

func CheckPing(client *elasticsearch.Client) bool {
    // Use the client's Ping method to check Elasticsearch's availability.
    res, err := client.Ping()
    if err != nil {
        // An error occurred while pinging Elasticsearch, indicating it's not available.
        return false
    }

    // Check the response status code.
    if res.StatusCode != http.StatusOK {
        // Elasticsearch returned a non-OK status code, indicating it's not available.
        return false
    }

    // Elasticsearch is available.
    return true
}

// CreateIndex creates an Elasticsearch index with the specified name and properties.
func CreateIndex(index string, properties map[string]interface{}) (map[string]interface{}, error) {
    // Prepare the request body
    reqBody := map[string]interface{}{
        "mappings": map[string]interface{}{
            "properties": properties,
        },
    }

    // Create the CreateIndex request
    req := esapi.IndicesCreateRequest{
        Index: index,
        Body:  helper.MakeReader(reqBody), // Serialize the request body to JSON
    }

    // Perform the request
    res, err := req.Do(context.Background(), EsClient)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    // Read the response body to get more information about the result
    var responseBody map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
        return nil, err
    }

    if res.StatusCode != 200 {
		panic(fmt.Sprint(responseBody["error"].(map[string]interface{})["reason"]))
    }

    // Prepare the result
    result := map[string]interface{}{
        "statusCode": res.StatusCode,
        "result":     responseBody,
    }

    return result, nil
}

// DeleteIndex deletes an Elasticsearch index by name.
func DeleteIndex(client *elasticsearch.Client, index string) (map[string]interface{}, error) {
    // Create the DeleteIndex request
    req := esapi.IndicesDeleteRequest{
        Index: []string{index},
    }

    // Perform the request
    res, err := req.Do(context.Background(), client)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    // Read the response body to get more information about the result
    var responseBody map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
        return nil, err
    }

	if res.StatusCode != 200 {
		panic(fmt.Sprint(responseBody["error"].(map[string]interface{})["reason"]))
    }

    // Prepare the result
    result := map[string]interface{}{
        "statusCode": res.StatusCode,
        "result":     "Index deleted successfully",
    }

    return result, nil
}


