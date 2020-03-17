package object

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var (
	testSeconds  int64 = 120
	testDuration       = int64(time.Second) * testSeconds
)

func TestDurationInputMarshal(t *testing.T) {
	type testStruct struct {
		Value *DurationInput
	}

	di := DurationInput(testDuration)

	t1 := testStruct{Value: &di}
	res, err := json.Marshal(t1)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%s\n", res)
}

func TestDurationInputNanoSeconds(t *testing.T) {
	// Nano seconds to DurationInput
	var ni DurationInput
	ni = DurationInput(testDuration)
	fmt.Printf("Duration ni in seconds is %d\n", ni.GetSeconds())
	if testSeconds != ni.GetSeconds() {
		t.Errorf("expected ni %d got %d", testSeconds, ni.GetSeconds())
	}
	fmt.Printf("Duration ni as a duration is %d\n", ni.GetDuration())
	if testDuration != int64(ni.GetDuration()) {
		t.Errorf("expected ni %d got %d", testDuration, ni.GetDuration())
	}
	fmt.Printf("Duration ni raw value is %d\n", ni.GetRawValue())
	if testDuration != int64(ni.GetRawValue()) {
		t.Errorf("expected ni %d got %d", testDuration, ni.GetRawValue())
	}
}

func TestDurationInputSeconds(t *testing.T) {
	// Seconds to DurationInput
	var ti DurationInput
	ti = DurationInput(testSeconds)
	fmt.Printf("Duration ti in seconds is %d\n", ti.GetSeconds())
	if testSeconds != ti.GetSeconds() {
		t.Errorf("expected ti %d got %d", testSeconds, ti.GetSeconds())
	}
	fmt.Printf("Duration ti as a duration is %d\n", ti.GetDuration())
	if testDuration != int64(ti.GetDuration()) {
		t.Errorf("expected ni %d got %d", testDuration, ti.GetDuration())
	}
	fmt.Printf("Duration ti raw value is %d\n", ti.GetRawValue())
	if testSeconds != int64(ti.GetRawValue()) {
		t.Errorf("expected ni %d got %d", testDuration, ti.GetRawValue())
	}
}
