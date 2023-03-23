API_DIR=api
V=v1

fix:
	./scripts/fix_openapi.sh ${API_DIR}/${V}/openapi.json

generate: 
	oapi-codegen -package metronome -generate "client, types, spec" \
	${API_DIR}/${V}/openapi.json > metronome.gen.go
