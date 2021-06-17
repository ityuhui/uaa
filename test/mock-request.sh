# Ref:
# http://www.ruanyifeng.com/blog/2019/09/curl-reference.html
# https://blog.csdn.net/b1303110335/article/details/78213281
# https://blog.csdn.net/afandaafandaafanda/article/details/45039463
# 

# POST (Add)
curl -X POST --data 'name=user1&password=p1' localhost:8080/api/user

# DELETE
#curl -v -X DELETE localhost:8080/api/user/3