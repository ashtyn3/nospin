import subprocess as sub
import os 
import getpass

def git_clone(url):
    proc = sub.Popen(['git', 'clone', '--progress',url],
        stdout=sub.PIPE,
        stderr=sub.STDOUT)

    while proc.poll() is None:
        out = proc.stdout.readline()
        if out.decode().strip().startswith("remote"):
            print(out.decode())
    print("finished cloning:",url)
try:
    sub.run('zap', capture_output=True)
    print("zap is installed")
    print("entering ../")
    url = input("url to server: ")
    pd = getpass.getpass(prompt='password to server: ')
    print("building with credentials")
    sub.run('cd ../cmd && pwd && zap -env pd='+pd+',url='+url+" -in ../auth/auth.go",shell=True,env=os.environ)
    print("moving build to bin")
    sub.run('mkdir ../bin && mv ../cmd/cmd ../bin/qt',shell=True)
except Exception:
    git_clone('https://github.com/doublequotation/zap/')
    #path = 'GOPATH='+os.environ["GOPATH"]
    print("entering zap/cmd")
    print("building zap")
    sub.run('cd zap/cmd && pwd && go build',shell=True)
    os.environ["zap"] = "zap/cmd/./cmd"
    print("leaving zap/cmd")
    print("entering ../")
    url = input("url to server: ")
    pd = getpass.getpass(prompt='password to server: ')
    print("building with credentials")
    sub.run('cd ../cmd && pwd && ../tools/zap/cmd/./cmd -env pd='+pd+',url='+url+" -in ../auth/auth.go",shell=True,env=os.environ)
    print("moving build to bin")
    sub.run('mkdir ../bin && mv ../cmd/cmd ../bin/qt',shell=True)
    print("Cleaning up")
    sub.run('sudo rm -r zap',shell=True)
