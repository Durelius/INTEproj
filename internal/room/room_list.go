package room

type RoomList struct {
	head *Room
	tail *Room
}

var Rooms = &RoomList{}

func (rl *RoomList) Add(r *Room) {
	if rl.head == nil {
		rl.head = r
		rl.tail = r
		return
	}
	rl.tail.next = r // link old tail forward
	r.prev = rl.tail // link new room back
	rl.tail = r      // update tail
}
