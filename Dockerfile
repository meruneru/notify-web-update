FROM golang:1.17 as build

WORKDIR /app
COPY . /app

WORKDIR /app/src
RUN go mod tidy
RUN GOOS=linux go build -o main .


# copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /app/src/main /main
ENTRYPOINT [ "/main" ]
