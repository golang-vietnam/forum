package envy

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

const (
	NIL = "<nil>"
)

var parselnTests = []struct {
	// input
	in string
	// output
	key string
	val string
	err string
}{
	{"PORT=9090", "PORT", "9090", NIL},
	{"PORT9090", "", "", "missing delimiter '='"},
	{"", "", "", NIL},
	{"# A comment line", "", "", NIL},
        {"MYVAR=2020", "MYVAR", "2020", NIL},
	{"PORT =9090", "PORT", "9090", NIL},
	{`PORT="9090"`, "PORT", "9090", NIL},
	{`PORT='9090'`, "PORT", "9090", NIL},
	{"PORT= 9090", "PORT", "9090", NIL},
	{"URL=postgres://localhost?option=1", "URL", "postgres://localhost?option=1", NIL},
}

func Test_Simple_Parseln(t *testing.T) {
	for _, tt := range parselnTests {
		key, val, err := parseln(tt.in)
		expect(t, key, tt.key)
		expect(t, val, tt.val)
		expect(t, fmt.Sprint(err), tt.err)
	}
}

func Test_Load(t *testing.T) {
	buf := bytes.NewBufferString("PORT=9090\nMARTINI_ENV=dev\nHELLO='world'")

	_, err := Load(buf)
	expect(t, fmt.Sprint(err), NIL)
	expect(t, os.Getenv("PORT"), "9090")
	expect(t, os.Getenv("MARTINI_ENV"), "dev")
	expect(t, os.Getenv("HELLO"), "world")
}

func Test_Env(t *testing.T) {
	buf := bytes.NewBufferString("PORT=9090\nMARTINI_ENV=dev\nHELLO='world'")

	env, err := Load(buf)
	expect(t, fmt.Sprint(err), NIL)
	expect(t, env["PORT"], "9090")
	expect(t, env["MARTINI_ENV"], "dev")
	expect(t, env["HELLO"], "world")
}

func Test_MustGet(t *testing.T) {
	os.Setenv("FOO_BAR", "batbaz")
	os.Setenv("NOT_HERE", "")

	func() {
		defer func() {
			if err := recover(); err == nil {
				t.Error("Expected a panic")
			}
		}()

		MustGet("NOT_HERE")
	}()

	f := MustGet("FOO_BAR")
	expect(t, f, "batbaz")
}

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
