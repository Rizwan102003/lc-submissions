package main
import "strconv"

const PACKETS_NOT_FOUND_FOR_GIVEN_TIMESTAMP_RANGE = 0
var THERE_ARE_NO_PACKETS = make([]int, 0)

type Packet struct {
    source      int
    destination int
    timestamp   int
}

func NewPacket(source int, destination int, timestamp int) Packet {
    packet := Packet{
        source:      source,
        destination: destination,
        timestamp:   timestamp,
    }
    return packet
}

type Timestamps struct {
    startIndex int
    timestamps []int
}

func NewTimestamps() *Timestamps {
    stamps := &Timestamps{
        startIndex: 0,
        timestamps: make([]int, 0),
    }
    return stamps
}

type Router struct {
    memoryLimit             int
    packets                 []Packet
    quickAccessPackets      HashSet
    destinationToTimestamps map[int]*Timestamps
}

func Constructor(memoryLimit int) Router {
    router := Router{
        memoryLimit:             memoryLimit,
        packets:                 make([]Packet, 0),
        quickAccessPackets:      NewHashSet(),
        destinationToTimestamps: map[int](*Timestamps){},
    }
    return router
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
    packet := NewPacket(source, destination, timestamp)
    packetAsString := convertPacketToString(packet)

    if this.quickAccessPackets.Contains(packetAsString) {
        return false
    }

    if len(this.packets) == this.memoryLimit {
        toRemove := this.packets[0]
        toRemoveAsString := convertPacketToString(toRemove)

        this.packets = this.packets[1:]
        this.quickAccessPackets.Remove(toRemoveAsString)
        this.destinationToTimestamps[toRemove.destination].startIndex++
    }

    this.packets = append(this.packets, packet)
    this.quickAccessPackets.Add(packetAsString)

    if _, contains := this.destinationToTimestamps[destination]; !contains {
        this.destinationToTimestamps[destination] = NewTimestamps()
    }
    this.destinationToTimestamps[destination].timestamps = append(this.destinationToTimestamps[destination].timestamps, timestamp)

    return true
}

func (this *Router) ForwardPacket() []int {
    if len(this.packets) == 0 {
        return THERE_ARE_NO_PACKETS
    }

    toRemove := this.packets[0]
    toRemoveAsString := convertPacketToString(toRemove)

    this.packets = this.packets[1:]
    this.quickAccessPackets.Remove(toRemoveAsString)
    this.destinationToTimestamps[toRemove.destination].startIndex++

    return []int{toRemove.source, toRemove.destination, toRemove.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
    if _, contains := this.destinationToTimestamps[destination]; !contains {
        return PACKETS_NOT_FOUND_FOR_GIVEN_TIMESTAMP_RANGE
    }

    timestamps := this.destinationToTimestamps[destination].timestamps
    var startIndex = this.destinationToTimestamps[destination].startIndex
    var endIndex = len(this.destinationToTimestamps[destination].timestamps) - 1

    if startIndex > endIndex ||
        startIndex == len(timestamps) ||
        timestamps[startIndex] > endTime ||
        timestamps[endIndex] < startTime {
        return PACKETS_NOT_FOUND_FOR_GIVEN_TIMESTAMP_RANGE
    }
    for startIndex < endIndex && timestamps[startIndex] < startTime {
        startIndex++
    }
    for endIndex > 0 && timestamps[endIndex] > endTime {
        endIndex--
    }
    return endIndex - startIndex + 1
}

func convertPacketToString(packet Packet) string {
    return strconv.Itoa(packet.source) + "," + strconv.Itoa(packet.destination) + "," + strconv.Itoa(packet.timestamp)
}

type HashSet struct {
    conainer map[string]bool
}

func NewHashSet() HashSet {
    return HashSet{conainer: map[string]bool{}}
}

func (this HashSet) Contains(packet string) bool {
    return this.conainer[packet]
}

func (this HashSet) Add(packet string) {
    this.conainer[packet] = true
}

func (this HashSet) Remove(packet string) {
    delete(this.conainer, packet)
}


/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */