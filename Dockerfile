FROM scratch

ENV PORT 8000
EXPOSE $PORT

COPY go-kuber /
CMD ["go-kuber"]