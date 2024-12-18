# HW 4. Elasticsearch & Kibana

## Фотоотчет

### Запуск команд в терминале

Сначала запустил скрипт без sudo, но пара команд скипнулась из-за отсутсвия прав рута. С sudo все норм было, но на команде restart elasticsearch почему-то скрипт завис, поэтому я его прервал и вручную отдельно запустил systemctl restart elasticsearch и systemctl restart kibana. Затем запустил манифест

![alt text](image.png)

![alt text](image-1.png)

Работаю на своей виртуалке, поэтому поменял в alias`ах ubuntu на misha и добавил в /home/misha/.bashrc

![alt text](image-2.png)

Вывел поды. Под с fluentd назывался fluentd-rhg7f. Вывел его логи
![alt text](image-3.png)

Потом в логах был большой конфигурационный файл (не стал его целиком вставлять)

![alt text](image-4.png)

А после файла вывело следующее:

![alt text](image-5.png)

### Работа в UI

В моем случае UI был доступен на localhost. Создал DataView:

![alt text](image-6.png)

А затем сделал дашборд с названиями контейнеров и количество записей по контейнеру.

![alt text](image-7.png)