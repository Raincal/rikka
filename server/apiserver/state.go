package apiserver

import (
	"net/http"

	"github.com/Raincal/rikka/common/util"
)

// stateHandleFunc is the base handle func of path /api/state/taskID
func stateHandleFunc(w http.ResponseWriter, r *http.Request) {

	taskID := util.GetTaskIDByRequest(r)

	l.Info("Receive a state request of task", taskID, "from ip", util.GetClientIP(r))

	var jsonData []byte
	var err error
	if jsonData, err = getStateJson(taskID); err != nil {
		l.Warn("Error happened when get state json of task", taskID, ":", err)
	} else {
		l.Debug("Get state json", string(jsonData), "of task", taskID, "successfully")
	}

	renderJsonOrError(w, taskID, jsonData, err, http.StatusInternalServerError)
}
