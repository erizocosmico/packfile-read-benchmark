package benchmark

import (
	"testing"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

var (
	r                  *git.Repository
	decoderFixtures    *DecoderFixtures
	repositoryFixtures *RepositoryFixtures
)

func init() {
	var err error
	r, repositoryFixtures, decoderFixtures, err = Fixtures("repo")
	if err != nil {
		panic(err)
	}
}

func BenchmarkDecoder(b *testing.B) {
	b.Run("first", func(b *testing.B) {
		runDecoderCase(b, decoderFixtures.First)
	})

	b.Run("first 100", func(b *testing.B) {
		runDecoderCase(b, decoderFixtures.First100)
	})

	b.Run("first 100 skip 5", func(b *testing.B) {
		runDecoderCase(b, decoderFixtures.First100Skip5)
	})

	b.Run("last", func(b *testing.B) {
		runDecoderCase(b, decoderFixtures.Last)
	})
}

func runDecoderCase(b *testing.B, offsets []uint64) {
	for i := 0; i < b.N; i++ {
		dec, err := LoadPackfile("repo", r.Storer)
		if err != nil {
			b.Fatal("unable to load packfile")
		}

		for _, offset := range offsets {
			_, err := dec.DecodeObjectAt(int64(offset))
			if err != nil {
				b.Errorf("error getting object at index %d: %s", offset, err)
			}
		}
	}
}

func runRepositoryCase(b *testing.B, hashes []plumbing.Hash) {
	for i := 0; i < b.N; i++ {
		r, err := git.PlainOpen("repo")
		if err != nil {
			b.Fatal("unable to load repo")
		}

		for _, hash := range hashes {
			_, err := r.Object(plumbing.BlobObject, hash)
			if err != nil {
				b.Errorf("error getting object %s: %s", hash, err)
			}
		}
	}
}

func BenchmarkRepository(b *testing.B) {
	b.Run("first", func(b *testing.B) {
		runRepositoryCase(b, repositoryFixtures.First)
	})

	b.Run("first 100", func(b *testing.B) {
		runRepositoryCase(b, repositoryFixtures.First100)
	})

	b.Run("first 100 skip 5", func(b *testing.B) {
		runRepositoryCase(b, repositoryFixtures.First100Skip5)
	})

	b.Run("last", func(b *testing.B) {
		runRepositoryCase(b, repositoryFixtures.Last)
	})
}
