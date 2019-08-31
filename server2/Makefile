.PHONY:watch release
watch:
	systemfd --no-pid -s http::8180 -- cargo watch -x run
release:
	-pkill server
	cargo build --release
	cp target/release/server .
	./server
bench:
	wrk -t12 -c100 -d20s http://0.0.0.0:8180/user/get
