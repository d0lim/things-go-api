package api

import (
	"time"
)

// StringPointer returns a pointer to a string
func StringPointer(str string) *string {
	return &str
}

// StatusPointer returns a pointer to a TaskStatus
func StatusPointer(val TaskStatus) *TaskStatus {
	return &val
}

// SchedulePointer returns a pointer to a TaskSchedule
func SchedulePointer(val TaskSchedule) *TaskSchedule {
	return &val
}

// TimePointer returns a pointer to a Time
func TimePointer(val time.Time) *Timestamp {
	ts := Timestamp(val)
	return &ts
}
