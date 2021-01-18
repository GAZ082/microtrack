#!/bin/bash
# -s -w son para reducir el tamaño del binario, Quita algunas cosas de debug.
go build -ldflags="-s -w"
chmod +x microtrack
scp microtrack sistemas@45.33.2.159:/home/sistemas/scripts/microtrack