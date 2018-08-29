package tailf

import (
	"fmt"
	"time"

	"github.com/AaronZheng815/loggather/log_agent/confData"
	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

type TailObj struct {
	tail *tail.Tail
	conf confData.CollectConf
}

type TextMsg struct {
	Msg   string
	Topic string
}

type TailMgr struct {
	tails   []*TailObj
	msgChan chan *TextMsg
}

var (
	tailMgr *TailMgr
)

func InitTailf(cconf []confData.CollectConf, chanSize int) (err error) {
	if len(cconf) == 0 {
		err = fmt.Errorf("invalid conf data.")
		return
	}

	tailMgr = &TailMgr{
		msgChan: make(chan *TextMsg, chanSize),
	}

	for _, v := range cconf {
		obj := &TailObj{
			conf: v,
		}

		tails, errt := tail.TailFile(v.LogPath, tail.Config{
			ReOpen:    true,
			Follow:    true,
			Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
			MustExist: false,
			Poll:      true,
		})

		if errt != nil {
			logs.Error("Failed to initialzie tail, err:", err)
			return
		}

		obj.tail = tails

		tailMgr.tails = append(tailMgr.tails, obj)

		go readFromTail(obj)
	}

	return
}

func readFromTail(tailObj *TailObj) {

	for true {
		line, ok := <-tailObj.tail.Lines
		if !ok {
			logs.Warn("tail file close reopen, filename:%s\n", tailObj.tail.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		textmsg := &TextMsg{
			Msg:   line.Text,
			Topic: tailObj.conf.Topic,
		}

		tailMgr.msgChan <- textmsg
	}
}

func GetOneLine() (msg *TextMsg) {
	msg = <-tailMgr.msgChan
	return

}

// func InitTailf() (err error) {

// 	tails, err := tail.TailFile(confData.AppConfig.CConf[0].LogPath, tail.Config{
// 		ReOpen:    true,
// 		Follow:    true,
// 		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
// 		MustExist: false,
// 		Poll:      true,
// 	})

// 	if err != nil {
// 		logs.Error("Failed to initialzie tail, err:", err)
// 		return
// 	}

// 	var msg *tail.Line
// 	var ok bool

// 	for true {
// 		msg, ok = <-tails.Lines
// 		if !ok {
// 			logs.Warn("tail file close reopen, filename:%s\n", tails.Filename)
// 			time.Sleep(100 * time.Millisecond)
// 			continue
// 		}

// 		timeStamp := msg.Time.Format("2006-01-02 15:04:05")
// 		msgStr := fmt.Sprintf("%s %s", timeStamp, msg.Text)
// 		kafkaAgent.Log2kafka(msgStr)
// 	}
// 	return
// }
