package core

// History is the struct for the history
type History struct {
	records  []string
	size     int // current size of the history
	capacity int // max size of the history
	head     int // index of the first record
}

// NewHistory is the constructor for the History
func NewHistory(capacity int) *History {
	return &History{
		records:  make([]string, capacity),
		size:     0,
		capacity: capacity,
		head:     0,
	}
}

// Add is the method to add a record to the history
func (h *History) Add(record string) {
	// calculate the position of the new record
	// if the history is full, the new record will replace the oldest one
	pos := (h.head + h.size) % h.capacity
	if h.size < h.capacity {
		h.size++
	}
	// add the new record to the history
	h.records[pos] = record
}

// GetRecords is the method to get the records from the history
func (h *History) GetRecords() []string {
	result := make([]string, h.size)
	for i := range h.size {
		pos := (h.head + i) % h.capacity
		result[i] = h.records[pos]
	}
	return result
}
