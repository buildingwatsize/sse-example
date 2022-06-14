package notisystem

import (
	"bufio"
	"fmt"
)

type Notifier interface {
	IsReady() bool
	Append(msg string)
	Get() chan string
	Clear()
	Send(clients map[string]*bufio.Writer, msg string) map[string]error
}

type notifier struct {
	ready   bool
	channel chan string
}

func NewNotifier() Notifier {
	return &notifier{
		ready:   true,
		channel: make(chan string),
	}
}

func (r *notifier) IsReady() bool {
	return r.ready
}
func (r *notifier) Append(msg string) {
	// fmt.Printf("INCOMING RESPONSE %v\n", msg)
	// 		var buf bytes.Buffer
	//    enc := json.NewEncoder(&buf)
	//    enc.Encode(msg)
	r.ready = true
	r.channel <- msg
}
func (r *notifier) Get() chan string {
	return r.channel
}
func (r *notifier) Clear() {
	r.ready = false
	close(r.channel)
}
func (r *notifier) Send(clients map[string]*bufio.Writer, msg string) map[string]error {
	errList := map[string]error{}
	for id, cli := range clients {
		fmt.Println("Sending to", id)
		fmt.Fprintf(cli, "data: Message: %s\n\n", msg)

		err := cli.Flush()
		if err != nil {
			fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
			errList[id] = err
		}
	}
	return errList
}
