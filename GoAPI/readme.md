# Pre-Checks and Running CockroachDB

Ensure that docker and docker-compose are installed on your system along with Go.
  1. `sudo dnf install docker docker-compose golang`
  2. `sudo apt install docker docker-compose golang`

## CockroachDB In Docker
In order to run CockroachDB in docker and have all nodes communicate, a network and volumes need to be created before containers are composed.
  1. Create network: `sudo docker network create -d bridge roachnet`
  2. Create volumes for each node: `sudo docker volume create roach1 && sudo docker volume create roach2 && sudo docker volume create roach3`
  3. Compose containers: `sudo docker-compose up -d`
  4. Initialize cluster `sudo docker exec -it roach1 ./cockroach --host=roach1:26357 init --insecure`
