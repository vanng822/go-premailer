test:
	make -j test_premailer test_cmd-server test_cmd-script

test_premailer:
	cd premailer && go test -v -cover

test_cmd-%:
	cd cmd/$* && go test -v -cover

bench:
	cd premailer && go test -bench=.
