package trigram

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func mktri(s string) T { return T(uint32(s[0])<<16 | uint32(s[1])<<8 | uint32(s[2])) }

func mktris(ss ...string) []T {
	var ts []T
	for _, s := range ss {
		ts = append(ts, mktri(s))
	}
	return ts
}

func TestExtract(t *testing.T) {

	tests := []struct {
		s    string
		want []T
	}{
		{"", nil},
		{"a", nil},
		{"ab", nil},
		{"abc", mktris("abc")},
		{"abcabc", mktris("abc", "bca", "cab")},
		{"abcd", mktris("abc", "bcd")},
	}

	for _, tt := range tests {
		if got := Extract(tt.s, nil); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Extract(%q)=%+v, want %+v", tt.s, got, tt.want)
		}
	}
}

func TestQuery(t *testing.T) {

	s := []string{
		"foo",
		"foobar",
		"foobfoo",
		"quxzoot",
		"zotzot",
		"azotfoba",
	}

	idx := NewIndex(s)

	tests := []struct {
		q   string
		ids []DocID
	}{
		{"", []DocID{0, 1, 2, 3, 4, 5}},
		{"foo", []DocID{0, 1, 2}},
		{"foob", []DocID{1, 2}},
		{"zot", []DocID{4, 5}},
		{"oba", []DocID{1, 5}},
	}

	for _, tt := range tests {
		if got := idx.Query(tt.q); !reflect.DeepEqual(got, tt.ids) {
			t.Errorf("Query(%q)=%+v, want %+v", tt.q, got, tt.ids)
		}
	}

	idx.Add("zlot")
	docs := idx.Query("lot")
	if len(docs) != 1 || docs[0] != 6 {
		t.Errorf("Query(`lot`)=%+v, want []DocID{6}", docs)
	}

	idx.Delete("foobar", 1)
	docs = idx.Query("fooba")
	if len(docs) != 0 {
		t.Errorf("Query(`fooba`)=%+v, want []DocID{}", docs)
	}
}

func TestFullPrune(t *testing.T) {

	s := []string{
		"foo",
		"foobar",
		"foobfoo",
		"quxzoot",
		"zotzot",
		"azotfoba",
	}

	idx := NewIndex(s)
	idx.Prune(0)

	tests := []struct {
		q   string
		ids []DocID
	}{
		{"", []DocID{0, 1, 2, 3, 4, 5}},
		{"foo", []DocID{0, 1, 2, 3, 4, 5}},
		{"foob", []DocID{0, 1, 2, 3, 4, 5}},
		{"zot", []DocID{0, 1, 2, 3, 4, 5}},
		{"oba", []DocID{0, 1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		if got := idx.Query(tt.q); !reflect.DeepEqual(got, tt.ids) {
			t.Errorf("Query(%q)=%+v, want %+v", tt.q, got, tt.ids)
		}
	}

	idx.Add("ahafoo")
	tests = []struct {
		q   string
		ids []DocID
	}{
		{"", []DocID{0, 1, 2, 3, 4, 5, 6}},
		{"foo", []DocID{0, 1, 2, 3, 4, 5, 6}},
		{"foob", []DocID{0, 1, 2, 3, 4, 5, 6}},
		{"zot", []DocID{0, 1, 2, 3, 4, 5, 6}},
		{"oba", []DocID{0, 1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		if got := idx.Query(tt.q); !reflect.DeepEqual(got, tt.ids) {
			t.Errorf("Query(%q)=%+v, want %+v", tt.q, got, tt.ids)
		}
	}
}

var result int
var format = "general.tuning.%s.%s.%s.%s.%s"
var podNames = getPodNames()
var globNames = getGlobNames()
var directoryNames = getDirectoryNames()
var appNames = getAppNames()
var metricNames = getMetricNames()
var f1 = getFileNames(10)
var f2 = getFileNames(11)
var f3 = getFileNames(12)
var f4 = getFileNames(14)
var f5 = getFileNames(15)


var idx1 = NewIndex(f1)
var idx2 = NewIndex(f2)
var idx3 = NewIndex(f3)
var idx4 = NewIndex(f4)
var idx5 = NewIndex(f5)


func BenchmarkQuery1_1(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], directoryNames[0], appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx1.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery1_2(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx1.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery1_3(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", "*", "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx1.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery2_1(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], directoryNames[0], appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx2.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a p	`ackage level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery2_2(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx2.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery2_3(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", "*", "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx2.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery3_1(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], directoryNames[0], appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx3.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery3_2(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx3.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery3_3(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", "*", "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx3.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery4_1(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], directoryNames[0], appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx4.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery4_2(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx4.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery4_3(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", "*", "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx4.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery5_1(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], directoryNames[0], appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx5.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery5_2(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", appNames[0], "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx5.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}

func BenchmarkQuery5_3(b *testing.B) {
	var resultLength int
	q := fmt.Sprintf(format, "*", globNames[0], "*", "*", "*")

	for n := 0; n < b.N; n++ {
		ts := extractTrigrams(q)
		r := idx5.QueryTrigrams(ts)
		resultLength = len(r)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = resultLength
}


func getFileNames(n int) []string {
	var fileNames []string

	for i := 0; i < n; i++ {
		podName := podNames[i]
		for j := 0; j < n; j++ {
			globName := globNames[j]
			for k := 0; k < n; k++ {
				directoryName := directoryNames[k]
				for l := 0; l < n; l++ {
					appName := appNames[l]
					for m := 0; m < n; m++ {
						metricName := metricNames[m]
						x := fmt.Sprintf(format, podName, globName, directoryName, appName, metricName)
						fileNames = append(fileNames, x)
					}
				}
			}
		}

	}
	return fileNames
}

func getPodNames() []string {
	var podNames []string
	for i := 0; i < 100; i++ {
		podNames[i] = RandStringRunes(8)
	}
	return podNames
}

func getGlobNames() []string {
	var globNames []string
	for i := 0; i < 100; i++ {
		globNames[i] = fmt.Sprintf("glob-%d", i)
	}
	return globNames
}

func getDirectoryNames() [100]string {
	var directoryNames [100]string
	for i := 0; i < 100; i++ {
		directoryNames[i] = fmt.Sprintf("dir-%d", i)
	}
	return directoryNames
}

func getAppNames() [100]string {
	var appNames [100]string
	for i := 0; i < 100; i++ {
		appNames[i] = fmt.Sprintf("app-%d", i)
	}
	return appNames
}

func getMetricNames() [100]string {
	var metricNames [100]string
	for i := 0; i < 100; i++ {
		metricNames[i] = fmt.Sprintf("app-%d", i)
	}
	return metricNames
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func extractTrigrams(query string) []T {

	if len(query) < 3 {
		return nil
	}

	var start int
	var i int

	var trigrams []T

	for i < len(query) {
		if query[i] == '[' || query[i] == '*' || query[i] == '?' {
			trigrams = Extract(query[start:i], trigrams)

			if query[i] == '[' {
				for i < len(query) && query[i] != ']' {
					i++
				}
			}

			start = i + 1
		}
		i++
	}

	if start < i {
		trigrams = Extract(query[start:i], trigrams)
	}

	return trigrams
}
