FROM golang

RUN go get -u github.com/dgageot/demoit

WORKDIR /pres

CMD ["demoit", "-host", "", "-shellhost", ""]