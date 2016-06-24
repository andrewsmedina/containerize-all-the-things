$ docker run -d busybox /bin/sh -c "while true; do echo caipyra; sleep 1; done"
3b5bdc3785bda0dfa0753d09c1b519b88107dca03b00544b31cc497175244c38

$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS               NAMES
3b5bdc3785bd        busybox             "/bin/sh -c 'while tr"   About a minute ago   Up About a minute                       evil_mestorf

$ docker logs 3b5bdc3785bd
caipyra
caipyra
caipyra
caipyra
caipyra
caipyra
