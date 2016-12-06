package status

import (
	"gensh.me/VirtualJudge/components/crawler/submitter"
	s "gensh.me/VirtualJudge/components/crawler/status"
)

func init() {
	go webSocketListener()
	s.InitializeStatusTasks(onStatusChanged)
	submitter.InitListener(onSubmittedStatusChangedListener)
}
