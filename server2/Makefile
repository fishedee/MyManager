.PHONY:watch release
watch:
	systemfd --no-pid -s http::8088 -- cargo watch -x run
release:
	-pkill server
	cargo build --release
	cp target/release/server .
	./server
bench:
	#wrk -t12 -c100 -d30s http://localhost:8088/user/get
	ab -n 10000 -c 100 -k 127.0.0.1:8180/user/get
	#ab -n 10000 -c 100 -k 127.0.0.1:8088/user/get
