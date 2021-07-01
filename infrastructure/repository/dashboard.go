package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"mongo-admin-backend/api/presenter/dashboard"
	"mongo-admin-backend/pkg/contextWrapper"
	"mongo-admin-backend/pkg/digest"
	"sync"
)

// mongoDB repo.
type DashBoardRepo struct {
	apiURL string
}

func NewDashBoardRepo(apiUrl string) *DashBoardRepo {
	return &DashBoardRepo{
		apiURL: apiUrl,
	}
}

func (r *DashBoardRepo) Get(u, p, groupId string) (*map[string]interface{}, error) {
	return getDashBoardData(contextWrapper.Ctx, u, p, r.apiURL, groupId)
}

func getDashBoardData(ctx context.Context, u, p, baseURL, groupId string) (*map[string]interface{}, error) {
	accessListPath := "/groups/" + groupId + "/accessList"
	databaseUsersPath := "/groups/" + groupId + "/databaseUsers"
	clusterDetailsPath := "/groups/" + groupId + "/clusters"
	processesDetailsPath := "/groups/" + groupId + "/processes"

	host := baseURL
	serviceCalls := []dashboard.Service{
		dashboard.Service{
			Name:             "numAccess",
			Host:             host,
			Url:              accessListPath,
			Method:           "GET",
			RequestBody:      nil,
			Process_Response: process_accessList,
		},
		dashboard.Service{
			Name:             "numDbUsers",
			Host:             host,
			Url:              databaseUsersPath,
			Method:           "GET",
			RequestBody:      nil,
			Process_Response: process_databaseUsers,
		},
		dashboard.Service{
			Name:             "clusterDetails",
			Host:             host,
			Url:              clusterDetailsPath,
			Method:           "GET",
			RequestBody:      nil,
			Process_Response: process_clusters,
		},
		dashboard.Service{
			Name:             "processDetails",
			Host:             host,
			Url:              processesDetailsPath,
			Method:           "GET",
			RequestBody:      nil,
			Process_Response: process_processesList,
		},
	}
	result := CallServices(u, p, serviceCalls)
	var returnObj = make(map[string]interface{})
	for _, response := range result.Results {
		returnObj[response.Name] = response.Responsebody
	}

	return &returnObj, nil
}

func CallServices(u, p string, services []dashboard.Service) dashboard.ServiceCallResult {
	output := make(chan dashboard.ServiceCallResult)
	input := make(chan dashboard.ServiceCallResponse)
	var wg sync.WaitGroup
	go handleResults(input, output, &wg)
	defer close(output)
	for _, service := range services {
		wg.Add(1)
		go ConcurrentCalls(u, p, service, input)
	}

	wg.Wait()
	close(input)
	return <-output
}
func handleResults(input chan dashboard.ServiceCallResponse, output chan dashboard.ServiceCallResult, wg *sync.WaitGroup) {
	var results dashboard.ServiceCallResult
	for result := range input {
		results.Results = append(results.Results, result)
		wg.Done()
	}
	output <- results
}

func ConcurrentCalls(u, p string, service dashboard.Service, input chan dashboard.ServiceCallResponse) {
	result := CallService(u, p, service)
	input <- result
}

func CallService(u, p string, service dashboard.Service) dashboard.ServiceCallResponse {
	fmt.Println("The service is called " + service.Url)
	fmt.Println(service)
	//time.Sleep(time.Second * 1)
	digestor := digest.NewDigestor(u, p)
	path := service.Url
	var reqBody []byte
	if service.RequestBody != nil {
		reqBody, _ = json.Marshal(service.RequestBody)
	} else {
		reqBody = nil
	}
	j, _ := json.Marshal(reqBody)
	result, err2 := digestor.Digest(service.Host, path, service.Method, j)
	var errDetails dashboard.ErrorDetail
	var responseBody interface{}
	var err3 error
	if err2 != nil {
		var val map[string]interface{}
		_ = json.Unmarshal(result, &val)
		jsonEncoded, _ := json.Marshal(val)
		_ = json.Unmarshal(jsonEncoded, &errDetails)
		responseBody = errDetails
		err3 = err2
	} else {
		responseBody, err3 = service.Process_Response(result)
	}

	return dashboard.ServiceCallResponse{
		ServiceCall:  service,
		Name:         service.Name,
		Responsebody: responseBody,
		Error:        err3,
	}
}

//Callback functions
func process_accessList(response []byte) (interface{}, error) {
	var resp map[string]interface{}
	marshalError := json.Unmarshal(response, &resp)
	type result struct{ TotalCount int }
	var resObj result
	jsonEncoded, _ := json.Marshal(resp)
	_ = json.Unmarshal(jsonEncoded, &resObj)
	fmt.Println(resObj)
	return resObj, marshalError
}

func process_databaseUsers(response []byte) (interface{}, error) {
	var resp map[string]interface{}
	marshalError := json.Unmarshal(response, &resp)
	type result struct{ TotalCount int }
	var resObj result
	jsonEncoded, _ := json.Marshal(resp)
	_ = json.Unmarshal(jsonEncoded, &resObj)
	fmt.Println(resObj)
	return resObj, marshalError
}

func process_clusters(response []byte) (interface{}, error) {
	var resp map[string]interface{}
	marshalError := json.Unmarshal(response, &resp)
	var resObj []dashboard.ClusterDetails
	jsonEncoded, _ := json.Marshal(resp["results"])
	_ = json.Unmarshal(jsonEncoded, &resObj)
	return resObj, marshalError
}

func process_processesList(response []byte) (interface{}, error) {
	var resp map[string]interface{}
	marshalError := json.Unmarshal(response, &resp)
	var resObj []dashboard.ProcessDetails
	jsonEncoded, _ := json.Marshal(resp["results"])
	_ = json.Unmarshal(jsonEncoded, &resObj)
	return resObj, marshalError
}
