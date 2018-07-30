package checkup

import (
	"fmt"
)

type Qianbao struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Channel  string `json:"channel"`
	Webhook  string `json:"webhook"`
}

// Notify implements notifier interface
func (q Qianbao) Notify(results []Result) error {
	for _, result := range results {
		//if !result.Healthy {
		q.Send(result)
		//}
	}
	return nil
}

func (q Qianbao) Send(result Result) error {
	fmt.Println("sended to wen", q.Channel, result)
	return nil
}
