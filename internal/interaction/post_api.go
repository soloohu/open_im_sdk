package interaction

import (
	"encoding/json"
	"errors"
	"github.com/soloohu/open_im_sdk/open_im_sdk_callback"
	"github.com/soloohu/open_im_sdk/pkg/common"
	"github.com/soloohu/open_im_sdk/pkg/log"
	"github.com/soloohu/open_im_sdk/pkg/network"
	"github.com/soloohu/open_im_sdk/pkg/utils"
	"time"
)

//no share
type PostApi struct {
	token      string
	apiAddress string
}

func NewPostApi(token string, apiAddress string) *PostApi {
	return &PostApi{token: token, apiAddress: apiAddress}
}

func (p *PostApi) PostFatalCallback(callback open_im_sdk_callback.Base, url string, data interface{}, output interface{}, operationID string) {
	log.Info(operationID, utils.GetSelfFuncName(), p.apiAddress, url, data)
	content, err := network.Post2Api(p.apiAddress+url, data, p.token)
	common.CheckErrAndRespCallback(callback, err, content, output, operationID)
}

func (pe *postErr) OnError(errCode int32, errMsg string) {
	pe.err = errors.New(errMsg)
}

func (pe *postErr) OnSuccess(data string) {
}

type postErr struct {
	err error
}

func (p *PostApi) PostReturn(url string, req interface{}, output interface{}) error {
	content, err := network.Post2Api(p.apiAddress+url, req, p.token)
	if err != nil {
		utils.Wrap(err, "post failed "+p.apiAddress+url)
	}
	err = common.CheckErrAndResp(err, content, output)
	return utils.Wrap(err, "CheckErrAndResp failed ")
}

func (p *PostApi) Post2UnmarshalRespReturn(url string, req interface{}, output interface{}) error {
	content, err := network.Post2Api(p.apiAddress+url, req, p.token)
	if err != nil {
		utils.Wrap(err, "post failed "+p.apiAddress+url)
	}
	err = json.Unmarshal(content, output)
	return utils.Wrap(err, "Unmarshal failed ")
}
func (p *PostApi) PostReturnWithTimeOut(url string, req interface{}, output interface{}, timeOut time.Duration) error {
	content, err := network.PostWithTimeOut(p.apiAddress+url, req, p.token, timeOut)

	err1 := common.CheckErrAndResp(err, content, output)
	if err1 != nil {
		log.Error("", "PostReturn failed ", err1.Error(), "input: ", string(content), " req:", req)
	}
	return err1
}
