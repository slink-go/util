# slink/util

Набор мелких утилит и подпакетов, общих для всех сервисов STS: работа с env, файлами, строками, датами, JWT, SQLite.

Go-модуль `go.slink.ws/util`, лицензия Apache-2.0. Отдельный git-репозиторий (`slink-go/util` на GitHub).

## Назначение

Собрать в одной библиотеке то, что иначе дублировалось бы в каждом сервисе: чтение конфигурации из окружения, файловые утилиты, строковые и датные хелперы, работа с токенами и локальной БД.

## Что это и что не

- Это библиотека: точки входа нет.
- Не фреймворк и не SDK: набор независимых пакетов без общего состояния.
- Не библиотека логирования: логирование — в отдельном модуле `slink/logging`.
- Не ORM и не валидатор конфигурации: `env` только читает значения.

## Структура

```
.
├── util.go              # HandleStopSignal — корректная остановка по SIGINT/SIGTERM
├── env/                 # чтение переменных окружения
├── files/               # EnsureDir, Exists, Copy, CopyDirectory, CopySymLink
├── str/                 # Slice, Random, Masked, Capitalize, LongestCommonPrefix, ...
├── date/                # Date (обёртка над time.Time), New/Now/Parse
├── duration/            # (Un)Marshal для time.Duration
├── jwt/                 # Init(secret) + клеймы
│   └── claims/          # имена и значения клеймов (UserId, AuthnType, TenantId, ...)
├── matcher/             # PatternMatcher на регулярных выражениях
├── sqlite/              # OpenDatabase с PRAGMA-настройками
├── cbor/                # заглушка (CBOR WebToken library)
├── go.mod
├── go.sum
├── .gitignore
└── LICENSE
```

## Использование

```go
// env
port := env.IntOrDefault("MON_PORT", 9001)
raw  := env.StringOrDefault("KAZE_NAMESPACE", "default")
arr  := env.StringArrayOrEmpty("SYMBOLS")

// файлы и строки
files.EnsureDir("/data/app")
masked := str.Masked("secret-value", 4)

// даты
d := date.Now()
d, err := date.Parse("2026-07-24", "2006-01-02")

// сопоставление по маскам
m := matcher.NewRegexPatternMatcher("^SBER.*", "^GAZP.*")
m.Matches("SBER@TQBR")

// SQLite
db := sqlite.OpenDatabase("/data/app.db")

// остановка процесса
util.HandleStopSignal(3 * time.Second)
```

## Сборка, тесты

```bash
cd src/lib/slink/util

go build ./...
go vet ./...
go test ./... -race -count=1
go test ./... -coverprofile=cover.out && go tool cover -func=cover.out
```

Модуль включён в `src/go.work`; автономная сборка — `GOWORK=off go build ./...`.

## Конфигурация

Собственных файлов конфигурации нет. Пакет `env` читает переменные процесса:

| Семейство функций | Поведение |
|-------------------|-----------|
| `String`, `StringMust`, `Int`, `IntMust`, `Int64` | строгая форма: ошибка или паника при отсутствии/некорректном значении |
| `StringOrDefault`, `IntOrDefault`, `Int64OrDefault`, `BoolOrDefault`, `Float64OrDefault`, `DurationOrDefault` | безопасная форма: значение по умолчанию |
| `StringWithPrefix`, `StringArrayOrEmpty` | разбор с префиксом и списком |

JWT: `jwt.Init(secret)` — секрет передаётся аргументом, в коде и конфигурации не хранится.

## Интеграция

- Зависимости: `github.com/golang-jwt/jwt`, `github.com/google/uuid`, `github.com/xhit/go-str2duration/v2`, `go.slink.ws/logging`.
- Потребители в STS: практически все сервисы и библиотеки — `svc/rireki`, `svc/kaze`, `lib/slink/monitoring`, `lib/sts/*`. `monitoring` берёт `util/env`, `sqlite` и `sql2` — `util/files`.
- При обновлении зависимостей в дереве STS используется `replace go.slink.ws/util => ...` в `go.mod` потребителей.

## Ограничения

- `env.Int64`, `env.Int`, `env.String` бросают/паникуют при пустой переменной — для сервисов это может приводить к падению на старте.
- `HandleStopSignal` завершает процесс через `os.Exit(0)` после паузы, не дожидаясь завершения горутин: «graceful» здесь означает «дать время на сохранение».
- `str.Random` использует `math/rand` — не криптографически стойкий генератор; для токенов и паролей не подходит.
- `sqlite.OpenDatabase` и `jwt` создают предусловия, при которых ошибка приводит к панике (`OpenDatabase`).
- `cbor` — пустая заглушка (один комментарий и `package`); использовать как CBOR/CWT-библиотеку нельзя.
- `matcher` работает только с регулярными выражениями: произвольные glob-маски не поддерживаются.
