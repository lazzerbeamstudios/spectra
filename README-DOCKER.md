## Docker

**docker**

    docker login
    docker logout

    docker ps
    docker stats
    docker images
    docker system prune
    docker system prune -a

    docker network create [network] --driver=bridge --attachable

**docker compose**

    docker compose -f [file].yaml ps

    docker compose -f [file].yaml build
    docker compose -f [file].yaml build --no-cache

    docker compose -f [file].yaml up
    docker compose -f [file].yaml up -d

    docker compose -f [file].yaml down
