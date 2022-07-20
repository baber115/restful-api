


# page
## 路径 E:\Go\pkg\mod\github.com\infraboard\mcube@v1.9.0\pb\page
```proto
syntax = "proto3";

package infraboard.mcube.page;
option go_package = "github.com/infraboard/mcube/http/request";

message PageRequest {
    uint64 page_size = 1;
    uint64 page_number = 2;
    int64 offset = 3;
}
```

# request 
## 路径 E:\Go\pkg\mod\github.com\infraboard\mcube@v1.9.0\pb\request
```proto
syntax = "proto3";

package infraboard.mcube.request;
option go_package = "github.com/infraboard/mcube/pb/request";

enum UpdateMode {
    PUT = 0;
    PATCH = 1;
}
```