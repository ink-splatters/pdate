package main

import (
	"errors"
	"log"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)



func parseDateparse(val string)(time.Time, error) {
	
	var (
		err error
		result time.Time
	)

	if result, err = dateparse.ParseStrict(val); err == nil {
		return result, err
	}

	valUpper := strings.ToUpper(val)
	if result, err = dateparse.ParseStrict(valUpper); err != nil {
		return result, err
	}

	if err = result.UnmarshalText([]byte(val)); err == nil {
		return result, err
	}

	if err = result.UnmarshalText([]byte(valUpper)); err == nil {
		return result, err
	}

	return result, err
}

// Parse attempts to parse the given string into a time.Time.
func Parse(val string) (time.Time, error) {
	if unixTs, err := strconv.ParseInt(val, 10, 64); err == nil {
		result := time.Unix(unixTs, 0)
		if result.Year() > 2070 {
			result = time.Unix(0, unixTs*int64(time.Millisecond))
		}
		if result.Year() < 1970 {
			result = time.Unix(0, unixTs)
		}
		return result, nil
	}

	if _, err := strconv.ParseFloat(val, 64); err == nil {
		split := strings.Split(val, ".")
		if lsplit:=len(split); lsplit != 2 {
			// should never happen, as we know this is parseable as a float
			log.Fatalf("[FATAL]: len(split) != 2: %d",lsplit)
		}
		secondsStr := split[0]
		subsecondsStr := split[1]
		if len(subsecondsStr) > 9 {
			subsecondsStr = subsecondsStr[:9]
		} else if len(subsecondsStr) < 9 {
			subsecondsStr = fmt.Sprintf("%s%s", subsecondsStr, strings.Repeat("0", 9-len(subsecondsStr)))
		}
		seconds, err := strconv.ParseInt(secondsStr, 10, 64)
		if err != nil {
			// should never happen, as we know this is parseable as a float
			panic(err)
		}
		nanos, err := strconv.ParseInt(subsecondsStr, 10, 64)
		if err != nil {
			// should never happen, as we know this is parseable as a float
			panic(err)
		}

		result := time.Unix(seconds, nanos)
		return result, nil
	}

	for _, perm := range Permutations(strings.Split(val," ")) {

		if result,err := parseDateparse(strings.Join(perm," ")); err == nil {
			return result,err
		}
	}

	return time.Time{}, errors.New("failed to parse date string")
}
