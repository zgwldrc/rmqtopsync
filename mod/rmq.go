package mod

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

type RocketMQConsole struct {
	URL *url.URL
	cli *http.Client
}
func NewRocketMQConsole(_url string) *RocketMQConsole {
	cli := &http.Client{
		Timeout: 5 * time.Second,
	}
	_u, err := url.Parse(_url)
	cobra.CheckErr(err)

	return &RocketMQConsole{
		URL: _u,
		cli: cli,
	}
}

func (rmq *RocketMQConsole) GetTopics () []string {

	_url := rmq.URL.JoinPath("/topic/list.query").String()
	res, err := rmq.cli.Get(_url)
	cobra.CheckErr(err)
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	cobra.CheckErr(err)
	getTopicsResp := &GetTopicsResp{}
	cobra.CheckErr(json.Unmarshal(resBody, getTopicsResp))

	return getTopicsResp.Data.TopicList
}

type GetTopicsResp struct {
	Status int `json:"status"`
	Data   struct {
		TopicList  []string    `json:"topicList"`
		BrokerAddr interface{} `json:"brokerAddr"`
	} `json:"data"`
	ErrMsg interface{} `json:"errMsg"`
}


func (rmq *RocketMQConsole) GetClusters () *GetClustersResp {

	_url := rmq.URL.JoinPath("/cluster/list.query").String()
	res, err := rmq.cli.Get(_url)
	cobra.CheckErr(err)
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	cobra.CheckErr(err)
	getClustersResp := &GetClustersResp{}
	cobra.CheckErr(json.Unmarshal(resBody, getClustersResp))

	return getClustersResp
}

type GetClustersResp struct {
	Status int `json:"status"`
	Data   struct {
		BrokerServer  map[string]Cluster               `json:"brokerServer"`
		ClusterInfo struct {
			BrokerAddrTable map[string]BrokerAddrTableItem `json:"brokerAddrTable"`
			ClusterAddrTable map[string][]string           `json:"clusterAddrTable"`
		} `json:"clusterInfo"`
	} `json:"data"`
	ErrMsg interface{} `json:"errMsg"`
}
type Cluster map[int]Broker
type Broker struct {
	MsgPutTotalTodayNow                             string `json:"msgPutTotalTodayNow"`
	CommitLogDiskRatio                              string `json:"commitLogDiskRatio"`
	GetFoundTps                                     string `json:"getFoundTps"`
	EndTransactionThreadPoolQueueCapacity           string `json:"EndTransactionThreadPoolQueueCapacity"`
	SendThreadPoolQueueHeadWaitTimeMills            string `json:"sendThreadPoolQueueHeadWaitTimeMills"`
	PutMessageDistributeTime                        string `json:"putMessageDistributeTime"`
	QueryThreadPoolQueueHeadWaitTimeMills           string `json:"queryThreadPoolQueueHeadWaitTimeMills"`
	RemainHowManyDataToFlush                        string `json:"remainHowManyDataToFlush"`
	MsgGetTotalTodayNow                             string `json:"msgGetTotalTodayNow"`
	PutMessageFailedTimes                           string `json:"putMessageFailedTimes"`
	CommitLogMaxOffset                              string `json:"commitLogMaxOffset"`
	QueryThreadPoolQueueSize                        string `json:"queryThreadPoolQueueSize"`
	GetMessageEntireTimeMax                         string `json:"getMessageEntireTimeMax"`
	MsgPutTotalTodayMorning                         string `json:"msgPutTotalTodayMorning"`
	PutMessageTimesTotal                            string `json:"putMessageTimesTotal"`
	BootTimestamp                                   string `json:"bootTimestamp"`
	MsgGetTotalTodayMorning                         string `json:"msgGetTotalTodayMorning"`
	MsgPutTotalYesterdayMorning                     string `json:"msgPutTotalYesterdayMorning"`
	MsgGetTotalYesterdayMorning                     string `json:"msgGetTotalYesterdayMorning"`
	PullThreadPoolQueueSize                         string `json:"pullThreadPoolQueueSize"`
	BrokerVersionDesc                               string `json:"brokerVersionDesc"`
	SendThreadPoolQueueSize                         string `json:"sendThreadPoolQueueSize"`
	StartAcceptSendRequestTimeStamp                 string `json:"startAcceptSendRequestTimeStamp"`
	CommitLogMinOffset                              string `json:"commitLogMinOffset"`
	PutMessageEntireTimeMax                         string `json:"putMessageEntireTimeMax"`
	PullThreadPoolQueueHeadWaitTimeMills            string `json:"pullThreadPoolQueueHeadWaitTimeMills"`
	Runtime                                         string `json:"runtime"`
	EarliestMessageTimeStamp                        string `json:"earliestMessageTimeStamp"`
	CommitLogDirCapacity                            string `json:"commitLogDirCapacity"`
	DispatchMaxBuffer                               string `json:"dispatchMaxBuffer"`
	BrokerVersion                                   string `json:"brokerVersion"`
	PutTps                                          string `json:"putTps"`
	RemainTransientStoreBufferNumbs                 string `json:"remainTransientStoreBufferNumbs"`
	GetMissTps                                      string `json:"getMissTps"`
	QueryThreadPoolQueueCapacity                    string `json:"queryThreadPoolQueueCapacity"`
	CommitLogDiskRatioDataRocketmq493StoreCommitlog string `json:"commitLogDiskRatio_/data/rocketmq/4.9.3/store/commitlog"`
	PutMessageAverageSize                           string `json:"putMessageAverageSize"`
	GetTransferedTps                                string `json:"getTransferedTps"`
	DispatchBehindBytes                             string `json:"dispatchBehindBytes"`
	PutMessageSizeTotal                             string `json:"putMessageSizeTotal"`
	SendThreadPoolQueueCapacity                     string `json:"sendThreadPoolQueueCapacity"`
	EndTransactionQueueSize                         string `json:"EndTransactionQueueSize"`
	GetTotalTps                                     string `json:"getTotalTps"`
	PullThreadPoolQueueCapacity                     string `json:"pullThreadPoolQueueCapacity"`
	ConsumeQueueDiskRatio                           string `json:"consumeQueueDiskRatio"`
	PageCacheLockTimeMills                          string `json:"pageCacheLockTimeMills"`
}

type BrokerAddrs map[int]string
type BrokerAddrTableItem struct {
	Cluster     string `json:"cluster"`
	BrokerName  string `json:"brokerName"`
	BrokerAddrs BrokerAddrs `json:"brokerAddrs"`
}

func (rmq *RocketMQConsole) CreateOrUpdate (req *CreateOrUpdateReq) *CreateOrUpdateResp {

	_url := rmq.URL.JoinPath("/topic/createOrUpdate.do").String()

	data, err := json.Marshal(req)
	cobra.CheckErr(err)
	res, err := rmq.cli.Post(_url, "application/json", bytes.NewReader(data))
	cobra.CheckErr(err)
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	cobra.CheckErr(err)
	createOrUpdateResp := &CreateOrUpdateResp{}
	cobra.CheckErr(json.Unmarshal(resBody, createOrUpdateResp))

	return createOrUpdateResp
}


type CreateOrUpdateReq struct {
	WriteQueueNums  int      `json:"writeQueueNums"`
	ReadQueueNums   int      `json:"readQueueNums"`
	Perm            int      `json:"perm"`
	Order           bool     `json:"order"`
	TopicName       string   `json:"topicName"`
	BrokerNameList  []string `json:"brokerNameList"`
	ClusterNameList []string `json:"clusterNameList"`
}

type CreateOrUpdateResp struct {
	Status int         `json:"status"`
	Data   bool        `json:"data"`
	ErrMsg interface{} `json:"errMsg"`
}