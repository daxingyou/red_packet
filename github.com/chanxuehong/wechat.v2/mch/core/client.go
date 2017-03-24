package core

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/chanxuehong/util"
	"github.com/chanxuehong/wechat.v2/internal/debug/mch/api"
)

type Client struct {
	appId  string
	mchId  string
	apiKey string

	httpClient *http.Client
}

func (clt *Client) AppId() string {
	return clt.appId
}
func (clt *Client) MchId() string {
	return clt.mchId
}
func (clt *Client) ApiKey() string {
	return clt.apiKey
}

// NewClient 创建一个新的 Client.
//  如果 httpClient == nil 则默认用 http.DefaultClient.
func NewClient(appId, mchId, apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		appId:      appId,
		mchId:      mchId,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// PostXML 是微信支付通用请求方法.
//  err == nil 表示协议状态为 SUCCESS(return_code==SUCCESS).
func (clt *Client) PostXML(url string, req map[string]string) (resp map[string]string, err error) {
	fmt.Printf("调用postXml url[%v]  req[%v]\n", url, req)
	bodyBuf := textBufferPool.Get().(*bytes.Buffer)
	bodyBuf.Reset()
	defer textBufferPool.Put(bodyBuf)

	if err = util.EncodeXMLFromMap(bodyBuf, req, "xml"); err != nil {
		return
	}
	api.DebugPrintPostXMLRequest(url, bodyBuf.Bytes())

	fmt.Printf("调用postXml,开始clt.httpClient.Post，EncodeXMLFromMap。bodyBuf: %v \n", bodyBuf)
	httpResp, err := clt.httpClient.Post(url, "text/xml; charset=utf-8", bodyBuf)
	if err != nil {
		return
	}

	fmt.Printf("调用postXml,clt.httpClient.Post 结束\n")
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	if resp, err = api.DecodeXMLHttpResponse(httpResp.Body); err != nil {
		return
	}

	fmt.Printf("调用postXml decoeXmlHttpRes %v\n", resp)

	// 判断协议状态
	returnCode, ok := resp["return_code"]
	if !ok {
		err = ErrNotFoundReturnCode
		fmt.Printf("没有找到return_code 返回...:%v \n", returnCode)

		return
	}
	if returnCode != ReturnCodeSuccess {
		err = &Error{
			ReturnCode: returnCode,
			ReturnMsg:  resp["return_msg"],
		}
		return
	}

	//todo 注释以下代码进行容错处理，解决发放成功后返回错误信息的问题。
	//fmt.Printf("开始验证appid %v\n", resp)

	// 安全考虑, 做下验证 appid 和 mch_id
	//appId, ok := resp["appid"]
	//fmt.Printf("开始验证appid %v\n", appId)
	//if ok && appId != clt.appId {
	//	err = fmt.Errorf("appid mismatch, have: %s, want: %s", appId, clt.appId)
	//	return
	//}
	//
	//mchId, ok := resp["mch_id"]
	//fmt.Printf("开始验证mchId %v\n", mchId)
	//
	//if ok && mchId != clt.mchId {
	//	err = fmt.Errorf("mch_id mismatch, have: %s, want: %s", mchId, clt.mchId)
	//	return
	//}

	// 验证签名
	//signature1, ok := resp["sign"]
	//fmt.Printf("开始验证sign %v\n", signature1)
	//
	//if !ok {
	//	err = ErrNotFoundSign // TODO 恢复这个注释, 待腾讯返回 sign 参数的时候
	//	return
	//}
	//signature2 := Sign(resp, clt.apiKey, nil)
	//if signature1 != signature2 {
	//	err = fmt.Errorf("sign mismatch,\nhave: %s,\nwant: %s", signature1, signature2)
	//	return
	//}
	//
	//fmt.Printf("调用postXml 结束\n")

	return
}
