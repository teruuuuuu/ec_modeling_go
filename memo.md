```
check(){
  SESSION=`echo $1 | grep Set-Cookie | grep -oE "mysession[^;]+" | awk 'END{print}'`
  echo $1
}
```

## ログイン
> RESULT=`curl -i -XPOST -H 'Content-Type:application/json' -d '{"name": "user1", "password": "password" }' http://localhost:8000/auth/login`    
> check $RESULT    

## 検索
> RESULT=`curl -i -XGET -b "$SESSION"  http://localhost:8000/product/search`    
> check $RESULT    

## 商品追加
> RESULT=`curl -i -XPOST -b "$SESSION" -H 'Content-Type:application/json' -d '{"product_id": 3, "number": 2 }' http://localhost:8000/order/updateItem`    
> check $RESULT    

## カート確認
> RESULT=`curl -i -XGET -b "$SESSION" http://localhost:8000/order/cartItems`    
> check $RESULT    

## クーポン追加
> RESULT=`curl -i -XPOST -b "$SESSION" -H 'Content-Type:application/json' -d '{"number": "5678901234" }' http://localhost:8000/order/addCoupon`    
> check $RESULT    

## 購入確定
> RESULT=`curl -i -XPOST -b "$SESSION" -H 'Content-Type:application/json' -d '{"paytype": 1 }' http://localhost:8000/order/confirm`    
> check $RESULT     

## ログアウト
> RESULT=`curl -i -XPOST -b "$SESSION" http://localhost:8000/auth/logout`    
> check $RESULT    


## DB確認
> sqlite3 tmp/gorm.db    
