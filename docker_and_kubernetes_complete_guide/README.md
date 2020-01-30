# Docker and Kubernetes: The Complete Guide

## Resources
draw.io

## Notes
docker run = docker create + docker start  

### Section 2  
**Restarting Stopped Containers**  
docker start does not show you output coming from the container whereas docker run does.  

Example:  
docker start -a b3433f5a91d38488b51d5f50f037ec6c27925151ef2ddc656965b588182fe0c0   
-a tells the container to show output to the terminal  

docker ps --all to show all containters, even ones run in the past.  
this command will list the container id, then using container id as follows:  
docker start -a <container id> will re-run whatever that container ran and show the   
container's output to the terminal.  

Retrieving Log Outputs  
    docker logs <container id>

Stopping Containers   
    docker stop <container id> - sends signal (sigterm) to running process in container.  If
    container does not automatically stop in 10 seconds then docker issues docker kill command.  
    
    docker kill <container id> - sends kill signal (sigkill) to running container.  Immediate stop,
    now.  

Multi-Command Containers  
   docker exec -it <container id> <command>   

The Purpose of the IT Flag  
    Every process run in a container has 3 communication channels, STDIN, STDOUT, STDERR.  
    The IT flag is actuall two flags, -i and -t.  
    The -i attaches our terminal gets directed to STDIN.  
    The -t flag ensures text entered in and output is formatted for the screen.  

Container Isolation  
    File systems are not automatically shared between 2 containers.  

## Section 3  

Docker tagging convention:

docker build -t <docker id>/<image name>:<version> .

## Section 4  

Started simpleweb  

## Section 5  

## Section 6

Flow focuses around GitHub Repo
Two branches: feature, and master

Review section 62 - Flow Specifics.  This is a very good description of development workflow.  

## Section 7  
If you are running on Windows, please read this: Create-React-App has some issues detecting when
files get changed on Windows based machines.  To fix this, please do the following:

In the root project directory, create a file called .env

Add the following text to the file and save it: CHOKIDAR_USEPOLLING=true

That's all!

For more on why this is required, you can check out:
https://facebook.github.io/create-react-app/docs/troubleshooting#npm-start-doesn-t-detect-changes
