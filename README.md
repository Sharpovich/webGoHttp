# Репозиторий HTTP GO git.

Доступна демонстрация сервиса для обзора

Для запуска демонстрации необходимо склонировать репозиторий на VM с установленным docker и docker-compose

    git clone https://github.com/Sharpovich/webGoHttp.git
    
На VM в каталоге склонированного проекта требуется вписать команду

    sudo sh runner_debug.sh
    
После чего перейти в браузере по адресу:

    localhost:8080
    
Предварительно произведя проброс портов (ssh:22/http:8080) из VM во внешнюю среду
