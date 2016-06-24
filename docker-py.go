>>> import docker
>>> client = Client(base_url='unix://var/run/docker.sock')
>>> client.version()
{u'KernelVersion': u'4.4.12-boot2docker', u'Os': u'linux', u'BuildTime': u'2016-06-01T21:20:08.558909126+00:00', u'ApiVersion': u'1.23', u'Version': u'1.11.2', u'GitCommit': u'b9f10c9', u'Arch': u'amd64', u'GoVersion': u'go1.5.4'}
>>> [container['Command'] for container in client.containers()]
[u'python3 -m http.server', u'python app.py', u"/bin/sh -c 'while true; do echo hello world; sleep 1; done'"]
>>>
