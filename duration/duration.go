package duration

import (
	"encoding/json"
	"fmt"
	"github.com/xhit/go-str2duration/v2"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	dur, err := str2duration.ParseDuration(value)
	if err != nil {
		return err
	}
	*d = Duration(dur)
	return nil
}
func (d Duration) MarshalJSON() ([]byte, error) {
	dur := time.Duration(d)
	return []byte(fmt.Sprintf(`"%s"`, str2duration.String(dur))), nil
}
