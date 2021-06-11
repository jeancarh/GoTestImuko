# Parte 1 IMUKO

1. Debes posicionarte sobre la consola del proyecto de la siguiente manera

```bash
cd .\parte1\main\
```
2. una vez alli, le das un go get e instalas todas las dependencias que se encuentran alli registradas en el go mod

3. cuando tengas todo listo e instalado, inicias el servidor con el siguiente comando
```bash
go run .\main.go
```
4. una vez hayas ejecutado este comando te aparecera un print que te dira si el servidor esta ON ```Ready to serve at :8081``` , en ese momento vas a probar en postman o en un navegador la siguiente url

```bash
http://localhost:8081/resumen/2019-12-01?dias=1
```
- La url te retornara los datos que se requerian en el ejercicio IMUKO parte 1

5. para salir de la consola de ejecucion usa el siguente atajo ```ctrl + c ```