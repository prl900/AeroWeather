package aeroweather_test

import (
	"testing"
	"time"

	"github.com/prl900/aeroweather"
)

func TestMetarParse(t *testing.T) {
	metars := []string{"METAR KAUS 092135Z 26018G25KT 8000 -TSRA BR SCT045CB BKN060 OVC080 30/21 Q987",
		"YSSY 142300Z 20020KT 9999 FEW028 SCT075 20/12 Q1015 FM2300 MOD TURB BLW 5000FT",
		"YSCB 162200Z 11008KT CAVOK 16/06 Q1021 NOSIG",
		"VEBS 152100Z 18005G12KT 160V200 2000 HZ NSC M14/M13 Q1013 NOSIG"}

	m := &aeroweather.Metar{}

	for _, metar := range metars {
		if m.Parse(metar, time.Now()) != nil {
			t.Errorf("Could not parse: %s", metar)
		}
	}
}
