# OTEL BASIC EXAMPLE


The service "A" will call service "B", both endpoints will generate traces and spans, you can see at ```localhost:16686```

To start service:
```bash
make run
```

To do a request:
```bash
make test
```

To stop service:
```bash
make stop
```

## Traces: service A and service B
![jaegerUI](Screenshot2024-04-02.png)

## Spans:
![jaegerUI](Screenshot2024-04-02(1).png)
