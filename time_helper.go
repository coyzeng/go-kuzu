package kuzu

// #include "kuzu.h"
// #include <stdlib.h>
import "C"
import "time"

func unixEpoch() time.Time {
	return time.Unix(0, 0)
}

func timeToKuzuDate(inputTime time.Time) C.kuzu_date_t {
	diff := inputTime.Sub(unixEpoch())
	diffDays := diff.Hours() / 24
	cKuzuDate := C.kuzu_date_t{}
	cKuzuDate.days = C.int32_t(diffDays)
	return cKuzuDate
}

func kuzuDateToTime(cKuzuDate C.kuzu_date_t) time.Time {
	diff := time.Duration(cKuzuDate.days) * 24 * time.Hour
	return unixEpoch().Add(diff)
}
