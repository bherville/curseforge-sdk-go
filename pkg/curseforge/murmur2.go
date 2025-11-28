package curseforge

import (
"io"
"os"
)

// Murmur2 hash implementation for CurseForge fingerprinting
// CurseForge uses a specific variant of Murmur2 with seed 1 and specific byte filtering

const (
murmur2Seed uint32 = 1
murmur2M    uint32 = 0x5bd1e995
murmur2R    int    = 24
)

// ComputeFileFingerprint computes the CurseForge fingerprint (Murmur2 hash) for a file
// This is used for matching files against the CurseForge fingerprint database
func ComputeFileFingerprint(filePath string) (int64, error) {
file, err := os.Open(filePath)
if err != nil {
return 0, err
}
defer file.Close()

return ComputeFingerprintFromReader(file)
}

// ComputeFingerprintFromReader computes the CurseForge fingerprint from an io.Reader
func ComputeFingerprintFromReader(reader io.Reader) (int64, error) {
data, err := io.ReadAll(reader)
if err != nil {
return 0, err
}

return ComputeFingerprint(data), nil
}

// ComputeFingerprint computes the CurseForge fingerprint from byte data
// CurseForge filters out whitespace characters (9, 10, 13, 32) before hashing
func ComputeFingerprint(data []byte) int64 {
// Filter out whitespace characters as CurseForge does
filtered := filterWhitespace(data)

hash := computeMurmur2(filtered)
return int64(hash)
}

// filterWhitespace removes whitespace characters that CurseForge ignores
// Characters filtered: tab (9), newline (10), carriage return (13), space (32)
func filterWhitespace(data []byte) []byte {
result := make([]byte, 0, len(data))
for _, b := range data {
if b != 9 && b != 10 && b != 13 && b != 32 {
result = append(result, b)
}
}
return result
}

// computeMurmur2 computes the Murmur2 hash used by CurseForge
func computeMurmur2(data []byte) uint32 {
length := len(data)
h := murmur2Seed ^ uint32(length)

// Process 4-byte chunks
nblocks := length / 4
for i := 0; i < nblocks; i++ {
idx := i * 4
k := uint32(data[idx]) |
uint32(data[idx+1])<<8 |
uint32(data[idx+2])<<16 |
uint32(data[idx+3])<<24

k *= murmur2M
k ^= k >> uint32(murmur2R)
k *= murmur2M

h *= murmur2M
h ^= k
}

// Process remaining bytes
tail := data[nblocks*4:]
switch len(tail) {
case 3:
h ^= uint32(tail[2]) << 16
fallthrough
case 2:
h ^= uint32(tail[1]) << 8
fallthrough
case 1:
h ^= uint32(tail[0])
h *= murmur2M
}

// Finalization
h ^= h >> 13
h *= murmur2M
h ^= h >> 15

return h
}

// ComputeNormalizedFingerprint computes fingerprint with additional normalization
// This normalizes line endings and other whitespace variations
func ComputeNormalizedFingerprint(data []byte) int64 {
// Normalize line endings to LF
normalized := normalizeLineEndings(data)
return ComputeFingerprint(normalized)
}

// normalizeLineEndings converts all line endings to LF
func normalizeLineEndings(data []byte) []byte {
result := make([]byte, 0, len(data))
i := 0
for i < len(data) {
if data[i] == '\r' {
if i+1 < len(data) && data[i+1] == '\n' {
// CRLF -> LF
result = append(result, '\n')
i += 2
} else {
// CR -> LF
result = append(result, '\n')
i++
}
} else {
result = append(result, data[i])
i++
}
}
return result
}
