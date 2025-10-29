FROM golang:1.25

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

# Copy everything (if they exist)
COPY go.mod* go.sum*

# Then ensure they exist (create empty ones if missing)
RUN [ -f go.mod ] || echo "module temp" > go.mod && \
    [ -f go.sum ] || touch go.sum
  
RUN go mod download

COPY . .

#RUN go build -v -o /usr/local/bin/app ./main.go

CMD ["air"]