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
