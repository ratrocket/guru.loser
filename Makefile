all:

progs: assemble serve

assemble: src/assemble.go
	go build -o bin/assemble src/assemble.go

serve: src/serve.go
	go build -o bin/serve src/serve.go

# For local testing.
localhost:
	rm -f index.html
	./bin/assemble > index.html

# For a "release"; trailing slashes are necessary.
guru.loser:
	rm -f index.html
	./bin/assemble -img http://d1nubnl3w13qxh.cloudfront.net/ \
		-asst http://d1nubnl3w13qxh.cloudfront.net/ \
		> index.html

clean:
	rm -f bin/{assemble,serve}
