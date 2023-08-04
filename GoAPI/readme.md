Create network: `sudo docker network create -d bridge roachnet`
Create volumes for each node: `sudo docker volume create roach1 roach2 roach3`
Compose containers: `sudo docker-compose up -d`
Initialize cluster `sudo docker exec -it roach1 ./cockroach --host=roach1:26357 init --insecure`
