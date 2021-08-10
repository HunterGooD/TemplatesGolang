function start {
    # Запускает только postgres без админера 
    if [$1 = "postgres"]
        then
            docker-compose -f "./deployments/docker/database/postgresql/docker-compose.yml" up -d postgres
    # Запускает только mysql без админера 
    elif [$1 = "mysql"]
        then
             docker-compose -f "./deployments/docker/database/mysql/docker-compose.yml" up -d
    elif [$1 = "postgres-all"]
        then 
            docker-compose -f "./deployments/docker/database/postgresql/docker-compose.yml" up -d
    fi
}

if [$1 = "start"]
    then
        if [ -z $2]
            then
            start $2
        else
            echo "Не указан контейнер для запуска"
        fi
fi