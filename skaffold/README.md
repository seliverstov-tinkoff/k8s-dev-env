## skaffold
Решение от google - cli 
про установку и фичи подробно есть на их стайте (https://skaffold.dev/)

Текущий конфигурационный файл `skaffold.yaml` имеет настройки `docker-desktop` как кластер и ореинтирован на локальный docker для хранилища.
Для запуска с пробросом портов в режиме разработчика надо выполнить команду `skaffold dev --port-forward`