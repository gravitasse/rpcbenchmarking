RM=rm -f
RMFORCE=rm -rf 
SRCS=client.go
COMP_NAME=client_protobuf
all: exe 

exe: $(SRCS)
	go build -o $(COMP_NAME) $(SRCS)

guard:
ifndef SR_CODE_BASE
	$(error SR_CODE_BASE is not set)
endif

install:
	@echo "Test Protobuf has no files to install"
clean:guard
	$(RM) $(COMP_NAME) 
