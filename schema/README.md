## 実行手順

1. データベースの作成

```mysql
create database example default character set utf8mb4;
```

2. `mysqldef` によってテーブルを作成 \
   設定を変更する場合は`Makefile`を編集する．

```
make sqldef_dry
make sqldef_apply
```

3. `sqlboiler`によって ORM を生成 \
   設定を変更する場合は`sqlboiler.yaml`を編集する．

```
make sqlboiler_apply
```
