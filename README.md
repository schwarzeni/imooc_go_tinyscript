# TinyScript Go Version

---

[慕课网：程序员三大浪漫--编译原理+操作系统+图形学](https://coding.imooc.com/class/432.html) 编译原理部分 Golang 版本

---

## 笔记

### 右递归

```text
A -> β A'
A' -> α A' | ε
```

例一 Expr -> Expr + 1 | 1 转换成右递归

```text
Expr -> 1 Expr'
Expr' -> +1 Expr' | ε
```

例二 A -> A α | A β | A γ | λ

λ[αβγ]\*

```text
A -> λ A'
A' -> α A' | β A' | γ A' | ε
```

例三 E -> E + E | E - E | d; d -> 0|1|2|3|4|5|6|7|8|9

```text
E -> d E'
E' -> +E E' | -E E' | ε
```

---

### 优先级控制

加减乘除控制

Expr -> Expr + Term | Expr - Term | Term

Term -> Term * Factor | Term / Factor | Factor

Factor -> [0-9]+

去左递归

Expr -> Term Expr_

Expr_ -> + Term Expr_ | - Term Expr_ | ε

Term -> Factor Term_

Term_ -> * Factor Term_ | / Factor Term_ | ε

Factor -> [0-9]+

---

### 领域模型架构

技术细节 --> 领域模型 --> 业务

- 技术细节 (编译器使用何种语言实现？应该使用怎样的数据结构？)
- 领域模型 (词法分析(符号) **-->** 语法分析(抽象语法树)) 需要抽象的思维
- 业务 (TinyScript)

---

## 参考

[elvin-du/tinyscript](https://github.com/elvin-du/tinyscript)
