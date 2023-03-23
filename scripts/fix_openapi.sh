#!/bin/bash

set -e

IN_FILE=$1
TMP_FILE=$(mktemp)

main() {
  GoNameAttr="x-go-name"
  GoNameAttrSuffix="Param"

  cp "$IN_FILE" "$TMP_FILE"

  for i in BillingProviderType InvoiceStatus ; do
  		contents=$(jq ".components.parameters.$i[\"$GoNameAttr\"] = \"$i$GoNameAttrSuffix\"" "$TMP_FILE")
  		echo -E "$contents" > "$TMP_FILE"
  done

  mv "$TMP_FILE" "$IN_FILE"
}

main