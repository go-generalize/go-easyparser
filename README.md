# go-easyparser
Parse Go packages with ease

## Behavior notes (v0.5.0+)

### Newly supported types
- `*types.Alias` (Go 1.23+ `gotypesalias=1`) is resolved via `types.Unalias`. The alias name itself is not preserved (e.g. `type UserID = string` is parsed as `string`).
- `*types.Signature` (function types) falls back to `Any{}`.
- `*types.TypeParam` (generic type parameters) falls back to `Any{}`.
- Generic type aliases (Go 1.24+, e.g. `type Vec[T any] = []T`) are also handled via `types.Unalias`.

### Migration from previous versions
- Inputs containing `func(...)` fields, generic type parameters, or type aliases used to **panic** with `unsupported named type: *types.<Kind>`. They now parse successfully and produce `Any{}` (or the resolved underlying type for aliases). If your downstream code relied on this panic for fail-fast validation, install a custom `Replacer` that explicitly rejects these kinds.
- The panic message wording changed from `unsupported named type:` to `unsupported type:` (the previous wording was misleading since `Map`, `Slice`, etc. are not named types).

### Replacer caveat
- When a `*types.Alias` is encountered, the user-supplied `Replacer` is invoked twice — once on the original alias type, then again on the Unalias-resolved type via the recursive `parseType` call. `Replacer` implementations should be idempotent.
