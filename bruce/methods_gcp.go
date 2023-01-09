package bruce

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/tukaianirban/sdk.go.common/errs"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

type GcpIndexer struct {
	ProjectID string
	Client    *secretmanager.Client
}

func initConfigModeGCP(projectId string) *GcpIndexer {

	return &GcpIndexer{
		ProjectID: projectId,
	}
}

func (indexer GcpIndexer) Close() {
	indexer.Client.Close()
}

func (indexer GcpIndexer) get(key string) ([]byte, error) {

	ctx := context.Background()
	name := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", indexer.ProjectID, key)

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// create the API request
	req := secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	result, err := client.AccessSecretVersion(ctx, &req)
	if err != nil {
		return nil, err
	}

	return result.Payload.Data, nil
}

func (indexer GcpIndexer) GetBoolean(key string) (bool, error) {

	resultbytes, err := indexer.get(key)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(string(resultbytes))
}

func (indexer GcpIndexer) GetString(key string) (string, error) {

	resultbytes, err := indexer.get(key)
	if err != nil {
		return "", err
	}

	// convert the payload data to string
	resultString := string(resultbytes)

	return resultString, nil
}

func (indexer GcpIndexer) GetInt(key string) (int, error) {

	resultbytes, err := indexer.get(key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(resultbytes))
}

func (indexer GcpIndexer) GetFloat64(key string) (float64, error) {

	resultbytes, err := indexer.get(key)
	if err != nil {
		return 0.0, err
	}

	return strconv.ParseFloat(string(resultbytes), 64)
}

func (indexer GcpIndexer) GetArray(key string) ([]interface{}, error) {

	return nil, errs.ErrBruceTypeNotSupported
}

func (indexer GcpIndexer) GetMap(key string) (map[string]interface{}, error) {

	resultbytes, err := indexer.get(key)
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})

	err = json.Unmarshal(resultbytes, &res)
	return res, err
}
