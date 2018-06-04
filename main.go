package benchmark

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/format/idxfile"
	"gopkg.in/src-d/go-git.v4/plumbing/format/packfile"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"

	git "gopkg.in/src-d/go-git.v4"
)

type RepositoryFixtures struct {
	First         []plumbing.Hash
	First100      []plumbing.Hash
	First100Skip5 []plumbing.Hash
	Last          []plumbing.Hash
}

type DecoderFixtures struct {
	First         []uint64
	First100      []uint64
	First100Skip5 []uint64
	Last          []uint64
}

func Fixtures(path string) (*git.Repository, *RepositoryFixtures, *DecoderFixtures, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, nil, nil, err
	}

	iter, err := r.BlobObjects()
	if err != nil {
		return nil, nil, nil, err
	}

	var (
		tc       RepositoryFixtures
		lastBlob *object.Blob
	)

	var i int
	err = iter.ForEach(func(b *object.Blob) error {
		lastBlob = b

		if len(tc.First) == 0 {
			tc.First = append(tc.First, b.Hash)
		}

		if len(tc.First100) < 100 {
			tc.First100 = append(tc.First100, b.Hash)
		}

		if len(tc.First100Skip5) < 100 && i%5 == 0 {
			tc.First100Skip5 = append(tc.First100Skip5, b.Hash)
		}

		i++
		return nil
	})
	if err != nil {
		return nil, nil, nil, err
	}

	tc.Last = append(tc.Last, lastBlob.Hash)
	index, err := loadIndex(path)
	if err != nil {
		return nil, nil, nil, err
	}

	var dt DecoderFixtures
	dt.First, err = hashesToOffsets(tc.First, index)
	if err != nil {
		return nil, nil, nil, err
	}

	dt.First100, err = hashesToOffsets(tc.First100, index)
	if err != nil {
		return nil, nil, nil, err
	}

	dt.First100Skip5, err = hashesToOffsets(tc.First100Skip5, index)
	if err != nil {
		return nil, nil, nil, err
	}

	dt.Last, err = hashesToOffsets(tc.Last, index)
	if err != nil {
		return nil, nil, nil, err
	}

	return r, &tc, &dt, nil
}

func hashesToOffsets(hashes []plumbing.Hash, index *packfile.Index) ([]uint64, error) {
	var result = make([]uint64, len(hashes))
	for i, hash := range hashes {
		entry, ok := index.LookupHash(hash)
		if !ok {
			return nil, fmt.Errorf("blob not found: %s", hash)
		}

		result[i] = entry.Offset
	}
	return result, nil
}

func loadIndex(path string) (*packfile.Index, error) {
	pattern := filepath.Join(path, ".git", "objects", "pack", "*.idx")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no index files found")
	}

	f, err := os.Open(files[0])
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var idxf idxfile.Idxfile
	if err := idxfile.NewDecoder(f).Decode(&idxf); err != nil {
		return nil, err
	}

	return packfile.NewIndexFromIdxFile(&idxf), nil
}

func LoadPackfile(path string, storer storer.EncodedObjectStorer) (*packfile.Decoder, error) {
	pattern := filepath.Join(path, ".git", "objects", "pack", "*.pack")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no pack files found")
	}

	f, err := os.Open(files[0])
	if err != nil {
		return nil, err
	}

	return packfile.NewDecoder(packfile.NewScanner(f), storer)
}
