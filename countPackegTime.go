package main

import "time"

type Event struct {
	count int
	time  time.Time
}

// Подход с плохой производительностью
type MinuteHourCounter struct {
	listEvent []Event
	__version string
}

func (mhc *MinuteHourCounter) Add(count int) {
	mhc.listEvent = append(mhc.listEvent, Event{count, time.Now()})
}

func (mhc *MinuteHourCounter) MinuteCount() int {
	return mhc.CountSince(time.Now().UnixNano()/int64(time.Second) - 60)
}

func (mhc *MinuteHourCounter) HourCount() int {
	return mhc.CountSince(time.Now().UnixNano()/int64(time.Second) - 3600)
}

func (mhc *MinuteHourCounter) CountSince(time_in_second int64) int {
	var count int = 0

	for i := len(mhc.listEvent) - 1; i >= 0; i-- {
		if mhc.listEvent[i].time.UnixNano()/int64(time.Second) > time_in_second {
			count += mhc.listEvent[i].count
		}
	}

	return count
}

func NewMinuteHourCounter() *MinuteHourCounter {
	return &MinuteHourCounter{__version: "1.0.1"}
}

// Реализация двухуровневого дизайна конвейера
type MinuteHourCounterConveer struct {
	minute_events []Event
	hour_events   []Event

	minute_count int
	hour_count   int

	__version string
}

func (mhcc *MinuteHourCounterConveer) Add(count int) {
	var now_secs time.Time = time.Now()
	mhcc.minute_events = append(mhcc.minute_events, Event{count, now_secs})

	mhcc.ShiftOldEvents(now_secs)

	mhcc.minute_count += 1
	mhcc.hour_count += 1
}

func (mhcc *MinuteHourCounterConveer) MinuteCount() int {
	mhcc.ShiftOldEvents(time.Now())
	return mhcc.minute_count
}

func (mhcc *MinuteHourCounterConveer) HourCount() int {
	mhcc.ShiftOldEvents(time.Now())
	return mhcc.hour_count
}

func (mhcc *MinuteHourCounterConveer) ShiftOldEvents(now_secs time.Time) {
	var minute_ago int64 = now_secs.UnixNano()/int64(time.Second) - 60
	var hour_ago int64 = now_secs.UnixNano()/int64(time.Second) - 3600

	for count_minute := 0; count_minute < len(mhcc.minute_events) &&
		mhcc.minute_events[count_minute].time.UnixNano()/int64(time.Second) <= minute_ago; count_minute++ {
		mhcc.hour_events = append(mhcc.hour_events, mhcc.minute_events[count_minute])
		mhcc.minute_events = append(mhcc.minute_events[:count_minute], mhcc.minute_events[count_minute+1:]...)
		mhcc.minute_count = len(mhcc.minute_events)
	}

	for count_hours := 0; count_hours < len(mhcc.hour_events) &&
		mhcc.hour_events[count_hours].time.UnixNano()/int64(time.Second) <= hour_ago; count_hours++ {
		mhcc.hour_events = append(mhcc.hour_events[:count_hours], mhcc.hour_events[count_hours+1:]...)
		mhcc.hour_count = len(mhcc.hour_events)
	}
}

func NewMinuteHourCounterConveer() *MinuteHourCounterConveer {
	return &MinuteHourCounterConveer{__version: "1.1.0"}
}

//Реализация с временными блоками
func Constructor() *MinuteHourCounterInsert {

	return &MinuteHourCounterInsert{minute_counts: TrailingBucketCounter{buckets: ConveyorQueue{max_items: 60}, secs_per_bucket: 1},
		hours_counts: TrailingBucketCounter{buckets: ConveyorQueue{max_items: 60}, secs_per_bucket: 60}, __version: "2.0.0"}
}

type MinuteHourCounterInsert struct {
	minute_counts TrailingBucketCounter
	hours_counts  TrailingBucketCounter
	__version     string
}

func (mhci *MinuteHourCounterInsert) Add(count int) {
	var now_time time.Time = time.Now()
	mhci.minute_counts.Add(count, now_time)
	mhci.hours_counts.Add(count, now_time)
}

func (mhci *MinuteHourCounterInsert) MinuteCount() int {
	var now_time time.Time = time.Now()
	return mhci.minute_counts.TrailingCount(now_time)
}

func (mhci *MinuteHourCounterInsert) HourCount() int {
	var now_time time.Time = time.Now()
	return mhci.hours_counts.TrailingCount(now_time)
}

//ConveyorQueue
type ConveyorQueue struct {
	queue     []int
	max_items int
	total_sum int
}

func (conv *ConveyorQueue) Shift(num_shifted int) {
	if num_shifted >= conv.max_items {
		conv.queue = make([]int, 10)
		conv.total_sum = 0
		return
	}
	for num_shifted > 0 {
		conv.queue = append(conv.queue, 0)
		num_shifted--
	}

	for len(conv.queue) > conv.max_items {
		var index int = len(conv.queue) - 1
		conv.total_sum -= conv.queue[index]
		conv.queue = append(conv.queue[:index], conv.queue[index+1:]...)
	}
}

func (conv *ConveyorQueue) AddToBack(count int) {
	if len(conv.queue) > 0 {
		conv.Shift(1)
	}
	conv.queue[len(conv.queue)-1] += 1
	conv.total_sum += count
}

//TrailingBucketCounter
type TrailingBucketCounter struct {
	buckets          ConveyorQueue
	secs_per_bucket  int64
	last_update_time time.Time
}

func (tbc *TrailingBucketCounter) Update(now time.Time) {
	var current_backet int = int((now.UnixNano() / int64(time.Second)) / tbc.secs_per_bucket)
	var last_update_bucket int = int((tbc.last_update_time.UnixNano() / int64(time.Second)) / tbc.secs_per_bucket)
	tbc.buckets.Shift(current_backet - last_update_bucket)
	tbc.last_update_time = now
}

func (tbc *TrailingBucketCounter) Add(count int, now time.Time) {
	tbc.Update(now)
	tbc.buckets.AddToBack(count)
}

func (tbc *TrailingBucketCounter) TrailingCount(now time.Time) int {
	tbc.Update(now)
	return tbc.buckets.total_sum
}
