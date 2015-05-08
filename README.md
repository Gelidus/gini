# gini


Gini is simple ini configuration loader in Go which uses reflection to populate given configuration structures from file or reader. Gini is wrapped around [go-ini library](https://github.com/vaughan0/go-ini).

## Usage

This may be your **configuration.ini** file:
```ini
[squirrel]
name=Alfred
age=5
nuts=17

[human]
name=Bob
age=19
atesquirrel=true

# this is a comment
; this is also a comment
```

Now you declare your configuration in Go struct notation and retrieve configuration.ini:
```go
import "github.com/gelidus/gini"

// declare your configuration structure with sections
type Configuration struct {
  Squirrel struct {
    Name string
    Age  int8
    Nuts int16 // cant have too many
  }
  
  Human struct {
    Name        string
    Age         int8
    AteSquirrel bool
  }
}

// create config and retrieve it
config := Configuration{}
gini.ReadFile(&config, "configuration.ini")
```

```go
// or, you can specify any io.Reader
file, err := os.Open("configuration.ini")
if err != nil {
	t.Fatal(err)
}
gini.Read(&config, file)
```

You can now simply access your structure with parsed data from ini file.
