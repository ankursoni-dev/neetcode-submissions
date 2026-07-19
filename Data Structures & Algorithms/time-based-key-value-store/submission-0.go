type TimeMap struct {
	entries map[string][]Detail
}

type Detail struct {
	timestamp int
	value string
}

func Constructor() TimeMap {
	return TimeMap {
		entries : make(map[string][]Detail),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.entries[key] = append(this.entries[key],  Detail{timestamp :timestamp, value : value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _,ok := this.entries[key];!ok {
		return ""
	}

	details := this.entries[key]
	low,high:=0,len(details)-1

	for low <= high {
		mid := low + (high-low)/2
		if details[mid].timestamp <= timestamp {
			if mid == len(details) - 1 ||  details[mid+1].timestamp > timestamp {
				return details[mid].value
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ""
}
