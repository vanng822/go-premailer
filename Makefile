test:
	make -j test_premailer test_cmd-server test_cmd-script

test_premailer:
	cd premailer && go test -v -cover

test_cmd-%:
	cd cmd/$* && go test -v -cover

gocyclo_all:
	make -j gocyclo-premailer gocyclo-cmd-script gocyclo-cmd-server

gocyclo-%:
	gocyclo -avg -over 15 $(subst -,/,$*)

bench:
	cd premailer && go test -bench=.
