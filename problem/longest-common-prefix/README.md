# Longest Common Prefix

## 考え方

### 水平スキャン法

- 配列の最初の要素を基準のprefixにする
- 残りの要素を比較し、共通箇所がなくなるまでprefixを短縮する
- 途中でprefixが空になったら、空文字列を返す
