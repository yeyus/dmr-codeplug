package tytera

import (
	"encoding/json"
	"fmt"
	"github.com/yeyus/dmr-codeplug/proto/tytera"
	"testing"
)

func TestEmergencySystemsParsing(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	es := GetEmergencySystemsGroup()

	b, err := json.MarshalIndent(es.Decode(content[:], 0x125), "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(b))
	fmt.Printf("%+v", es.Systems)
}

func TestEmergencySystemsProto(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	es := GetEmergencySystemsGroup()

	es.Decode(content[:], 0x125)

	if es.Systems.RadioDisableDecode != true {
		t.Errorf("expected radio disable decode to be true, got %t", es.Systems.RadioDisableDecode)
	}

	if es.Systems.RemoteMonitorDecode != false {
		t.Errorf("expected remote monitor decode to be false, got %t", es.Systems.RemoteMonitorDecode)
	}

	if es.Systems.EmergencyRemoteMonitorDecode != true {
		t.Errorf("expected emergency remote monitor decode to be true, got %t", es.Systems.EmergencyRemoteMonitorDecode)
	}

	if es.Systems.RemoteMonitorDuration != 10 {
		t.Errorf("expected remote monitor duration to be 10, got %d", es.Systems.RemoteMonitorDuration)
	}

	if es.Systems.TxSyncWakeup != 150 {
		t.Errorf("expected tx sync wakeup to be 10, got %d", es.Systems.TxSyncWakeup)
	}

	if es.Systems.TxWakeupMessageLimit != 3 {
		t.Errorf("expected tx wakeup message limit to be 3, got %d", es.Systems.TxWakeupMessageLimit)
	}

	// testing entries
	if len(es.Systems.Entries) != 32 {
		t.Errorf("expected number of emergency system entries to be 32, got %d", len(es.Systems.Entries))
	}
}

func TestEmergencyEntries(t *testing.T) {
	content := getRDTBytes("../../packing/tytera/testdata/usa_codeplug.rdt")

	es := GetEmergencySystemsGroup()

	es.Decode(content[:], 0x125)

	testEntry(t, es.Systems.Entries[0], "System1", tytera.AlarmType_REGULAR_TYPE, tytera.AlarmMode_EMERGENCY_ALARM, 15, 5, 10, 0)
	testEntry(t, es.Systems.Entries[1], "Call DMR Call", tytera.AlarmType_REGULAR_TYPE, tytera.AlarmMode_EMERGENCY_ALARM, 15, 5, 10, 1)
	testEntry(t, es.Systems.Entries[2], "Call Hotspot", tytera.AlarmType_SILENT_TYPE, tytera.AlarmMode_EMERGENCY_ALARM_WITH_CALL, 10, 10, 100, 2)
	testEntry(t, es.Systems.Entries[3], "Call WorldWide", tytera.AlarmType_SILENT_WITH_VOICE_TYPE, tytera.AlarmMode_EMERGENCY_ALARM_WITH_VOICE_TO_FOLLOW, 15, 5, 10, 3)

}

func testEntry(t *testing.T, entry *tytera.EmergencySystemEntry, systemName string, alarmType tytera.AlarmType, alarmMode tytera.AlarmMode, impolite uint32, polite uint32, hotMic uint32, revertChannel uint32) {
	if entry.SystemName != systemName {
		t.Errorf("expected system name to be %s, but got %s", systemName, entry.SystemName)
	}

	if entry.AlarmType != alarmType {
		t.Errorf("expected alarm type to be %v, but got %v", alarmType, entry.AlarmType)
	}

	if entry.AlarmMode != alarmMode {
		t.Errorf("expected alarm mode to be %v, but got %v", alarmMode, entry.AlarmMode)
	}

	if entry.ImpoliteRetries != impolite {
		t.Errorf("expected impolite retries to be %d, but got %d", impolite, entry.ImpoliteRetries)
	}

	if entry.PoliteRetries != polite {
		t.Errorf("expected polite retries to be %d, but got %d", polite, entry.PoliteRetries)
	}

	if entry.HotMicDuration != hotMic {
		t.Errorf("expected hot mic duration to be %d, but got %d", hotMic, entry.HotMicDuration)
	}

	if entry.RevertChannel != revertChannel {
		t.Errorf("expected revert channel to be %d, but got %d", revertChannel, entry.RevertChannel)
	}
}
