
.PHONY: /tmp/bi.yaml-new /tmp/poly.yaml-new

all: /tmp/bi.yaml-new /tmp/poly.yaml-new

SRC = main.go tidy.go ../../reslist/fs.go

clean:
	rm -f /tmp/bi.yaml-new /tmp/poly.yaml-new
	rm -f makepublicyaml

makepublicyaml:  $(SRC)
	go build -o makepublicyaml main.go tidy.go

/tmp/bi.yaml-new: makepublicyaml
	./makepublicyaml /tmp/resbi /tmp/bi.yaml-new

/tmp/poly.yaml-new: makepublicyaml
	./makepublicyaml /tmp/respoly /tmp/poly.yaml-new

diff:
	@echo ========== BI
	-diff /tmp/bi.yaml /tmp/bi.yaml-new
	@echo ========== POLY
	-diff /tmp/poly.yaml /tmp/poly.yaml-new
