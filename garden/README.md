## Garden

для установки решение воспользуйтесь инструкцией https://docs.garden.io/getting-started/1-installation

структура тестового проекта:
* Dockerfile - инструкции по созданию образа
* main.go - http сервер, который запускается в k8s
* project.garden.yml - инициализация проект и провайдеров
* service.garden.yml - деплоймент и настройки развертывания, включает в себя deployment

Для запуска проекта:
* запуск проекта `garden dev` (Эта версия не имеет hot reload)
* открываем `localhost/hello` или `http://127.0.0.1:8080/`
