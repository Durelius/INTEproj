package room

<<<<<<< Updated upstream
type RoomList struct {
	head         *Room
	tail         *Room
	levelCounter int
}

func NewRoomList(r *Room) *RoomList {
	rl := RoomList{head: nil, tail: nil, levelCounter: 0}
	rl.Add(r)
	return &rl
}
func LoadRoomList(r *Room, levelCounter int) *RoomList {
	rl := RoomList{levelCounter: levelCounter}
	if r.next != nil {
		rl.head = r.next
	}
	if r.prev != nil {
		rl.tail = r.prev
	}
	return &rl
}

func (rl *RoomList) Add(r *Room) {
	rl.levelCounter++
	r.level = rl.levelCounter
	if rl.head == nil {
		rl.head = r
		rl.tail = r
		return
	}
	rl.tail.next = r // link old tail forward
	r.prev = rl.tail // link new room back
	rl.tail = r      // update tail
}
func (rl *RoomList) GetHead() *Room {
	return rl.head
}
func (rl *RoomList) GetTail() *Room {
	return rl.tail
}
func (rl *RoomList) GetLevelCounter() int {
	return rl.levelCounter
}
=======
var STARTING_AREA = NewRandomRoom("Starting area", Location{0, 0}, 25, 50, nil, nil)
>>>>>>> Stashed changes
