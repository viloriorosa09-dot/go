# Slices en Go

Un slice es una colección de tamaño dinámico. A diferencia del array, puede crecer.

## Crear
```go
marcas := make([]string, 2)
marcas[0] = "Mazda"
marcas[1] = "Toyota"
```

## Agregar con append
```go
marcas = append(marcas, "Nissan")
marcas = append(marcas, "Suzuki", "Honda")
```

## Notas
- Asignar por índice fuera del tamaño da panic: `index out of range`.
- `append` devuelve un slice nuevo, así que hay que guardarlo.
- `make([]T, len)` crea el slice con ese tamaño.
