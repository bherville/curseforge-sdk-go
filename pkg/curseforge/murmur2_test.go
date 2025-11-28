package curseforge

import (
"bytes"
"testing"
)

func TestFilterWhitespace(t *testing.T) {
tests := []struct {
name     string
input    []byte
expected []byte
}{
{
name:     "no whitespace",
input:    []byte("hello"),
expected: []byte("hello"),
},
{
name:     "with spaces",
input:    []byte("hello world"),
expected: []byte("helloworld"),
},
{
name:     "with tabs",
input:    []byte("hello\tworld"),
expected: []byte("helloworld"),
},
{
name:     "with newlines",
input:    []byte("hello\nworld"),
expected: []byte("helloworld"),
},
{
name:     "with carriage return",
input:    []byte("hello\rworld"),
expected: []byte("helloworld"),
},
{
name:     "with CRLF",
input:    []byte("hello\r\nworld"),
expected: []byte("helloworld"),
},
{
name:     "all whitespace types",
input:    []byte("a b\tc\nd\re"),
expected: []byte("abcde"),
},
{
name:     "empty input",
input:    []byte{},
expected: []byte{},
},
{
name:     "only whitespace",
input:    []byte(" \t\n\r"),
expected: []byte{},
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
result := filterWhitespace(tt.input)
if !bytes.Equal(result, tt.expected) {
t.Errorf("filterWhitespace(%q) = %q, want %q", tt.input, result, tt.expected)
}
})
}
}

func TestComputeFingerprint(t *testing.T) {
// Test that fingerprint computation filters whitespace correctly
input1 := []byte("hello")
input2 := []byte("hel lo")
input3 := []byte("h e l l o")

fp1 := ComputeFingerprint(input1)
fp2 := ComputeFingerprint(input2)
fp3 := ComputeFingerprint(input3)

// All should produce the same fingerprint since whitespace is filtered
if fp1 != fp2 {
t.Errorf("Fingerprints differ: %d vs %d", fp1, fp2)
}
if fp1 != fp3 {
t.Errorf("Fingerprints differ: %d vs %d", fp1, fp3)
}
}

func TestComputeFingerprintDeterministic(t *testing.T) {
// Ensure fingerprint is deterministic
data := []byte("test data for fingerprint")

fp1 := ComputeFingerprint(data)
fp2 := ComputeFingerprint(data)

if fp1 != fp2 {
t.Errorf("Fingerprint not deterministic: %d vs %d", fp1, fp2)
}
}

func TestComputeFingerprintFromReader(t *testing.T) {
data := []byte("test data for fingerprint")
reader := bytes.NewReader(data)

fp, err := ComputeFingerprintFromReader(reader)
if err != nil {
t.Fatalf("ComputeFingerprintFromReader failed: %v", err)
}

expectedFp := ComputeFingerprint(data)
if fp != expectedFp {
t.Errorf("Fingerprint from reader = %d, want %d", fp, expectedFp)
}
}

func TestNormalizeLineEndings(t *testing.T) {
tests := []struct {
name     string
input    []byte
expected []byte
}{
{
name:     "no line endings",
input:    []byte("hello"),
expected: []byte("hello"),
},
{
name:     "LF only",
input:    []byte("hello\nworld"),
expected: []byte("hello\nworld"),
},
{
name:     "CR only",
input:    []byte("hello\rworld"),
expected: []byte("hello\nworld"),
},
{
name:     "CRLF",
input:    []byte("hello\r\nworld"),
expected: []byte("hello\nworld"),
},
{
name:     "mixed line endings",
input:    []byte("line1\r\nline2\rline3\nline4"),
expected: []byte("line1\nline2\nline3\nline4"),
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
result := normalizeLineEndings(tt.input)
if !bytes.Equal(result, tt.expected) {
t.Errorf("normalizeLineEndings(%q) = %q, want %q", tt.input, result, tt.expected)
}
})
}
}

func TestFingerprintDifferentInputs(t *testing.T) {
// Different inputs should produce different fingerprints
fp1 := ComputeFingerprint([]byte("hello"))
fp2 := ComputeFingerprint([]byte("world"))

if fp1 == fp2 {
t.Error("Different inputs produced same fingerprint")
}
}
