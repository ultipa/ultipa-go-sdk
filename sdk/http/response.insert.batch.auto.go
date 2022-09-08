/**
 * @Author: zhaohaichao
 * @Description:
 * @File:  response.insert.batch.auto.go
 * @Date: 2022/9/2 5:20 下午
 */

package http

type InsertBatchAutoResponse struct {
	ErrorCode string
	Msg       string
	Statistic *Statistic
	Resps     map[string]*InsertResponse
	ErrorItem map[int]int
}
