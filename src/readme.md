restful文档: https://blog.csdn.net/qq_41606973/article/details/86352787
微服务文档 https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design


URI 表示资源，资源一般对应服务器端领域模型中的实体类。
不用大写；
用中杠-不用下杠_;
URI中的名词表示资源集合，使用复数形式。

资源集合：
/zoos //所有动物园
/zoos/1/animals //id为1的动物园中的所有动物

单个资源：
/zoos/1 //id为1的动物园
/zoos/1;2;3 //id为1，2，3的动物园


/在url中表达层级，用于按实体关联关系进行对象导航，一般根据id导航。
过深的导航容易导致url膨胀，不易维护，如 GET /zoos/1/areas/3/animals/4，尽量使用查询参数代替路径中的实体导航，
如GET /animals?zoo=1&area=3；

对Composite资源的访问
一个常见的例子是 User — Address，Address是对User表中zipCode/country/city三个字段的简单抽象，无法独立于User存在。
必须通过User索引到Address：GET /user/1/addresses

response 的 body 直接就是数据，不要做多余的包装。错误示例：
{
    "success":true,
    "data":{"id":1,"name":"xiaotuan"},
}
·	response 格式
GET	       单个对象、集合
POST	   新增成功的对象
PUT/PATCH  更新成功的对象
DELETE	   空

entity 与数据库表一一对应
dto 中间转换层
vo 与界面一一对应


在更复杂的系统中，提供允许客户端在多个级别的关系中导航的URI很有诱惑力，例如/customers/1/orders/99/products。
然而，如果资源之间的关系在未来发生变化，则这种复杂程度可能难以维护并且不灵活。相反，尝试保持URI相对简单。一旦应用程序引用了资源，
就应该可以使用此引用来查找与该资源相关的项目。可以使用URI替换上述查询/customers/1/orders以查找客户1的所有订单，
然后/orders/99/products按此顺序查找产品。
1.尽量最多两层
/customers/1/orders post,get,put,delete 对某一客户的订单的增删改查
/orders post,get,put,delete 对所有订单的增删改查
/orders/99/products 对某一订单的产品增删改查

避免在Web API和底层数据源之间引入依赖关系。例如，如果您的数据存储在关系数据库中，则Web API不需要将每个表公开为资源集合。
事实上，这可能是一个糟糕的设计。相反，将Web API视为数据库的抽象。如有必要，在数据库和Web API之间引入映射层。
这样，客户端应用程序就与底层数据库方案的更改隔离开来。

可能无法将Web API实现的每个操作映射到特定资源。您可以通过调用函数的HTTP请求来处理此类非资源方案，并将结果作为HTTP响应消息返回。
例如，实现简单计算器操作（如加法和减法）的Web API可以提供将这些操作公开为伪资源的URI，并使用查询字符串指定所需的参数。
例如，对URI / add？operand1 = 99＆operand2 = 1的GET请求将返回一个响应消息，其中正文包含值100.但是，只能谨慎使用这些形式的URI。

资源                   POST                        GET                    PUT                            DELETE
/customers             创建一个新客户               检索所有客户            批量更新客户                     删除所有客户
/customers/1           错误                        检索客户1的详细信息      更新客户1的详细信息（如果存在）   删除客户1
/customers/1/orders    为客户1创建新订单            检索客户1的所有订单      批量更新客户1的订单              删除客户1的所有订单
/orders                创建一个新订单(正常用不到)    检索所有订单            批量更新订单                     删除所有订单
/orders/1              错误                        检索订单1的详细信息      更新订单1的详细信息              删除订单1
/orders/1/products     为订单1创建个产品            检索订单1的所有产品      批量更新订单1的产品              删除订单1所有产品

