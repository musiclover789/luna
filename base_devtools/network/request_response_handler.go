package network

import (
	"sync"
)

type Data struct {
	Request      map[string]interface{}
	Response     map[string]interface{}
}

type DataProcessor struct {
	DataMap map[string]*Data
	Lock    sync.Mutex
}

func NewDataProcessor() *DataProcessor {
	return &DataProcessor{
		DataMap: make(map[string]*Data),
		Lock:    sync.Mutex{},
	}
}

func (dp *DataProcessor) ProcessRequest(requestID string, requestData map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	dp.Lock.Lock()
	data, ok := dp.DataMap[requestID]
	if !ok {
		data = &Data{}
		dp.DataMap[requestID] = data
	}
	data.Request = requestData
	dp.Lock.Unlock()

	if data.Response != nil {
		return dp.processMatchingData(requestID, data)
	}
	return nil, nil
}

func (dp *DataProcessor) ProcessResponse(requestID string, responseData map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	dp.Lock.Lock()
	data, ok := dp.DataMap[requestID]
	if !ok {
		data = &Data{}
		dp.DataMap[requestID] = data
	}
	data.Response = responseData
	dp.Lock.Unlock()

	if data.Request != nil{
		return dp.processMatchingData(requestID, data)
	}
	return nil, nil
}


func (dp *DataProcessor) processMatchingData(requestID string, data *Data) (map[string]interface{}, map[string]interface{}) {
	dp.Lock.Lock()
	delete(dp.DataMap, requestID)
	dp.Lock.Unlock()
	return data.Request, data.Response
}
