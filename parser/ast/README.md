# 递归向下

## Left:

```text
E(k) -> E(k) op(k) E(k+1) | E(k+1)
```

---

## Right

combine

```text
E(k) -> E(k+1) E_(k)
```

```text
var e = new Expr()

e.left = E(k+1)
e.right = E_(k).child(0)
```

race

```text
E_(k) -> op(k) E(k+1) E_(k) | ε
```

终结的版本

```text
// U -> (E) | ++E | --E
E(t) = F E_(t) | U E_(t)
```
