
all: data

SRC = main.go 

clean:

airtable2publicyaml:  $(SRC)
	go build -o airtable2publicyaml main.go

.PHONY: data
data: airtable2publicyaml
	source .env ; export AIRTABLE_APIKEY ; export AIRTABLE_BASE_ID ; ./airtable2publicyaml 
