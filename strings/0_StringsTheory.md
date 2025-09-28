## Основы

- Тип string в Go — это неизменяемая последовательность байт.

- Символы хранятся в UTF-8, поэтому один символ (rune) может занимать несколько байт.

- Длина строки через len(s) возвращает количество байт, а не символов.

```go
s := "Привет"
fmt.Println(len(s)) // 12 (6 букв × 2 байта)
```

## Работа с символами
> Для доступа к символам используют []rune.

- Перебор строки через for range возвращает позицию в байтах и rune.
```go
runes := []rune("Привет!")
fmt.Println(len(runes)) // 6
fmt.Println(string(runes[0])) // П
```


## Изменяемость
> Строки иммутабельные.

- Чтобы изменить, переводим в []rune или []byte, потом обратно в string.
```go
runes := []rune("молоко")
runes[0] = 'x'
s := string(runes)
fmt.Println(s) //  холоко
```

## Создание и объединение
Конкатенация через +.

Форматирование через fmt.Sprintf.

Многострочные строки через `…` (raw string literal).
```go
name := "Alice"
age := 30
s := fmt.Sprintf("Меня зовут %s, мне %d", name, age)
```

## Работа со строками (strings пакет)

Поиск и проверка:
- strings.Contains(s, "go")
- strings.HasPrefix(s, "http")
- strings.Index(s, "lang")
Разделение и склейка:
- strings.Split(s, " ")
- strings.Join(slice, "-")
Замена и регистр:
- strings.ReplaceAll(s, "cat", "dog")
- strings.ToUpper(s)
- strings.ToLower(s)
Обрезка:
- strings.TrimSpace(s)
- strings.Trim(s, "!?")
Повторение
- strings.Repeat("ha", 3)

## Байты и строки

> Преобразование в []byte и обратно — часто для ввода/вывода и сетевых операций.
```go
b := []byte("hello") // строка -> байты
s := string(b) // бакйты -> строка
```

## Сравнение строк

> ==, <, > работают лексикографически (по Unicode).
strings.EqualFola(a, b) -> сравнение без учета регистра


## Подстроки (срезы)
- Можно брать срез строки, но это срез по байтам, не по символам.
- Для кириллицы и других символов лучше переводить в []rune.
```go
s := "hello"
fmt.Println(s[:2]) // "he"

runes := []rune("Привет")
fmt.Println(string(runes[:3])) // "При"
```