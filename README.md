# TeamService

Сервис для создание команды, турнира, просмотр статистики

| Запрос |        Путь         |      Статусы       | Возвращает 
 |:------:|:-------------------:|:------------------:|:----------:
|  POST  |   /v1/team-create   | 200, 400, 500, 404 |    JSON    |
|  POST  | /v1/team-useradd:id | 200, 400, 500, 404 |    JSON    |
|  POST  |  /team-turadd/:id   | 200, 400, 500, 404 |    JSON    |
|  GET   |      /v1/teams      | 200, 400, 500, 404 |    JSON    |
|  GET   |    /v1/team/:id     | 200, 400, 500, 404 |    JSON    |
|  GET   |    /team-stats/     | 200, 400, 500, 404 |    JSON    |
|  GET   |   /team-stats/:id   | 200, 400, 500, 404 |    JSON    |
| PATCH  |    /v1/team-edit    | 200, 400, 500, 404 |    JSON    |
| PATCH  |  /v1/team-roleedit  | 200, 400, 500, 404 |    JSON    |
| PATCH  | /team-statsedit/:id | 200, 400, 500, 404 |    JSON    |
| PATCH  | /team-modifyelo/:id | 200, 400, 500, 404 |    JSON    |
| DELETE |    /team-del/:id    | 200, 400, 500, 404 |    JSON    |
| DELETE |  /team-userdel/:id  | 200, 400, 500, 404 |    JSON    |
| DELETE | /team-userleave/:id | 200, 400, 500, 404 |    JSON    |
| DELETE |  /team-turdel/:id   | 200, 400, 500, 404 |    JSON    |

