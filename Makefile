INDEX=index.html

all:

progs: assemble serve

assemble: src/assemble.go
	go build -o bin/assemble src/assemble.go

serve: src/serve.go
	go build -o bin/serve src/serve.go

# For local testing.
localhost:
	rm -f $(INDEX)
	./bin/assemble > $(INDEX)

# For a "release"; trailing slashes are necessary.
guru.loser:
	rm -f $(INDEX)
	./bin/assemble -img http://d1nubnl3w13qxh.cloudfront.net/ \
		-asst http://d1nubnl3w13qxh.cloudfront.net/ \
		> $(INDEX)

clean:
	rm -f bin/{assemble,serve}
