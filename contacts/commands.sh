docker run -d --name pg -e POSTGRES_PASSWORD=admin -e POSTGRES_USER=jiten -e POSTGRES_DB=sample -p 5432:5432 postgress

go run main.go -stderrthreshold=FATAL