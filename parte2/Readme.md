# Parte 2

1. Debes posicionarte sobre la consola del proyecto de la siguiente manera

```bash
cd .\parte2\main\
```
2. una vez alli, le das un go get e instalas todas las dependencias que se encuentran alli registradas en el go mod

3. cuando tengas todo listo e instalado, inicias el servidor con el siguiente comando
```bash
go run .\main.go
```
4. una vez hayas ejecutado este comando te aparecera un json en la consola donde estara todo el json formado como la estructura esperada.


```bash
PS F:\WORK\assesment\GoTestImuko\parte2\main> go run .\main.go
Successfully Opened CSV file
[{"organization":"org1","users":[{"username":"jperez","roles":["admin ","superadmin"]}]},{"organization":"org1","users":[{"username":"asosa","roles":["writer"]}]},{"organization":"org2","users":[{"username":"jperez","roles":["admin "]}]},{"organization":"org2","users":[{"username":"rrodriguez","roles":["writer","editor"]}]}]
```
- La url te retornara los datos que se requerian en el ejercicio IMUKO parte 1