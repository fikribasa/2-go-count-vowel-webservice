FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go build -o main .
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
#EXPOSE 3030
CMD ["./main"]

# to run build, use the example command below

# docker build -t countvowel .
# docker run -p 3030:3030 countvowel