package elasticsearchLib

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"tampayang-backend/core/helper"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var EsClient *elasticsearch.Client

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : InitElasticsearchClient
* Kegunaan      : RAW ElasticSearch Function
* Parameter     : -
* Balikan       : Update variable Client
 */
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

// ////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : CheckPing
* Kegunaan      : untuk cek koneksi ke ElasticSearch dan memberitahukan hasil koneksinya.
* Parameter     : client
* Balikan       : Bool
* Dokumentasi :
*   > Mappings Object : https://www.elastic.co/guide/en/elasticsearch/reference/7.x/mapping.html
 */
func CheckPing() bool {
	// Use the client's Ping method to check Elasticsearch's availability.
	res, err := EsClient.Ping()
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

// ////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : CreateIndex
* Kegunaan      : untuk membuat index tanpa menambahkan data
* Parameter     : index (string), body (string), data (object)
* Balikan       : Object
* Dokumentasi :
*   > Mappings Object : https://www.elastic.co/guide/en/elasticsearch/reference/7.x/mapping.html
 */
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

// ////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : DeleteIndex
* Kegunaan      : untuk menghapus index
* Parameter     : index (string)
* Balikan       : Object
* Dokumentasi :
*   > indices.delete : https://www.elastic.co/guide/en/elasticsearch/client/javascript-api/current/api-reference.html#_indices_delete
 */
// DeleteIndex deletes an Elasticsearch index by name.
func DeleteIndex(index string) (map[string]interface{}, error) {
	// Create the DeleteIndex request
	req := esapi.IndicesDeleteRequest{
		Index: []string{index},
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
		"result":     "Index deleted successfully",
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : insertData
* Kegunaan      : untuk membuat index sekaligus menambahkan data kedalam index yang baru dibuat ATAU update data
* Parameter     : index (string), id (string), data (object)
* Balikan       : Object
 */
// InsertData inserts data into an Elasticsearch index with the specified ID.
func InsertData(index string, id string, data map[string]interface{}) (map[string]interface{}, error) {
	// Serialize the data to JSON
	reqBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Create the Index request
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(string(reqBody)), // Serialize the request body to JSON
	}

	// Perform the request
	res, err := req.Do(context.Background(), EsClient)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Read the response body to get more information about the result
	var responseBody map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		panic(err)
	}

	// panic(responseBody)

	if res.StatusCode < 200 || res.StatusCode > 299 {
		// Handle the error here
		panic(responseBody)
	}

	// Prepare the result
	result := map[string]interface{}{
		"statusCode": res.StatusCode,
		"result":     responseBody,
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : updateIndex
* Kegunaan      : untuk merubah index satu per satu
* Parameter     : index (string), id (string), data (object)
* Balikan       : Object
 */
// UpdateIndex updates a document in an Elasticsearch index with the specified ID.
func UpdateIndex(index string, id string, data map[string]interface{}) (map[string]interface{}, error) {
	// Serialize the data to JSON
	reqBody := map[string]interface{}{
		"doc": data,
	}

	// Serialize the request body to JSON
	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// Create the Update request
	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: id,
		Body:       strings.NewReader(string(reqBodyJSON)),
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

	if res.StatusCode < 200 || res.StatusCode > 299 {
		// Handle the error here
		panic(responseBody)
	}

	// Prepare the result
	result := map[string]interface{}{
		"statusCode": res.StatusCode,
		"result":     responseBody,
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : existIndex
* Kegunaan      : untuk memeriksa apakah ada "index" yang dimaksud
* Parameter     : index (string)
* Balikan       : Bool
 */
// IndexExists checks if an index already exists in Elasticsearch.
func IndexExists(indexName string) (bool, error) {
	// Create the IndicesExistsRequest
	req := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}

	// Perform the request
	res, err := req.Do(context.Background(), EsClient)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	// Check the response status code
	if res.IsError() {
		return false, nil // Index does not exist
	}

	return true, nil // Index exists
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : DocumentExists
* Kegunaan      : untuk memeriksa apakah ada data / document di "index" dengan "id" yang dimaksud ?
* Parameter     : index (string), id (string)
* Balikan       : Bool
 */
// DocumentExists checks if a document with the specified ID exists in the given index.
func DocumentExists(indexName, documentID string) (bool, error) {
	// Create the Get request to retrieve the document by ID
	req := esapi.GetRequest{
		Index:      indexName,
		DocumentID: documentID,
	}

	// Perform the request
	res, err := req.Do(context.Background(), EsClient)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	// Check the response status code
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return false, nil // Document does not exist
	}

	return true, nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : deleteData
* Kegunaan      : untuk menghapus data / dokumen dengan  "index" dan "id" yang dimaksud.
* Parameter     : index (string), id (string)
* Balikan       : Object
 */
// DeleteDocument deletes a document with the specified ID in the given index.
func DeleteDocument(indexName, documentID string) (map[string]interface{}, error) {
	// Create the Delete request to delete the document by ID
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: documentID,
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

	if res.StatusCode < 200 || res.StatusCode > 299 {
		panic(responseBody)
	}

	// Prepare the result
	result := map[string]interface{}{
		"statusCode": res.StatusCode,
		"result":     responseBody,
	}

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////

// ////////////////////////////////////////////////////////////////////////////////////////////////////
/*
* Nama Fungsi   : GetDocument
* Kegunaan      : untuk mengambil satu dokumen dengan kirim nama index dan id dokumennya.
* Parameter     : index (string), id (string).
* Balikan       : Object
 */
// GetDocument retrieves a document from the specified index in Elasticsearch by ID.
func GetDocument(indexName, documentID string) (map[string]interface{}, error) {
	// Create the Get request to retrieve the document by ID
	req := esapi.GetRequest{
		Index:      indexName,
		DocumentID: documentID,
	}

	// Perform the request
	res, err := req.Do(context.Background(), EsClient)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Prepare the result
	result := map[string]interface{}{
		"statusCode": res.StatusCode,
		"resultEs":   make(map[string]interface{}),
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		// If the status code is not 200 (OK), set the error information in the result
		result["resultEs"].(map[string]interface{})["statusCode"] = res.StatusCode
		result["resultEs"].(map[string]interface{})["error"] = "Document not found"
		return result, nil
	}

	// Read the response body to get the document data
	var responseBody map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return nil, err
	}

	// Set the document data in the result
	result["resultEs"] = responseBody["_source"]

	return result, nil
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
