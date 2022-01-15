
all: data

SRC = main.go tidy.go

clean:

airtable2publicyaml:  $(SRC)
	go build -o airtable2publicyaml $(SRC)

.PHONY: data
data: airtable2publicyaml
	source .env ; export AIRTABLE_APIKEY ; export AIRTABLE_BASE_ID ; ./airtable2publicyaml 
