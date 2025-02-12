# GetShippingDetailsById

A Gin-based microservice created with ginshot.

## Structure

```bash
get_shipping_details_by_id/
├── app/
│   └── main.go
├── config/
│   ├── cors/
│   │   └── cors.go
│   └── kafka/
│       └── offset_manager.go
├── deployments/
│   └── docker-compose.yml
├── docs/
├── internal/
│   ├── bootstrap/
│   │   └── setup/
│   │       └── setup.go
│   ├── controller/
│   │   └── get_shipping_details_by_id-controller.go
│   ├── data/
│   │   ├── messages/
│   │   │   ├── replication_topic_message.go
│   │   │   └── shipping_details_topic_message.go
│   │   ├── models/
│   │   │   ├── address.go
│   │   │   └── shipping_details.go
│   │   ├── requests/
│   │   │   └── get_shipping_details_by_id_request.go
│   │   └── responses/
│   │       └── get_shipping_details_by_id_response.go
│   ├── event/
│   │   ├── consumer.go
│   │   ├── producer.go
│   │   └── start_consumer.go
│   ├── middleware/
│   ├── repository/
│   │   ├── dbcontext.go
│   │   ├── shipping_details_replicator.go
│   │   └── shipping_details_repository.go
│   ├── secrets/
│   │   └── shipping_details-env.go
│   └── service/
│       ├── get_shipping_details_by_id-service-impl.go
│       └── get_shipping_details_by_id-service.go
├── pkg/
├── router/
│   └── router.go
├── test/
├── .dockerignore
├── .env
├── Dockerfile
├── README.md
├── go.mod
└── go.sum
```