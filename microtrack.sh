#!/bin/bash

from=$1
to=$(date -d "$(date +$1)+1month-1day" +%Y-%m-%d)

echo "Importando desde $from a $to del resumen mensual."

./microtrack -host=45.33.2.159 -database=microtrack -schema=public -user=leer_insertar_user -pass='7E$VjsNDaCyz#gaE' -from="$from" -to="$to" -whatdata="monthBrief"