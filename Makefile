all:
	docker run -v$(HOME):/root --rm golang:1.8 /bin/bash -c 'cd /root/go/src/github.com/aandryashin/plug/plugins/one && GOPATH=/root/go go build -buildmode plugin -o one.so'
	docker run -v$(HOME):/root --rm golang:1.8 /bin/bash -c 'cd /root/go/src/github.com/aandryashin/plug/plugins/two && GOPATH=/root/go go build -buildmode plugin -o two.so'
	docker run -v$(HOME):/root --rm golang:1.8 /bin/bash -c 'cd /root/go/src/github.com/aandryashin/plug && GOPATH=/root/go go build'
	docker build -t plug:0.1.0 .

clean:
	rm plug
	rm plugins/one/one.so
	rm plugins/two/two.so
	docker rmi -f plug:0.1.0
	