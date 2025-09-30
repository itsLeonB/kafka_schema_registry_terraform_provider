package restapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type schemaRegistryClient struct {
	httpClient *http.Client
	createURI  string
	updateURI  string
	readURI    string
	deleteURI  string
	subject    string
	schema     string
}

const subjectsString = "/subjects/"
const versionsString = "/versions/"

func NewSchemaRegistryClient(uri string, subject string, schema string) (*schemaRegistryClient, error) {
	client := schemaRegistryClient{
		createURI:  uri + subjectsString + subject + versionsString,
		updateURI:  uri + subjectsString + subject + versionsString,
		readURI:    uri + subjectsString + subject + versionsString,
		deleteURI:  uri + subjectsString + subject,
		subject:    subject,
		schema:     schema,
		httpClient: &http.Client{},
	}

	return &client, nil
}

func (client schemaRegistryClient) createSubject() error {
	jsonData := map[string]string{"schema": client.schema}
	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	response, err := http.Post(client.createURI, "application/vnd.schemaregistry.v1+json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response code is %d: %s", response.StatusCode, data)
	}

	return nil
}

func (client schemaRegistryClient) deleteSubject() error {
	request, err := http.NewRequest(http.MethodDelete, client.deleteURI, nil)
	if err != nil {
		return err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response code is %d: %s", response.StatusCode, data)
	}

	return nil
}
