# Allocated Notifications Service

Проект на **Go** с использованием **Kafka** и **PostgreSQL** для организации системы уведомлений.  
Сервис подписывается на события от различных сервисов продюссеров, обрабатывает их сообщения и отправляет уведомления (например, по Email через SendGrid).  

---

## Архитектура

**Producers (prds/service_X)**  
Несколько сервисов, которые публикуют события в разные Kafka-топики.  
Например: `service_1`, `service_2`, … — каждые N секунд отправляют JSON с информацией.

**Kafka**  
Используется как брокер сообщений для доставки событий между сервисами.

**Consumer (app/main.go)**  
Основное приложение:
 - Запускает HTTP API.
 - Запускает Kafka consumer.
 - Подключается к PostgreSQL.
 - Делегирует обработку сообщений сервисам уведомлений.

**DB Service (db_service/requests.go)**  
 - Получение информации о пользователях.
 - Сохранение истории уведомлений.

**Notifications Service (notifications_service/by_email.go)**  
Отправка email через [SendGrid](https://sendgrid.com).  
В реальном проекте можно добавить поддержку других каналов (SMS, push-уведомления и т. д.).

**Config (config/config.go, config.yaml)**  
Все настройки (Kafka, БД, SendGrid и пр.) вынесены в YAML-файл.

