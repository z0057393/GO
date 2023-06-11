FROM  golang:alpine3.18

WORKDIR /app

COPY . .

# Install debugger ( delve )
RUN go install github.com/go-delve/delve/cmd/dlv@latest 
# Install Go-please
RUN go install golang.org/x/tools/gopls@latest