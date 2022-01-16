
all: data

SRC = main.go tidy.go

clean:

airtable2publicyaml:  $(SRC)
	go build -o airtable2publicyaml $(SRC)

.PHONY: data
data: airtable2publicyaml bi-entries.yaml poly-entries.yaml

.PHONY: demo-entries.yaml
demo: demo-entries.yaml
demo-entries.yaml: .env-demo airtable2publicyaml
	source .env-demo ; export AIRTABLE_APIKEY ; export AIRTABLE_BASE_ID ; ./airtable2publicyaml demo-entries.yaml

.PHONY: bi
bi: bi-entries.yaml
.PHONY: bi-entries.yaml
bi: bi-entries.yaml
bi-entries.yaml: .env-bi airtable2publicyaml
	source .env-bi ; export AIRTABLE_APIKEY ; export AIRTABLE_BASE_ID ; ./airtable2publicyaml bi-entries.yaml

.PHONY: poly
poly: poly-entries.yaml
.PHONY: poly-entries.yaml
poly: poly-entries.yaml
poly-entries.yaml: .env-poly airtable2publicyaml
	source .env-poly ; export AIRTABLE_APIKEY ; export AIRTABLE_BASE_ID ; ./airtable2publicyaml poly-entries.yaml

