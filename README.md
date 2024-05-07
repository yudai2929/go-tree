# go-tree

## 概要

`go-tree`は、Go 言語で実装された二分木と B 木の操作のデモ

## 例

二分木と B 木の操作を提供します。

```golang
func main() {
    // 二分木の操作
    bt := binarytree.New()
    bt.Insert(1)
    bt.Delete(1)

    // B木の操作
    b := btree.New(2)
    b.Insert(1)
    b.Delete(1)
}
```

## 実行

```
make run
```
