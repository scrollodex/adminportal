
all: data

SRC = main.go tidy.go

clean:

airtable2publicyaml:  $(SRC)
	go build -o airtable2publicyaml $(SRC)

.PHONY: data
data: airtable2publicyaml
	source .env-bi ; export AIRTABLE_APIKEY ; export AIRTABLE_BASE_ID ; ./airtable2publicyaml bi-entries.yaml
	source .env-poly ; export AIRTABLE_APIKEY;export AIRTABLE_BASE_ID ; ./airtable2publicyaml poly-entries.yaml
