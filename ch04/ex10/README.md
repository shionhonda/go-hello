# GitHub issue tracker

```
$ go run main.go 
Found 64 issues.

Issues created in the last month:
#48950 Alexander encoding/json: calculate correct SyntaxError.Offset in the stream
#48646    piersy encoding/json: unclear documentation for how `json.Unmarshal` and `json.Unmarshaler` interact with `null`

Issues created in the last year:
#48298     dsnet proposal: encoding/json: add Decoder.DisallowDuplicateFields
#42571     dsnet    encoding/json: clarify Decoder.InputOffset semantics
#43716 ggaaooppe encoding/json: increment byte counter when using decoder.Token
#48277 Windsooon encoding/json: add an example for InputOffset() function
#45628 pgundlach                      encoding/xml: add Decoder.InputPos
#43401  opennota             encoding/csv: add Reader.InputOffset method

Issues Older:
#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report different values for SyntaxError.Offset for the same input
#33416   bserdar          encoding/json: This CL adds Decoder.InternKeys
#11046     kurin    encoding/json: Decoder internally buffers full input
#40128  rogpeppe proposal: encoding/json: garbage-free reading of tokens
#5901        rsc encoding/json: allow per-Encoder/per-Decoder registration of marshal/unmarshal functions
#36225     dsnet encoding/json: the Decoder.Decode API lends itself to misuse
#30301     zelch encoding/xml: option to treat unknown fields as an error
#28143    arp242             proposal: encoding/json: add "readonly" tag
#32779       rsc            encoding/json: memoize strings during decode
#29035    jaswdr proposal: encoding/json: add error var to compare  the returned error when using json.Decoder.DisallowUnknownFields()
#14750 cyberphon  encoding/json: parser ignores the case of member names
#31701    lr1980     encoding/json: second decode after error impossible
#40982   Segflow encoding/json: use different error type for unknown field if they are disallowed 
#28923     mvdan            encoding/json: speed up the decoding scanner
#16212 josharian      encoding/json: do all reflect work before decoding
#40983   Segflow encoding/json: return a different error type for unknown field if they are disallowed
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#40127  rogpeppe           encoding/json: add Encoder.EncodeToken method
#41144 alvaroale encoding/json: Unmarshaler breaks DisallowUnknownFields
#34564  mdempsky go/internal/gcimporter: single source of truth for decoder logic
#33854     Qhesz encoding/json: unmarshal option to treat omitted fields as null
#26946    deuill encoding/json: clarify what happens when unmarshaling into a non-empty interface{}
```