# DeleteShippingDetails

A Gin-based microservice created with ginshot.

## Structure

```bash
delete_shipping_details/
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
│   │   └── delete_shipping_details-controller.go
│   ├── data/
│   │   ├── messages/
│   │   │   ├── replication_topic_message.go
│   │   │   └── shipping_details_topic_message.go
│   │   ├── models/
│   │   │   ├── address.go
│   │   │   └── shipping_details.go
│   │   ├── requests/
│   │   │   └── delete_shipping_details_request.go
│   │   └── responses/
│   │       └── delete_shipping_details_response.go
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
│       ├── delete_shipping_details-service-impl.go
│       └── delete_shipping_details-service.go
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

	