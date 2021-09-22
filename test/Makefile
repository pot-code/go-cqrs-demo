.PHONY: stress-test up down

up:
	@docker compose up -d

down:
	@docker compose down

stress-test:
	@k6 --out influxdb=http://localhost:8086/k6test run stress.test.js