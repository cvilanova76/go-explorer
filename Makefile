# Makefile for go-explorer

# Command package path
CMD := ./cmd/cli

# Run the command with the COUNTRY variable (default: Portugal)
run:
	@COUNTRY=$(or $(COUNTRY),Portugal) \
	&& echo "Running with COUNTRY=$${COUNTRY}" \
	&& go run $(CMD) $${COUNTRY}

