# item-service

## Run 
`go run server/main.go` 


## Call the service
`buf curl --schema proto --data '{"item_id": "1"}' http://localhost:8080/item.v1.ItemStoreService/GetItem`