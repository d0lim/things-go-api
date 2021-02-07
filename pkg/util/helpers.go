package util

import (
	"time"

	"github.com/d0lim/things-go-api/pkg/common"
)

// StringPointer returns a pointer to a string
func StringPointer(str string) *string {
	return &str
}

// StatusPointer returns a pointer to a TaskStatus
func StatusPointer(val common.TaskStatus) *common.TaskStatus {
	return &val
}

// SchedulePointer returns a pointer to a TaskSchedule
func SchedulePointer(val common.TaskSchedule) *common.TaskSchedule {
	return &val
}

// TimePointer returns a pointer to a Time
func TimePointer(val time.Time) *common.Timestamp {
	ts := common.Timestamp(val)
	return &ts
}
