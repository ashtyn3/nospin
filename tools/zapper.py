import subprocess as sub
import os 
import getpass
from dotenv import load_dotenv
import os.path
import shutil
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
    os.chdir("zap/cmd")
    sub.run('go build',shell=True)
    print("leaving zap/cmd")
    url = ""
    pd = ""
    if os.path.isfile("../../../.env"):
        if input("Build with credentials in .env file [y/n]: ") == "y":
            load_dotenv(dotenv_path="../../../.env")
            url = os.environ["url"]
            pd = os.environ["pd"]
        else:
            url = input("url to server: ")
            pd = getpass.getpass(prompt='password to server: ')
    print("building with credentials")
    print("entering ../cmd")
    os.chdir("../../../cmd")
    sub.call('../tools/zap/cmd/cmd -env pd='+pd+',url='+url+" -in ../auth/auth.go",shell=True,env=os.environ)
    print("moving build to bin")
    if os.path.isdir("../bin"):
        shutil.rmtree("../bin")
    os.mkdir("../bin")
    os.rename("cmd", "../bin/qt")
    print("leaving cmd")
    print("entering ../tools")
    os.chdir("../tools")
    print("Cleaning up")
    shutil.rmtree("zap")
