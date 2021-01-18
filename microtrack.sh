#!/bin/bash

echo "Importando período $1 del resumen mensual con plantillaID: $2"

./microtrack -whatdata=monthBrief -period=$1 -templateID=$2