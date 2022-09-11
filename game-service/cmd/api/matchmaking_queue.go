package main

import (
	"sync"
)

type matchQueueItem struct {
	client *matchClient
	low    float64
	high   float64
}

type matchQueue struct {
	m    sync.Mutex
	c    sync.Cond
	data []matchQueueItem
}

func newMatchQueue() *matchQueue {
	q := new(matchQueue)
	q.c = sync.Cond{L: &q.m}
	return q
}

func (q *matchQueue) addToQ(item *matchQueueItem) {
	q.c.L.Lock()
	defer q.c.L.Unlock()

	q.data = append(q.data, *item)
	q.c.Signal()
}

func (q *matchQueue) removeClientFromQ(item *matchQueueItem) {
	var idx int
	for i := 0; i < len(q.data); i++ {
		if q.data[i].client.userId == item.client.userId {
			idx = i
		}
	}

	if idx == len(q.data)-1 {
		q.data = q.data[:idx]
	} else {
		q.data = append(q.data[:idx], q.data[idx+1:]...)
	}
}

func (q *matchQueue) removeMultipleClientsFromQ(items []*matchQueueItem) {
	for i := 0; i < len(items); i++ {
		q.removeClientFromQ(items[i])
	}
}

func isMatch(a *matchQueueItem, b *matchQueueItem) bool {
	return (a.low >= b.low && a.low <= b.high) || (a.high >= b.low && a.high <= b.low)
}

func (q *matchQueue) findMatch(user *matchQueueItem) []*matchQueueItem {
	var res []*matchQueueItem
	var count int
	for i := 0; i < len(q.data); i++ {
		if isMatch(&q.data[i], user) && q.data[i].client.userId != user.client.userId {
			res = append(res, &q.data[i])
			count++
		}
	}

	if count < 2 || count > 4 {
		return nil
	} else {
		return append(res, user)
	}
}

func (q *matchQueue) findQueueItem(c *matchClient) int {
	for i := 0; i < len(q.data); i++ {
		if q.data[i].client.userId == c.userId {
			return i
		}
	}

	return -1
}
