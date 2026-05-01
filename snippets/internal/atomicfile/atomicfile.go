// Package atomicfile writes a file by way of a same-directory tempfile,
// fsync, and rename. Crash-safe and concurrent-safe across parallel
// renders that target distinct destinations. Used by every render
// adapter that emits files into the consumer checkout.
package atomicfile

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path/filepath"
)

// Write atomically replaces path with data. Preserves the destination's
// existing permission bits (so a checkout that has tightened a file to
// read-only mode for CODEOWNER reasons doesn't quietly get reset to
// 0644). New files are created with mode 0644.
func Write(path string, data []byte) error {
	dir := filepath.Dir(path)
	mode := os.FileMode(0o644)
	if info, err := os.Stat(path); err == nil {
		mode = info.Mode().Perm()
	}
	var sfx [8]byte
	if _, err := rand.Read(sfx[:]); err != nil {
		return err
	}
	tmp := filepath.Join(dir, "."+filepath.Base(path)+".sdk-snippets."+hex.EncodeToString(sfx[:])+".tmp")
	f, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, mode)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		os.Remove(tmp)
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		os.Remove(tmp)
		return err
	}
	if err := f.Close(); err != nil {
		os.Remove(tmp)
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return err
	}
	return nil
}
