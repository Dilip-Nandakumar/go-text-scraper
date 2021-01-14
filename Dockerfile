FROM alpine:latest

ENV name=word-highlights-scraper

COPY ./out/build/word-highlights-scraper-linux /${name}

CMD [ "/word-highlights-scraper" ]