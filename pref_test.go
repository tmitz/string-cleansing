package stringcleansing

import "testing"
import "reflect"

func TestPrefixPref(t *testing.T) {
	cases := []struct {
		address, pref, want string
	}{
		{"島根県松江市春日町543-2", "島根県", "島根県松江市春日町543-2"},
		{"東津田町1877番地3", "島根県", "島根県東津田町1877番地3"},
		{"松江市玉湯町湯町１９２４－１", "島根県", "島根県松江市玉湯町湯町１９２４－１"},
		{"", "", ""},
	}

	for _, c := range cases {
		if got := PrefixPref(c.address, c.pref); got != c.want {
			t.Errorf("PrefixPref(%s, %s) == %q, want: %q", c.address, c.pref, got, c.want)
		}
	}
}

func TestPrefCodeString(t *testing.T) {
	cases := []struct {
		in   string
		want Pref
	}{
		{"02", Pref{Code: "aomori", Address: "青森県"}},
		{"44", Pref{Code: "oita", Address: "大分県"}},
	}

	for _, c := range cases {
		if got, _ := PrefCodeString(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("PrefCodeString(%s) == %q, want: %q", c.in, got, c.want)
		}
	}
}
