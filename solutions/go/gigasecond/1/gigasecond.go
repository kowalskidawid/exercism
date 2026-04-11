package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond adds one gigasecond (1,000,000,000 seconds) to the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
