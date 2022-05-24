package config

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/PubApiADN/pkg/logger"
)

const (
	logErrorCommunicationAPI      = "there was customPool communication error with: %s, Error: %s [Class: restClient][Method:%s]"
	logErrorUnmarshallInformation = "error when scanning the information an try convert to struct [Class: restClient][Method:%s]"
	getMethod                     = "Get"
	postMethod                    = "Post"
	putMethod                     = "Put"
)

type CustomRestClient struct {
}

func (customRestClient *CustomRestClient) Get(url string, apiName string, typeOfDataToMap interface{}) error {

	client := resty.New()

	response, err := client.R().EnableTrace().Get(url)

	if err != nil {
		logger.Errorf(fmt.Sprintf(logErrorCommunicationAPI, apiName, err, getMethod), err)
		return err
	}

	err = MapRequestToStruct(response.Body(), typeOfDataToMap)
	if err != nil {
		logger.Errorf(fmt.Sprintf(logErrorUnmarshallInformation, getMethod), err)
		return err

	}
	return err
}

// Post  the API in the specified URL with the body info
func (customRestClient *CustomRestClient) Post(url string, body interface{}, typeOfDataToMap interface{}, apiName string) error {

	logger.Infof("RestClient - Post , apiName= %s - body -> %+v , url =%s", apiName, body, url)
	client := resty.New()
	response, err := client.R().SetBody(body).Post(url)

	if err != nil {

		logger.Errorf(fmt.Sprintf(logErrorCommunicationAPI, apiName, err, postMethod), err)
		return err
	}

	err = MapRequestToStruct(response.Body(), typeOfDataToMap)

	if err != nil {

		logger.Errorf(fmt.Sprintf(logErrorUnmarshallInformation, postMethod), err)
		return err
	}

	return err
}

// Put update the API in the specified URL with the body info
func (customRestClient *CustomRestClient) Put(url string, body interface{}, apiName string) error {

	logger.Infof("RestClient - PUT - body -> %+v", body)
	client := resty.New()
	response, err := client.R().SetBody(body).Put(url)

	if err != nil {

		logger.Errorf(fmt.Sprintf(logErrorCommunicationAPI, apiName, err, putMethod), err)
		return err
	}

	logger.Infof("RestClient - PUT - responseStatus -> %v", response.Status())

	return err
}
