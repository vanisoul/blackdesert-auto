# 物品分為三種類(img/items)
insertStrToFilenameTailArr(imgs, "type")
insertStrToFilenameTail(img, "type")
type :
    bag 背包上的物品樣式
    bank 銀行上的物品樣式
    Formula 加工時的物品樣式

# 等待某個畫面出現
whilescreen(img string)
=> x y 

# 在某個位置輸入文字
clickLocation(img, x, y, text, 輸入文字前動作func )

# 按鈕
robotgo.KeyTap("a", "control")

# 文字輸入
robotgo.TypeStr("aaaaaa", 0.1)

# leftMosue
leftMosue(x y)

# 新增加工流程
增加 種類.go
呼叫processingTask(.Status, "種類", .Arms, .PearlArms, .Method)
img/is種類.png
img/not種類.png

# clone 需要覆蓋robotgo 原始程式

func FindBitmap 中

```
if len(args) <= 0 {
    FreeBitmap(sbit)
}
```
To
```
if len(args) <= 0 {
    FreeBitmap(sbit)
} else if len(args) > 0 && args[0] == nil {
    FreeBitmap(sbit)
}
```

