# Go Server

Source: https://cryptic.io/go-http/

https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.0.html


# DB

## CMD to start mysql db docker

```
docker run \
-d \
--name=test-mysql \
--env="MYSQL_ROOT_PASSWORD=mypassword" \
-p 6604:3306 \
-v $(pwd)/mysql/conf.d:/etc/mysql/conf.d \
-v $(pwd)/db:/var/lib/mysql \
mysql
```
[doc mysql docker link](https://hub.docker.com/_/mysql/)