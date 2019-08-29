.PHONY:watch release
watch:
	systemfd --no-pid -s http::8088 -- cargo watch -x run
release:
	-pkill server2
	cargo build
	./target/debug/server2
