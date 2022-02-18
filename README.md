# HuangcoolyServer

## 目標

| 功能  |是否達成 |
| ------------- |:-------------:|
| 透過Golang並基於gin框架打造RESTFul API服務     | V     |
| 加入Swagger，以供查看API內容，作為API的線上文件      | Ｖ     |
| 架設mongoDB Altas雲端資料庫，做後Golang後台資料庫|  Ｖ    |
| 透過mongo-go-driver存取mongoDB|  Ｖ    |
| 建立DockerFile為此後台的建立image|  Ｖ    |
|於AWS EC2部署此Docker image|  Ｖ    |
| 建立Unit Test，進行測試|      |
| 將所需要的資源放置AWS S3|      |
| 進行與前台互動的功能開發|      |

## 支線說明
### develope
資料庫MongoDB是使用MongoDB Altas的服務。後續應開發需求的便利，後基於此架構繼續開發


### developDocekrIncludeMongo
1. 將MongoDB的部署也包裝在Docker裡面
1. 此部署方式對於獨立、單一、不需要共用資料庫的服務相當方便
1. 可以加初始資料都放置Volume裡面，再給予MondoDB使用
1. 詳細設定可以參考此支線的`docker-compose.yml`


## 專案使用
1. clone此專案到本機端
1. 安裝docker，可參考[Get Docker ](https://docs.docker.com/get-docker/)
1. 安裝docker-compose，可以參考[Install Compose ](https://docs.docker.com/compose/install/).
1. 用終端機到專案中docker-compose.yml同層資料夾下
1. 進行`docker-compose pull`和`docker-compose up`即可執行此專案
1. 可連線至`http://localhost:8888/swagger/index.html`查看此後台的API
1. MonoDB Altas目前設定有鎖定特定IP位置才能使用，所以目前只提供透過Swagger進行API的查看