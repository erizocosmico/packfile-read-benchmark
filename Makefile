bench: repo
	go test -bench=. -benchmem

repo:
	git clone https://github.com/git/git repo

.PHONY: bench