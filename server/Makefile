watch:
	-pkill -9 fishcmd
	-pkill -9 mymanager
	fishcmd run --watch &
	tail -f nohup.out
release:
	go install mymanager
	rm -rf mymanager
	cp ${GOPATH}/bin/mymanager .
	sudo supervisorctl restart mymanager_server
	sudo tail -f ../data/log/access.log
stop:
	sudo supervisorctl stop mymanager_server 
	-pkill -9 fishcmd
	-pkill -9 mymanager
gen:
	go generate
