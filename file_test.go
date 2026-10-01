package sie_test

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/omniboost/go-sie"
)

// TestMarshalExample rebuilds the OPB-870 example file
func TestMarshalExample(t *testing.T) {
	date := sie.Date{time.Date(2026, 8, 28, 0, 0, 0, 0, time.UTC)}

	f := sie.File{
		Header: sie.Header{
			Flagga:  0,
			Program: "OMNIBOOST",
			Format:  "PC8",
			Gen:     date,
			SieTyp:  4,
			Prosa:   "Bokföringsorder",
			FNamn:   "Omniboost Hotel",
			KPTyp:   "EUBAS97",
		},
		Accounts: sie.Accounts{
			{Number: 1512, Name: "Faktura"},
			{Number: 1513, Name: "Master Card & VISA"},
			{Number: 1514, Name: "Amex"},
			{Number: 1575, Name: "Boende Gäster"},
			{Number: 1580, Name: "Förskott Nets"},
			{Number: 1680, Name: "Bankgiro"},
			{Number: 2420, Name: "Förskott"},
			{Number: 2610, Name: "25% moms"},
			{Number: 2620, Name: "12% moms"},
			{Number: 2630, Name: "6% moms"},
			{Number: 2820, Name: "Dricks"},
			{Number: 2850, Name: "Lunchhäfte"},
			{Number: 2860, Name: "Presentkort"},
			{Number: 3010, Name: "Logi"},
			{Number: 3010, Name: "Logi observation"},
			{Number: 3120, Name: "Mat 25"},
			{Number: 3121, Name: "Mat 12"},
			{Number: 3121, Name: "Frukost"},
			{Number: 3122, Name: "MAT 6%"},
			{Number: 3210, Name: "Kaffe 25"},
			{Number: 3211, Name: "Kaffe 12"},
			{Number: 3221, Name: "Vatten 12"},
			{Number: 32210, Name: "SPA Vatten 12%"},
			{Number: 3240, Name: "Starköl"},
			{Number: 32400, Name: "SPA Starköl"},
			{Number: 324000, Name: "Hotell starköl"},
			{Number: 3250, Name: "Vin"},
			{Number: 32500, Name: "SPA Vin"},
			{Number: 325000, Name: "Hotell Vin"},
			{Number: 3260, Name: "Sprit"},
			{Number: 32600, Name: "SPA sprit"},
			{Number: 354100, Name: "Hotell Konfektyr 6%"},
			{Number: 3610, Name: "Hyra lokal"},
			{Number: 3614, Name: "Hyra festlokal"},
			{Number: 3616, Name: "Spa entré"},
			{Number: 3630, Name: "Gym entré"},
			{Number: 3650, Name: "Garage"},
			{Number: 3660, Name: "SPA Produkter"},
			{Number: 6050, Name: "Försprov lunchhäfte"},
		},
		Dimensions: sie.Dimensions{
			{Number: 1, Name: "Resultatenhet"},
		},
		Objects: sie.Objects{
			{Dimension: 1, Number: "101", Name: "Restaurang /"},
			{Dimension: 1, Number: "200", Name: "SPA /"},
			{Dimension: 1, Number: "30", Name: "Konferens"},
			{Dimension: 1, Number: "307", Name: "Hotell /Reception"},
			{Dimension: 1, Number: "600", Name: "Gym"},
			{Dimension: 1, Number: "80", Name: "Betalningar"},
		},
		Vouchers: sie.Vouchers{
			{
				Series: "",
				Number: "1",
				Date:   date,
				Text:   "Kassarapport",
				Transactions: sie.Transactions{
					{Account: 1512, Amount: sie.RequireAmountFromString("42110.61")},
					{Account: 1513, Amount: sie.RequireAmountFromString("150451.71")},
					{Account: 1514, Amount: sie.RequireAmountFromString("11879.85")},
					{Account: 1575, Amount: sie.RequireAmountFromString("60751.23")},
					{Account: 1580, Amount: sie.RequireAmountFromString("3111.12")},
					{Account: 1680, Amount: sie.RequireAmountFromString("17720.16")},
					{Account: 2420, Amount: sie.RequireAmountFromString("13580.70")},
					{Account: 2610, Amount: sie.RequireAmountFromString("-19279.45")},
					{Account: 2620, Amount: sie.RequireAmountFromString("-20288.42")},
					{Account: 2630, Amount: sie.RequireAmountFromString("-693.62")},
					{Account: 2820, Objects: sie.TransactionObjects{{Dimension: 1, Number: "80"}}, Amount: sie.RequireAmountFromString("-661.20")},
					{Account: 2850, Amount: sie.RequireAmountFromString("-1456.38")},
					{Account: 2860, Amount: sie.RequireAmountFromString("522.00")},
					{Account: 3010, Objects: sie.TransactionObjects{{Dimension: 1, Number: "307"}}, Amount: sie.RequireAmountFromString("-98588.87")},
					{Account: 3120, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("-11492.29")},
					{Account: 3121, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("-57126.38")},
					{Account: 3122, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("-229.81")},
					{Account: 3210, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("-2199.42")},
					{Account: 3211, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("-3586.42")},
					{Account: 3221, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("-1901.23")},
					{Account: 32210, Objects: sie.TransactionObjects{{Dimension: 1, Number: "200"}}, Amount: sie.RequireAmountFromString("-195.75")},
					{Account: 3240, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("-3832.18")},
					{Account: 32400, Objects: sie.TransactionObjects{{Dimension: 1, Number: "200"}}, Amount: sie.RequireAmountFromString("-1117.78")},
					{Account: 324000, Objects: sie.TransactionObjects{{Dimension: 1, Number: "307"}}, Amount: sie.RequireAmountFromString("-292.32")},
					{Account: 3250, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("-28789.34")},
					{Account: 32500, Objects: sie.TransactionObjects{{Dimension: 1, Number: "200"}}, Amount: sie.RequireAmountFromString("-699.48")},
					{Account: 325000, Objects: sie.TransactionObjects{{Dimension: 1, Number: "307"}}, Amount: sie.RequireAmountFromString("-471.19")},
					{Account: 3260, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("-3938.27")},
					{Account: 32600, Objects: sie.TransactionObjects{{Dimension: 1, Number: "200"}}, Amount: sie.RequireAmountFromString("-153.12")},
					{Account: 354100, Objects: sie.TransactionObjects{{Dimension: 1, Number: "307"}}, Amount: sie.RequireAmountFromString("-20.51")},
					{Account: 3610, Objects: sie.TransactionObjects{{Dimension: 1, Number: "307"}}, Amount: sie.RequireAmountFromString("-9509.45")},
					{Account: 3614, Objects: sie.TransactionObjects{{Dimension: 1, Number: "30"}}, Amount: sie.RequireAmountFromString("-7767.86")},
					{Account: 3616, Objects: sie.TransactionObjects{{Dimension: 1, Number: "200"}}, Amount: sie.RequireAmountFromString("-11115.12")},
					{Account: 3630, Objects: sie.TransactionObjects{{Dimension: 1, Number: "600"}}, Amount: sie.RequireAmountFromString("-11310.00")},
					{Account: 3650, Objects: sie.TransactionObjects{{Dimension: 1, Number: "307"}}, Amount: sie.RequireAmountFromString("-3097.20")},
					{Account: 3660, Objects: sie.TransactionObjects{{Dimension: 1, Number: "200"}}, Amount: sie.RequireAmountFromString("-410.64")},
					{Account: 6050, Objects: sie.TransactionObjects{{Dimension: 1, Number: "101"}}, Amount: sie.RequireAmountFromString("96.32")},
				},
			},
		},
	}

	got, err := sie.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}

	want, err := os.ReadFile("testdata/example.si")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(got, want) {
		gotLines := strings.Split(string(got), "\r\n")
		wantLines := strings.Split(string(want), "\r\n")
		for i := 0; i < len(gotLines) || i < len(wantLines); i++ {
			var g, w string
			if i < len(gotLines) {
				g = gotLines[i]
			}
			if i < len(wantLines) {
				w = wantLines[i]
			}
			if g != w {
				t.Fatalf("line %d:\n got: %q\nwant: %q", i+1, g, w)
			}
		}
	}
}
